# Mini C Compiler
A Project for CUHK **CSCI3120 Compiler Construction**

This is a mini C compiler that aims to compile simple C language into intermidate code(i.e. Quadruples) from scratch  
Notes that this compiler do not fully follow ISO C standard

## Usage
The main entrance is `main.go`, and you can run it by:
```
go run main.go -s {source file path} -o {destination directory}
```
To build the source code into binary:
```
go build main.go
./main -s {source file path} -o {destination directory}
```