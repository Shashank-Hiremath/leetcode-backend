# Leetcode backend

1. Consumes code submissions from a messaging queue like Kafka(ToDo, using dummy for now)
2. Processes and run them in a docker container
3. Autoscales if number of problems per second increases(20-100 per second)

1. TLE if submisison takes more than 5 seconds
2. Compilation error. Send the error message
3. Output not as expected. Return WA
4. Passes all test cases. Return AC

1. Allow user input or testcases
2. 

## Quick Run
sh run.sh