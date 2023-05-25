python stringToPythonFile.py "$1"
python stringToInputFile.py "$2"
python submissionFile.py < inputFile.txt
