package submissions

import (
	"leetcode-backend/constants"
)

const (
	AC      string = "Answered Correct"
	WA             = "Wrong Answer"
	TLE            = "Time Limit Exceeded"
	CE             = "Compilation Error"
	MLE            = "Memory limit exceeded"
	SIGSEGV        = "Invalid Memory access"
)

type Submission struct {
	Id             string
	UserId         string
	ProblemId      string
	Language       string
	Code           string
	Testcase       string
	ExpectedOutput string
}

type Result struct {
	OriginalSubmission    Submission
	ActualOutputEncrypted []byte //Encrypted so that the user doesn't findout test cases by printing them. Still sending if want to show diff with expected output or if it is compiler error message
	VerdictMessage        string
	ErrorMessage          string
}

func GetSubmissions() *Submission {
	//TODO:: consume from a messaging queue
	return &Submission{
		UserId:         "u11",
		ProblemId:      "p1",
		Language:       constants.PYTHON,
		Code:           "# import inbuilt standard input output\nfrom sys import stdin, stdout\n\nn = stdin.readline()\narr = [int(x) for x in stdin.readline().split()]\n\nsummation = 0\nfor num in arr:\n    summation += num\nminimum = min(arr)\nmaximum = max(arr)\nprint(summation, minimum, maximum)\n",
		Testcase:       "3\n1 2 3",
		ExpectedOutput: "6 1 3",
	}
}
