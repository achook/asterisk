# asterisk

## Description

- If the file is not provided or if the file does not exist, the program will not run the checking function but instead will return a non-zero exit code.
- In other cases (the file exists and is provided), the program will read the file.
- If the supplied file is not a valid JSON file the program will run the checking function that will return "true".
- Duplicate keys in the JSON file are allowed, as per the JSON specification. Only the last value of the key will be considered.
- Multiple "PolicyDocument/Statement" keys are allowed, as per the AWS::IAM::Role Policy document specification. 
- The program will check for an asterisk (*) in each "Resource" key value and return "false" if a single asterisk is found in all of them.
- If even a single "Resource" key value does not contain an asterisk, or the key is missing altogether, the program will return "true".
- The program will also return "true" if either the "PolicyDocument" or "Statement" key is missing.
- Other missing keys and values are allowed, as the program doesn't check the structure of the AWS::IAM::Role Policy document itself.

## Prerequisites
[Go](https://go.dev/) (any fairly recent version should work) 

## How to run
From the main directory, run the following command:
### Windows
````bash
go build -o asterisk.exe
.\asterisk.exe <path_to_file>
````

### Linux / MacOS
````bash
go build -o asterisk
./asterisk <path_to_file>
````

## How to test
From the main directory, run the following command:
````bash
go test
````
