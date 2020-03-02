Go HMMM Assembler Emulator
--------------------------
[Language specification][https://www.cs.hmc.edu/~cs5grad/cs5/hmmm/documentation/documentation.html]

Warnings

There should not be any tabs in the file 

If there are numbers they must be no tab between the number and the commands

There must not be any whitespace at the end of the file

Typing the command with no arguments creates a shell, however jumps aren't possible within it

$ hmmm

Typing the command and a file as an argument will run the commands in the file line by line

$ hmmm [file]

Installing
-----------

The hmmm binary should work for Linux, MacOS, etc. if not you may need to install the Go language/compiler/etc. Full details of installation and set up can be found on the Go language website. Once installed you have two options.
Compiling

go get && go build

This will create a binary for you. If you want to install it in the $GOPATH/bin folder you can run:

go install

On windows: try running hello.exe clicking it goes to shell 

to run use windows command prompt and type hmmm.exe [file to compile]
