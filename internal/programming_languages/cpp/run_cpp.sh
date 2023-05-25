echo $1 > submissionFile.cpp
echo $2 > inputFile.txt

gcc  -lstdc++ -o submissionFile submissionFile.cpp && ./submissionFile < ./inputFile.txt
