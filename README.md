# TC-EXE

> exe runner for windows.

A Go module that provides a alternative to `exec.LookPath()` on Windows.

The following, relatively common approach to running external commands has a vulnerability on Windows:
```go
import "os/exec"

func goMod() error {
    cmd := exec.Command("go", "mod", "download)
    return cmd.Run()
}
```

Searching the current directory (surprising behavior) before searching folders listed in the PATH environment variable (expected behavior) seems to be intended in Go and unlikely to be changed: https://github.com/golang/go/issues/38736

Example use:
```go
import (
    "os/exec"
    tcexe "github.com/Timothee-Cardoso/tc-exe"
)

func goMod() error {
    goPath, err := tcexe.LookPath("git")

    if err != nil {
        return err
    }

    cmd := exec.Command(goPath, "mod", "download")
    return cmd.Run()
}
```
