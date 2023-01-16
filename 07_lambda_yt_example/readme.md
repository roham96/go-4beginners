#
#
#
##  Make Lambda Yt Example App

### Before starting the web server run the following command:
```sh
go mod init github.com/username/lambda-yt-example
go mod tidy
```

### To start web server, run :
```sh
go get github.com/aws/aws-lambda-go/lambda
```

test: 
```sh 
aws iam create-role --role-name lambda-ex --assume-role-policy-document file://trust-policy.json
```
```sh 
go build main.go

ls
```
```sh 
zip function.zip main

ls
```
#
```sh 
aws lambda create-function --function-name lambda-yt-example \
 > --zip-file fileb://function.zip --handler main --runtime go1.x \
 > --role arn:aws:iam::3234234234:role/lambda-ex
```
#
extra mode:
```sh 
aws lambda invoke --function-name lambda-yt-example --cli-binary-format raw-in-base64-out --payload '{"what is your name?":"Reo","How old are you?":26}' output.txt

```

#
Enjoy it :)
