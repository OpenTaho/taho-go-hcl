# Taho HCL Go Project

A Go Module for working with HCL.

## Why this Module Exists

In `taho@0.0.x`, we used file based parsing to access comments because at the
time, the [Hashicorp HCLv2 Toolkit] package did not provide a way of working
with comments or whitespace. The code in `taho@0.0.x` is a mess with some parts
of the Hcl files parsed using the toolkit while other parts are parsed using
rudimentary file operations. Content is often parsed using the HCLv2 toolkit then
individual blocks are written out to the file system so they can be reparsed
usin file system based techniques.

This module was created to allow fugure versions of Taho to work with HCL
through the abstraction layer provided by it's interfaces and avoid performing
file based parsing.

## What this Module Provides

This module provides a method to read HCL files and directories and returns an
interface that allows callars to interact with the content in an object oriented
way.

Various implementations of the interface defined by this module are planned.
The [taho-go-hcl-thin] implementation aims to parse content without using the
[Hashicorp HCLv2 Toolkit] while the [taho-go-hcl-hcl-v2] will be based on the
`taho@0.0.x`.

## How this Module Works

The `hcl` package this module provides is built around the concept of
representing documents as a collection of **HclNodes**. After a file or
directory of files has been read information is provided as a collection of
nodes with subnodes.

Reading a file results in one node to represent the file and child nodes
representing the contents within the file. When needed, **HclNodes** have sub
**HclNodes**. Each **HclNodes** also has a type.

As a file is read nodes become more complex based on processing rules.

### Empty Directories or Files

The least complex possible HCL content is an empty file or empty directory.
Reading an empty file would result in one node where the value of the node is
the file's simple name. This hypothetical node would have no subnodes.

## Space, Line Feed, and Tabs

HCL files may contain space, line feed, and/or tab charactors such that the
simplest possible legal HCL file would be one that has one space or one new
line.

The following command will create some very simple but syntactically legal hcl
files.

```zsh
# [0] - File type node with value of "000.tfvars"
touch 000.tfvars

# [0] ---- File type node with value of "001.tfvars"
# [0][0] - Space type node with value of " "
printf ' ' > 001.tfvars

# [0] ---- File type node with value of "002.tfvars"
# [0][0] - Space type node with value of "\t "
printf "\t " > 002.tfvars

# [0] ---- File type node with value of "003.tfvars"
# [0][0] - Space type node with value of " "
# [0][1] - Rune type node with value of "\n"
printf " \n" > 003.tfvars

# [0] ---- File type node with value of "004.tfvars"
# [0][0] - Space type node with value of "\t "
# [0][1] - Token type node with value of "\n"
printf "\t \n" > 004.tfvars
```

### HclLineFeed

<!-- links -->
[Hashicorp HCLv2 Toolkit]: https://pkg.go.dev/github.com/hashicorp/hcl/v2
