# wc-tool

```txt
NAME:
   wc-tool - count bytes in a file

USAGE:
   wc-tool [OPTIONS] FILE

DESCRIPTION:
   A command-line tool for counting various metrics in a file.

COMMANDS:
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --bytes, -c  Print the byte count
   --lines, -l  Print the line count
   --words, -w  Print the words count
   --chars, -m  Print the chars count
   --help, -h   show help
```

**Results printed in order: `lines words chars bytes`**

**Currently supports a single file.**

**If no stdin or file given, the text can be copy pasted.**
```bash
./wc-tool

Professor Michael S. Hart was the originator of the Project
Gutenberg™ concept of a library of electronic works that could be
freely shared with anyone. For forty years, he produced and
distributed Project Gutenberg™ eBooks with only a loose network of
volunteer support.
```
`CTRL+D` to stop input and get the results.

**Usage**

<a href="https://asciinema.org/a/608819?t=10" target="_blank"><img src="https://asciinema.org/a/608819.svg?t=10" width=500px/></a>
