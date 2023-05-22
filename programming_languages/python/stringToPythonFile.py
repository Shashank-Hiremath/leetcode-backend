import sys

with open('/tmp/submissionFile.py', 'w') as submissionUpdatedFile:
    submissionUpdatedFile.write(sys.argv[1])