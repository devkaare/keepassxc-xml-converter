# KeePassXC XML to CSV convertion tool

## Introduction
This is a convertion tool which converts [KeePassXC](https://keepassxc.org/) XML exported files into CSV format.
Which is useful if you're like me, and exported your KeePassXC passwords to a XML file, and realised KeePassXC doesn't support importing XML files. 

## Requirements
A [Go compiler](https://go.dev/dl/) is **required** to run this program.

## Getting Started
1. Make sure the XML file you want to convert is inside the `tmp/` directory.
2. Run `go run main.go` to start the program.
3. Enter the name of your file (and press enter) **DON'T INCLUDE THE FILE EXTENSTION**
4. If everything went smoothly, a copy of your file (with the CSV filetype) should appear inside the `tmp/` directory.

Now you can finally import the new CSV file in KeePassXC if you wish to do so!
