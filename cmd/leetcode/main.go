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

	// if err = c.EnsureImage("cpp_image"); err != nil {
	// 	logger.Info("", zap.Error(err))
	// 	os.Exit(1)
	// }

	// if err = c.EnsureImage("python_image"); err != nil {
	// 	logger.Info("", zap.Error(err))
	// 	os.Exit(1)
	// }

	//TODO:: Should delete if some volume already existing same name?
	// createdNow, volume, err1 := c.EnsureVolume("pythonVolume")
	createdNow, volume, err1 := c.EnsureVolume("cppVolume")

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

	// expectedOutput := "6 1 3"
	// submittedCode := "# import inbuilt standard input output\nfrom sys import stdin, stdout\n\nn = stdin.readline()\narr = [int(x) for x in stdin.readline().split()]\n\nsummation = 0\nfor num in arr:\n    summation += num\nminimum = min(arr)\nmaximum = max(arr)\nprint(summation, minimum, maximum)\n"
	// input := "3\n1 2 3"
	// statusCode, actualOutput, err := c.ContainerRunAndClean("python_image", []string{"sh", "runPython.sh", submittedCode, input}, mounts)
	// actualOutput = strings.ReplaceAll(actualOutput, "\r\n", "\n")
	// actualOutput = strings.TrimRight(actualOutput, "\n")
	// expectedOutput = strings.TrimRight(expectedOutput, "\n")

	expectedOutput := "6 1 3"
	submittedCode := `#include<bits/stdc++.h>\nusing namespace std;\n\nint main(){\n    int n;\n    cin>>n;\n    int arr[n], i, mi=INT_MAX, ma = INT_MIN, sum=0;\n    for(i=0;i<n;i++){\n        cin>>arr[i];\n        mi = min(mi, arr[i]);\n        ma = max(ma, arr[i]);\n        sum += arr[i];\n    }\n    cout<<sum<<" "<<*min_element(arr, arr+n)<<" "<<ma<<"\\n";\n    return 0;\n}`
	input := "3\n1 2 3"
	statusCode, actualOutput, err := c.ContainerRunAndClean("cpp_image", []string{"sh", "runCpp.sh", submittedCode, input}, mounts)
	actualOutput = strings.ReplaceAll(actualOutput, "\r\n", "\n")
	actualOutput = strings.TrimRight(actualOutput, "\n")
	expectedOutput = strings.TrimRight(expectedOutput, "\n")

	//go listen submissions

	//validate actual vs expected output
	fmt.Println(actualOutput)
	fmt.Println(expectedOutput)
	fmt.Println(strings.Compare(actualOutput, expectedOutput))
	logger.Info("\n---", zap.Int64("statusCode", statusCode), zap.String("actualOutput", actualOutput), zap.Error(err))
}
