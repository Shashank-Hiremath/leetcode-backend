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

	// err = c.EnsureImage("golang")

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
		logger.Info("Created a new volume as it wasn't already present")
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

	expectedOutput := "Hello111\napple\nball\nBye456\n"
	submittedCode := "print('Hello111')\nfor line in ['apple', 'ball']:\n    print(line)\nprint('Bye456')"
	statusCode, actualOutput, err := c.ContainerRunAndClean("python_image", []string{"sh", "runPython.sh", submittedCode}, mounts)
	actualOutput = strings.ReplaceAll(actualOutput, "\r\n", "\n")
	actualOutput = strings.TrimRight(actualOutput, "\n")
	expectedOutput = strings.TrimRight(expectedOutput, "\n")

	//go listen submissions

	//validate actual vs expected output
	fmt.Println(strings.Compare(actualOutput, expectedOutput))
	logger.Info("\n---", zap.Int64("statusCode", statusCode), zap.String("actualOutput", actualOutput), zap.Error(err))
}
