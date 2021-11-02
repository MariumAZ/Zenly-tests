
In order to explain to the users the goal and the usage of our app, we need to edit the root.go file:

The first step when working with Protocol Buffers is to define the structure for the data you want to serialize in a .proto file.

protoc-gen-go needs to be in your shell path, i.e. one of the directories listed in the PATH environment variable, which is different from the Go path. You can test this by simply typing protoc-gen-go at the command line: If it says "command not found" (or similar) then it's not in your PATH.

```bash
export PATH=$PATH:$(go env GOPATH)/bin
```

The command used to gerate gopher.pb.go is : 

```bash
protoc -I=. --go_out=. ./gopher.proto 
```
