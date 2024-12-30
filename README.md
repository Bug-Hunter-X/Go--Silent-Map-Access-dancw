# Go: Silent Map Access

This repository demonstrates a common, yet subtle, error in Go related to accessing keys in maps.  When you try to access a key that doesn't exist in a Go map, it doesn't result in a runtime panic. Instead, it returns the zero value for the map's value type. This can lead to unexpected behavior and difficult-to-debug issues.

The `bug.go` file shows an example of this behavior.  The `bugSolution.go` file provides a solution to mitigate this issue. 

## Reproducing the Bug

1. Clone this repository.
2. Navigate to the repository's directory.
3. Run `go run bug.go`.

You'll see that accessing a non-existent key simply returns 0, which might not be immediately apparent as an error.