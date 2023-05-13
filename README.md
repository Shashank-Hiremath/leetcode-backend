Leetcode backend

1. Consumes code submissions from Kafka
2. Processes and run them in a docker
3. Autoscales if number of problems per second increases(20-100 per second)

1. TLE if submisison takes more than 2 seconds
2. Compilation error. Send the error message
3. Output not as expected. Return WA
4. Passes all test cases. Return AC
