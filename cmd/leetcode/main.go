package main

import (
	"leetcode-backend/configs"
	"leetcode-backend/constants"
	"leetcode-backend/internal/submissions"
	container "leetcode-backend/pkg"
	"os"
	"strings"
	"time"

	"go.uber.org/zap"
)

type Context struct {
	logger     *zap.Logger
	controller *container.Controller
}

func main() {
	//Initialisation
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	c, errNewController := container.NewController()

	if errNewController != nil {
		logger.Info("", zap.Error(errNewController))
		os.Exit(1)
	}
	ctx := &Context{logger, c}

	createVolumes(ctx)
	submissionsChan := make(chan submissions.Submission, configs.SUBMISSIONS_CHAN_BUFFER)
	resultChan := make(chan submissions.Result, configs.SUBMISSIONS_CHAN_BUFFER)

	for i := 0; i < 3; i++ {
		go worker(ctx, submissionsChan, resultChan)
		go printResult(ctx, resultChan)
		//TODO:: Return result to client
	}

	//TODO:: Listen to a message queue like kafka or RabbitMQ
	for i := 0; 1 < 2; i++ {
		// Print summation, min and max of array
		submissionsChan <- submissions.Submission{
			UserId:         "u2",
			ProblemId:      "p1",
			Language:       constants.PYTHON,
			Code:           "# import inbuilt standard input output\nfrom sys import stdin, stdout\n\nn = stdin.readline()\narr = [int(x) for x in stdin.readline().split()]\n\nsummation = 0\nfor num in arr:\n    summation += num\nminimum = min(arr)\nmaximum = max(arr)\nprint(summation, minimum, maximum)\n",
			Testcase:       "3\n1 2 3",
			ExpectedOutput: "6 1 3",
		}
		time.Sleep(time.Millisecond * 200)

		//CPP
		submissionsChan <- submissions.Submission{
			UserId:         "u2",
			ProblemId:      "p1",
			Language:       constants.CPP,
			Code:           `#include<bits/stdc++.h>\nusing namespace std;\n\nint main(){\n    int n;\n    cin>>n;\n    int arr[n], i, mi=INT_MAX, ma = INT_MIN, sum=0;\n    for(i=0;i<n;i++){\n        cin>>arr[i];\n        mi = min(mi, arr[i]);\n        ma = max(ma, arr[i]);\n        sum += arr[i];\n    }\n    cout<<sum<<" "<<*min_element(arr, arr+n)<<" "<<ma<<"\\n";\n    return 0;\n}`,
			Testcase:       "3\n1 2 3",
			ExpectedOutput: "6 1 3",
		}
		time.Sleep(time.Millisecond * 200)

		//Python Wrong answer
		submissionsChan <- submissions.Submission{
			UserId:         "u2",
			ProblemId:      "p1",
			Language:       constants.PYTHON,
			Code:           "# import inbuilt standard input output\nfrom sys import stdin, stdout\n\nn = stdin.readline()\narr = [int(x) for x in stdin.readline().split()]\n\nsummation = 0\nfor num in arr:\n    summation += num\nminimum = min(arr)\nmaximum = max(arr)\nprint(summation, minimum, maximum+1)\n",
			Testcase:       "3\n1 2 3",
			ExpectedOutput: "6 1 3",
		}
		time.Sleep(time.Millisecond * 200)

		//CPP Wrong answer
		submissionsChan <- submissions.Submission{
			UserId:         "u2",
			ProblemId:      "p1",
			Language:       constants.CPP,
			Code:           `#include<bits/stdc++.h>\nusing namespace std;\n\nint main(){\n    int n;\n    cin>>n;\n    int arr[n], i, mi=INT_MAX, ma = INT_MIN, sum=0;\n    for(i=0;i<n;i++){\n        cin>>arr[i];\n        mi = min(mi, arr[i]);\n        ma = max(ma, arr[i]);\n        sum += arr[i];\n    }\n    cout<<sum<<" "<<*min_element(arr, arr+n)<<" "<<ma+1<<"\\n";\n    return 0;\n}`,
			Testcase:       "3\n1 2 3",
			ExpectedOutput: "6 1 3",
		}
		time.Sleep(time.Millisecond * 200)

		//CPP Compilation error
		submissionsChan <- submissions.Submission{
			UserId:         "u2",
			ProblemId:      "p1",
			Language:       constants.CPP,
			Code:           `#exclude<bits/stdc++.h>\nusing namespace std;\n\nint main(){\n    int n;\n    cin>>n;\n    int arr[n], i, mi=INT_MAX, ma = INT_MIN, sum=0;\n    for(i=0;i<n;i++){\n        cin>>arr[i];\n        mi = min(mi, arr[i]);\n        ma = max(ma, arr[i]);\n        sum += arr[i];\n    }\n    cout<<sum<<" "<<*min_element(arr, arr+n)<<" "<<ma+1<<"\\n";\n    return 0;\n}`,
			Testcase:       "3\n1 2 3",
			ExpectedOutput: "6 1 3",
		}
		time.Sleep(time.Millisecond * 200)

		//Python Compilation error
		submissionsChan <- submissions.Submission{
			UserId:         "u2",
			ProblemId:      "p1",
			Language:       constants.PYTHON,
			Code:           "# import inbuilt standard input output\nfrom sys export stdin, stdout\n\nn = stdin.readline()\narr = [int(x) for x in stdin.readline().split()]\n\nsummation = 0\nfor num in arr:\n    summation += num\nminimum = min(arr)\nmaximum = max(arr)\nprint(summation, minimum, maximum+1)\n",
			Testcase:       "3\n1 2 3",
			ExpectedOutput: "6 1 3",
		}
		time.Sleep(time.Millisecond * 200)

		// //Golang Wrong Answer
		// submissionsChan <- submissions.Submission{
		// 	UserId:    "u2",
		// 	ProblemId: "p1",
		// 	Language:  constants.GOLANG,
		// 	Code:      `package main\n\nimport (\n\t"fmt"\n\t"math"\n)\n\nfunc main() {\n\tvar n int\n\tmi := math.MaxInt\n\tma := math.MinInt\n\tsum := 0\n\tfmt.Scanln(&n)\n\tarr := make([]int, n)\n\tfor i := 0; i < n; i++ {\n\t\tfmt.Scanf("%d", &arr[i])\n\t\tif arr[i] < mi {\n\t\t\tmi = arr[i]\n\t\t}\n\t\tif arr[i] > ma {\n\t\t\tma = arr[i]\n\t\t}\n\t\tsum += arr[i]\n\t}\n\tfmt.Printf("%d %d %d\\n", sum, mi, ma)\n}\n`,
		// 	// Code:           `package main\n\nimport (\n\t"fmt"\n)\n\nfunc main() {\n\tfmt.Printf("Hello World\\n")\n}\n`,
		// 	Testcase:       "3\n1 2 3",
		// 	ExpectedOutput: "6 1 3",
		// }
		// time.Sleep(time.Millisecond * 200)
	}
}

