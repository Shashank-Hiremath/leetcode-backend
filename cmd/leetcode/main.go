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

	// if err = c.EnsureImage("python_image"); err != nil {
	// 	logger.Info("", zap.Error(err))
	// 	os.Exit(1)
	// }

	createVolumes(ctx)
	submissionsChan := make(chan submissions.Submission, configs.SUBMISSIONS_CHAN_BUFFER)
	resultChan := make(chan submissions.Result, configs.SUBMISSIONS_CHAN_BUFFER)

	for i := 0; i < 1; i++ {
		go worker(ctx, submissionsChan, resultChan)
		go printResult(ctx, resultChan)
	}

	//go listen submissions
	//TODO:: Listen to a message queue like kafka or RabbitMQ
	for i := 0; 1 < 2; i++ {
		submissionsChan <- *submissions.GetSubmissions()
		time.Sleep(time.Millisecond * 100)
		// ctx.logger.Info(fmt.Sprint("===", i))
		// time.Sleep(10 * time.Minute)
	}
}

func printResult(ctx *Context, resultChan chan submissions.Result) {
	for result := range resultChan {
		ctx.logger.Info(result.VerdictMessage + "\n")
		ctx.logger.Info(result.ErrorMessage + "\n")
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
		return "WA", submissions.WA
	}
	return "AC", submissions.AC //TODO::BUG:: 2nd value should not be sent as errorMessage, should be empty
}
