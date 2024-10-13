# KeePassXC XML to CSV convertion tool
## Requierments
This tool ONLY supports ONE type of STRUCTURE for your KeePassXC XML file:

    Root/ (SUPPORTED STRUCTURE)
        Group1/
            Entry1
            Entry2
            ... (It supports an infinite amount of entries)
        Group2/
            Entry1
            Entry2
            ... (It supports an infinite amount of entries)

So basically, you can only have 1 nested layer of groups (inside root). BUT, If you have root and multiple nested layers of groups inside other groups, that doesn't work e.g:

    Root/ (UNSUPPORTED STRUCTURE)
        Group1/
            Entry1
            Entry2
                Group2InsideGroup1/
                    Entry1
                    Entry2
        Group2/
            Entry1
            Entry2
                Group3InsideGroup2/
                    Entry1
                    Entry2

Having all your entries inside the root itself also doesn't work, e.g:

    Root/ (UNSUPPORTED STRUCTURE)
        Entry1
        Entry2
        Entry3
        Entry4

## How to use
1. Open main.go
2. Edit line 132 and change "example" to your XML files name and save. e.g "old-passwords-file" instead of "example"
(don't include the file extension when typing the name of your file, e.g don't add .xml at the end)
3. Ensure your XML file is inside the tmp directory
(you can delete the example files if you'd like btw)
5. Check if you have go installed on your system using the command: go version
(if go is not installed, please install it on your system https://go.dev/dl/)
6. In the directory containing the main.go file, run the command: go run main.go
