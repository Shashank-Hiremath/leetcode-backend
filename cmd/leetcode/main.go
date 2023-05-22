package main

import (
	"fmt"
	container "leetcode-backend/pkg"
	"os"
	"strings"

	"go.uber.org/zap"
)

func main() {
	//Initialisation
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	c, err := container.NewController()

	if err != nil {
		logger.Info("", zap.Error(err))
		os.Exit(1)
	}

	// err = c.EnsureImage("gcc")

	if err = c.EnsureImage("python"); err != nil {
		logger.Info("", zap.Error(err))
		os.Exit(1)
	}

	//TODO:: Should delete if some volume already existing same name?
	createdNow, volume, err1 := c.EnsureVolume("pythonVolume")

	if err1 != nil {
		logger.Info("", zap.Error(err1))
		os.Exit(1)
	}

	if createdNow != true {
		logger.Info("Did not create a new volume as it was already present")
	}

	mounts := []container.VolumeMount{
		{
			HostPath: "/volume",
			Volume:   volume,
		},
	}

	// submissions := make(chan, int)

	//factory pattern
	//go process submissions

	expectedOutput := "6 1 3"
	submittedCode := "# import inbuilt standard input output\nfrom sys import stdin, stdout\n\nn = stdin.readline()\narr = [int(x) for x in stdin.readline().split()]\n\nsummation = 0\nfor num in arr:\n    summation += num\nminimum = min(arr)\nmaximum = max(arr)\nprint(summation, minimum, maximum)\n"
	input := "3\n1 2 3"
	statusCode, actualOutput, err := c.ContainerRunAndClean("python_image", []string{"sh", "runPython.sh", submittedCode, input}, mounts)
	actualOutput = strings.ReplaceAll(actualOutput, "\r\n", "\n")
	actualOutput = strings.TrimRight(actualOutput, "\n")
	expectedOutput = strings.TrimRight(expectedOutput, "\n")

	//go listen submissions

	//validate actual vs expected output
	fmt.Println(strings.Compare(actualOutput, expectedOutput))
	logger.Info("\n---", zap.Int64("statusCode", statusCode), zap.String("actualOutput", actualOutput), zap.Error(err))
}