func printResult(ctx *Context, resultChan chan submissions.Result) {
	for result := range resultChan {
		ctx.logger.Info(result.VerdictMessage + ": " + result.ErrorMessage)
	}
}

func createVolumes(ctx *Context) {
	for language := range constants.LanguagesMap {
		createdNow, _, err := ctx.controller.EnsureVolume(language + "_volume")
		if err != nil {
			ctx.logger.Error(language+" volume not created", zap.Error(err))
			os.Exit(1)
		}
		if createdNow == true {
			ctx.logger.Info(language + " volume created now as it was not already present")
		}
	}
}

func worker(ctx *Context, submissionsChan <-chan submissions.Submission, resultChan chan<- submissions.Result) {
	for submission := range submissionsChan {
		go getResult(ctx, submission, resultChan)
	}
}

func getResult(ctx *Context, submission submissions.Submission, resultChan chan<- submissions.Result) {
	createdNow, volume, err := ctx.controller.EnsureVolume(submission.Language + "_volume")
	if err != nil {
		ctx.logger.Error(submission.Language+" volume not created", zap.Error(err))
		os.Exit(1)
	}
	if createdNow == true {
		ctx.logger.Info(submission.Language + " volume created now as it was not already present")
	}

	image := submission.Language + "_image"
	volumeMounts := []container.VolumeMount{{HostPath: "/volume", Volume: volume}}
	command := []string{"sh", "run_" + submission.Language + ".sh", submission.Code, submission.Testcase}
	statusCode, actualOutput, err := ctx.controller.ContainerRunAndClean(image, command, volumeMounts)

	verdictMessage, errorMessage := getVerdictAndErrorMessage(ctx, submission, statusCode, actualOutput, err)
	resultChan <- submissions.Result{
		OriginalSubmission:    submission,
		ActualOutputEncrypted: []byte(actualOutput),
		VerdictMessage:        verdictMessage,
		ErrorMessage:          errorMessage,
	}
}

func getVerdictAndErrorMessage(ctx *Context, submission submissions.Submission, statusCode int64, actualOutput string, err error) (string, string) {
	if statusCode == 256 { //Own definition
		return "TLE", submissions.TLE
	}

	actualOutput = strings.ReplaceAll(actualOutput, "\r\n", "\n")
	actualOutput = strings.TrimRight(actualOutput, "\n")
	expectedOutput := strings.TrimRight(submission.ExpectedOutput, "\n")
	if strings.Compare(actualOutput, expectedOutput) != 0 {
		return "WA", actualOutput
	}
	return "AC", ""
}
