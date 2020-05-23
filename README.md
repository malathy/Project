# Project
# This project is to build a REST API In GOLang 
# First  we have to download the go programming and WVC(Windows Visualstudio code) for writing required codes
# Created the go-res-api folder in src/github.com/malathy/
# IN CODE , we start to import requied packges,here am using encoding/json.fmt,log,//"net/http" and os
# created struct for url and products
# In main function,we created instances for  these struct(url and product) and gave required information
# By using Marshall (For encoding json data we can use Marshall) encode jason data in to another variable(byArray)
# print that variable in string format
# Code for run go file is "go run filename.go"(go run assignment.go)
# Next printed these json response in text file using following code
# creating  response text file
# file,e:=os.OpenFile("Response_new.text",os.O_CREATE|os.O_WRONLY|os.O_APPEND,0666)
# if e!=nil{
# }
# //Writing into text file
# log.SetOutput(file)
# log.PrintLn(string(byteArray))
# }
