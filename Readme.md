# Mini C Compiler
A Project for CUHK **CSCI3120 Compiler Construction**

This is a mini C compiler that aims to compile simple C language into intermidate code(i.e. Quadruples) from scratch  
Notes that this compiler do not fully follow ISO C standard

## Usage
The main entrance is inside the directory `\main`, and you can run the `main.go` by:
```
cd main
go run main.go -s {scource file path}
```
To build the source code into binary:
```
cd main
go build main.go -o {destination}
```