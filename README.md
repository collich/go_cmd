# GO_CMD_Development
Project for Go Command-Line Development.

## Projects
- ### go-start-proj
    - To bootstrap an instant go project with file structure.
        - `CMD` folder with `main.go`
        - `tests` folder to run `test.go`
        - `bin` folder to hold exe or binary files
        - `internal` folder to hold other types of file/folder
    - Run `go build cmd/main.go` for linux or `env GOOS=<target-OS> GOARCH=<target-architecture> go build cmd/main.go`
        ```js
        Target OS - Target Platform
        ==================
        android 	arm
        darwin 	    386
        darwin 	    amd64
        darwin 	    arm
        darwin 	    arm64
        dragonfly 	amd64
        freebsd 	386
        freebsd 	amd64
        freebsd 	arm
        linux 	    386
        linux 	    amd64
        linux 	    arm
        linux 	    arm64
        linux 	    ppc64
        linux 	    ppc64le
        linux 	    mips
        linux 	    mipsle
        linux 	    mips64
        linux 	    mips64le
        netbsd 	    386
        netbsd 	    amd64
        netbsd 	    arm
        openbsd 	386
        openbsd 	amd64
        openbsd 	arm
        plan9 	    386
        plan9 	    amd64
        solaris 	amd64
        windows 	386
        windows 	amd64
        ```
