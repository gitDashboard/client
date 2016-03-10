package client

import (
	"fmt"
	git "gopkg.in/libgit2/git2go.v22"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
)

const (
	DELTA_Y int = 30
	DELTA_X int = 20
)

type GraphCommit struct {
	Commit  *git.Commit
	Message string
	Branch  string
	X, Y    int
	Color   string
}

func NewGraphCommit(commit *git.Commit, branch string) *GraphCommit {
	msg := strings.Replace(commit.Message(), "\n", "<br/>", -1)
	msg = strings.Replace(msg, "\"", "'", -1)
	return &GraphCommit{Commit: commit, Branch: branch, Message: msg}
}

func addTimeCommit(graphCommit *GraphCommit, TimeCommits map[string][]*GraphCommit) {
	year, monthS, day := graphCommit.Commit.Committer().When.Date()
	month := int(monthS)
	timeKey := fmt.Sprintf("%d-%02d-%02d", year, month, day)
	if _, timeExist := TimeCommits[timeKey]; !timeExist {
		TimeCommits[timeKey] = make([]*GraphCommit, 0)
	}
	TimeCommits[timeKey] = append(TimeCommits[timeKey], graphCommit)
}

func repoWalker(AllCommits map[string]*GraphCommit, BranchCommits map[string][]*GraphCommit, TimeCommits map[string][]*GraphCommit, Connections map[string][]string, branch string, commit *git.Commit) bool {
	_, alreadyPresent := AllCommits[commit.Id().String()]
	if len(BranchCommits[branch]) == 0 {
		BranchCommits[branch] = make([]*GraphCommit, 0)
	}
	if !alreadyPresent {
		grpCommit := NewGraphCommit(commit, branch)
		AllCommits[commit.Id().String()] = grpCommit
		BranchCommits[branch] = append(BranchCommits[branch], grpCommit)
		addTimeCommit(grpCommit, TimeCommits)

		//check parents
		for p := uint(0); p < commit.ParentCount(); p++ {
			parent := commit.Parent(p)
			if parent != nil {
				_, parentAlreadyPresent := AllCommits[parent.Id().String()]
				if parentAlreadyPresent {
					if _, connPresent := Connections[parent.Id().String()]; !connPresent {
						Connections[parent.Id().String()] = make([]string, 1)
					}
					Connections[parent.Id().String()] = append(Connections[parent.Id().String()], commit.Id().String())
				}
			}
		}
	}
	return true

}

type ByTime []*GraphCommit

func (a ByTime) Len() int      { return len(a) }
func (a ByTime) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a ByTime) Less(i, j int) bool {
	return a[i].Commit.Committer().When.UnixNano() < a[j].Commit.Committer().When.UnixNano()
}

func drawCommit(cmt *GraphCommit) string {
	message := "Id:" + cmt.Commit.Id().String()
	parent := cmt.Commit.Parent(0)
	if parent != nil {
		message = message + "  parent:" + parent.Id().String() + "date:" + strconv.FormatInt(cmt.Commit.Committer().When.Unix(), 10)
	}
	return fmt.Sprintf("var %s = createCommit(stage,%d,%d,\"%s\");\n", "cmt_"+cmt.Commit.Id().String(), cmt.X, cmt.Y, message)
}

func drawCommitConnection(src, dst *GraphCommit) string {
	var result string
	if dst != nil && src != nil {
		result = result + fmt.Sprintf("connectCommit(stage,%s,%s);\n", "cmt_"+src.Commit.Id().String(), "cmt_"+dst.Commit.Id().String())
	}
	return result
}

func drawDate(y, x1, x2 int, color, date string) string {
	return fmt.Sprintf("drawDate(stage,%d,%d,%d,'%s','%s');\n", y, x1, x2, color, date)
}

func splitDay(day string) (string, string, string) {
	dayParts := strings.Split(day, "-")
	return dayParts[0], dayParts[0] + "-" + dayParts[1], day
}

