# go-GridScorer

go-GridScorer takes in a linear array and turns into an n matrix; based on the input it will calculate the highest score from itself and the surrounding location. This will then return a string result with the following format (x, y, total value)

Example array: `[4,2,3,2,0,1,2,2,1,3,0,2,2,0,1,5]`

Example matrix from the above array:

```
4 | 2 | 3 | 2

0 | 1 | 2 | 2

1 | 3 | 0 | 2

2 | 0 | 1 | 5
```

Example of input:

```
count_of_high_scores = 2

row_length = 4

array = [4,2,3,2,0,1,2,2,1,3,0,2,2,0,1,5]
```

Example of output:

`(1, 2, 17)(1, 1, 16)`

## Setup

Make sure to have go installed - https://go.dev/dl/

You can use your IDE of your choice, but I will provide the setup for Visual Studio Code:

1) Install the "Go" extension (this aids with development and the ability to build the project)
    - VS Code may prompt you to install additional tools required for Go development, such as gopls (the Go language server).
2) Open the Command Palette by pressing Ctrl+Shift+P and search for "Go: Install/Update Tools". 
    - Select this option to open a list of Go tools that can be installed. Ensure that "gopls" (Go Language Server) and "gotests" are selected, as they are necessary for testing.

## Running the Go Code

You can run the Go program directly from the VS Code by:
- Right-clicking in the editor window and selecting Run Go File.
- Or pressing Ctrl+F5 (which runs without debugging).
- Alternatively, you can press F5 to start debugging the Go program.

## Using the Terminal

You can also run Go programs using the integrated terminal in VS Code (as long as you are in the directory containing the file):
- `go build`
- `go test` 