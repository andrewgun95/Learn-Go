## Go commands

* go version
* go env
* go help
* go fmt
    * ./…
* go run
    * needs a file name, eg, go run main.go
    * go run <file name>
    * go run *.go
* go build
    * for an executable:
        1. builds the file
        2. reports errors, if any
        3. if there are no errors, it puts an executable into the current folder
    * for a package:
        1. builds the file
        2. reports errors, if any
        3. throws away binary
* go install
    * for an executable:
        1. compiles the program (builds it)
        2. names the executable 
            * mac: the folder name holding the code
            * windows: file name
        3. puts the executable in workspace / bin
            * $GOPATH / bin
    * for a package:
        1. compiles the package (builds it)
        2. puts the executable in workspace / pkg
            * $GOPATH / pkg
        3. makes it an archive file
* flags
    * -race