func GenerateGraph(gitDirPath string) (string, error) {
	var BranchCommits map[string][]*GraphCommit
	var AllCommits map[string]*GraphCommit
	var TimeCommits map[string][]*GraphCommit
	var Connections map[string][]string
	BranchCommits = make(map[string][]*GraphCommit)
	AllCommits = make(map[string]*GraphCommit)
	TimeCommits = make(map[string][]*GraphCommit)
	Connections = make(map[string][]string)

	gitDir, _ := filepath.Abs(gitDirPath)

	repo, err := git.OpenRepository(gitDir)
	if err != nil {
		return "", err
	}
	defer repo.Free()
	result := "function drawGraph(stage){"

	//finding all branch
	refIt, err := repo.NewReferenceIterator()
	if err != nil {
		return "", err
	}

	defer refIt.Free()
	refNameIt := refIt.Names()
	refName, refNameErr := refNameIt.Next()

	//set master as reference branch
	branches := make([]string, 1)
	branches[0] = "refs/heads/master"
	for refNameErr == nil {
		if refName != "refs/heads/master" {
			branches = append(branches, refName)
		}
		refName, refNameErr = refNameIt.Next()
	}
	var notEmptyBranches []string
	//read branch commits
	for _, branch := range branches {
		walk, err := repo.Walk()
		if !strings.Contains(branch, "refs/tags") && !strings.Contains(branch, "refs/remotes") {
			walk.PushRef(branch)
			if err != nil {
				return "", err
			}
			defer walk.Free()
			walk.Sorting(git.SortTopological | git.SortTime)
			walk.Iterate(func(commit *git.Commit) bool {
				return repoWalker(AllCommits, BranchCommits, TimeCommits, Connections, branch, commit)
			})
			//remove merged branches
			if len(BranchCommits[branch]) > 0 {
				notEmptyBranches = append(notEmptyBranches, branch)
			}
		}
	}

	//setting Y coord
	baseY := 45
	maxY := 0
	for b, branch := range notEmptyBranches {
		branchCmts, _ := BranchCommits[branch]
		for _, cmt := range branchCmts {
			cmt.Y = baseY + (b+1)*DELTA_Y

		}

	}
	maxY = baseY + (len(notEmptyBranches))*DELTA_Y

	yearCoords := make(map[string][]int)
	monthCoords := make(map[string][]int)
	dayCoords := make(map[string][]int)

	//set X coord
	days := make([]string, 0)
	for day, _ := range TimeCommits {
		days = append(days, day)
	}
	sort.Sort(sort.StringSlice(days))
	maxX := 0
	for _, day := range days {
		maxXDay := -1
		dayCommits := TimeCommits[day]
		sort.Sort(ByTime(dayCommits))

		//preparing date min/max coord maps
		year, month, day := splitDay(day)
		if _, present := yearCoords[year]; !present {
			yearCoords[year] = make([]int, 2, 2)
			yearCoords[year][0] = maxX + (DELTA_X / 2)
		}
		if _, present := monthCoords[month]; !present {
			monthCoords[month] = make([]int, 2, 2)
			monthCoords[month][0] = maxX + (DELTA_X / 2)
		}
		dayCoords[day] = make([]int, 2, 2)
		dayCoords[day][0] = maxX + (DELTA_X / 2)

		for d, cmt := range dayCommits {
			cmt.X = maxX + (d+1)*DELTA_X
			if maxXDay < cmt.X {
				maxXDay = cmt.X
			}
			result += drawCommit(cmt)
		}
		maxX = maxXDay
		dayCoords[day][1] = maxX + (DELTA_X / 2)
		monthCoords[month][1] = maxX + (DELTA_X / 2)
		yearCoords[year][1] = maxX + (DELTA_X / 2)
	}

	//horizontal branch line
	for _, branch := range notEmptyBranches {
		lastBranch := BranchCommits[branch][0]
		firstBranch := BranchCommits[branch][len(BranchCommits[branch])-1]
		result += drawCommitConnection(firstBranch, lastBranch)
	}

	for src, dsts := range Connections {
		for _, dst := range dsts {
			result += drawCommitConnection(AllCommits[src], AllCommits[dst])
		}
	}

	//draw calendar
	for year, coords := range yearCoords {
		result += drawDate(0, coords[0], coords[1], "#3D3E40", year)
	}
	for month, coords := range monthCoords {
		monthParts := strings.Split(month, "-")
		result += drawDate(20, coords[0], coords[1], "#2D2E2E", monthParts[1])
	}
	for day, coords := range dayCoords {
		dayParts := strings.Split(day, "-")
		result += drawDate(40, coords[0], coords[1], "#6D6E70", dayParts[2])
	}
	result += "}"
	result += "function getHeight(){return " + strconv.Itoa(maxY) + ";}"
	result += "function getWidth(){return " + strconv.Itoa(maxX) + ";}"
	result += "return {'drawGraph':drawGraph,'getHeight':getHeight,'getWidth':getWidth};"

	return result, nil
}
