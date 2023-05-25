echo $1 > submissionFile.go
echo $2 > inputFile.txt

go run submissionFile.go < ./inputFile.txt
