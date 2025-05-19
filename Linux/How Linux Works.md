# How Linux Works

## Table of Contents

1. [BIG PICTURE](#1-big-picture)  
   1.1 [Levels and Layers of Abstraction in a Linux System](#11-levels-and-layers-of-abstraction-in-a-linux-system)  
   1.2 [Hardware: Understanding Main Memory](#12-hardware-understanding-main-memory)  
   1.3 [The Kernel](#13-the-kernel)  
   1.3.1 [Process Management](#131-process-management)  
   1.3.2 [Memory Management](#132-memory-management)  
   1.3.3 [Device Drivers and Management](#133-device-drivers-and-management)  
   1.3.4 [System Calls and Support](#134-system-calls-and-support)  
   1.4 [User Space](#14-user-space)  
   1.5 [Users](#15-users)  

2. [BASIC COMMANDS AND DIRECTORY HIERARCHY](#2-basic-commands-and-directory-hierarchy)  
   2.1 [Using the Shell](#using-the-shell)  
   2.2 [Basic Commands](#basic-commands)  
   2.3 [Navigating Directories](#navigating-directories)  
   2.4 [Intermediate Commands](#intermediate-commands)  
   2.5 [Changing Your Password and Shell](#changing-your-password-and-shell)  
   2.6 [Dot files](#dot-files)  
   2.7 [Environment and Shell Variables](#environment-and-shell-variables)  
   2.8 [The Command Path](#the-command-path)  
   2.9 [Special Characters](#special-characters)  
   2.10 [Command-Line Editing](#command-line-editing)  
   2.11 [Text Editors](#text-editors)  
   2.12 [Getting Help](#getting-help)  
   2.13 [Shell Input and Output](#shell-input-and-output)  
   2.14 [Understanding Error Messages](#understanding-error-messages)  
   2.15 [Listing and Manipulating Processes](#listing-and-manipulating-processes)  
   2.16 [File Modes and Permissions](#file-modes-and-permissions)  
   2.17 [Archiving and Compressing Files](#archiving-and-compressing-files)  

---

## 1. BIG PICTURE

### 1.1 Levels and Layers of Abstraction in a Linux System  
hardware → kernel → User Processes  

**User processes:**  
- GUI  
- Servers  
- Shell  
- etc.  

**Linux Kernel:**  
- System Calls  
- Process Management  
- Memory Management  
- Device Drivers  

**Hardware:**  
- Processor (CPU)  
- Main Memory  
- Disks  
- Network Ports  

User Processes — the running programs that the kernel manages — collectively make up the system's upper abstraction level, called **user space**.

**Kernel space VS user space**

---

### 1.2 Hardware: Understanding Main Memory

---

### 1.3 The Kernel

The kernel is in charge of managing tasks for four general system areas:  
- **Processes:** kernel determines which processes can use CPU.  
- **Memory:** kernel keeps track of all memory.  
- **Device drivers:** kernel is an interface between processes and hardware.  
- **System calls and support:** processes use system calls to communicate with the kernel.

---

#### 1.3.1 Process Management

Process management describes starting, pausing, resuming, scheduling, and terminating processes.  

- **Concurrency vs Parallelism**  
- **Context switch:** act of one process giving up control of the CPU to another.  
- **Time slice:** minimal time needed for a process to complete minimal significant computation between context switches.  
- Kernel runs between process time slices during a context switch.

---

#### 1.3.2 Memory Management

- **Memory management unit (MMU):** enables 'virtual memory'.

---

#### 1.3.3 Device Drivers and Management

Device drivers are usually accessible only to the kernel.

---

#### 1.3.4 System Calls and Support

System calls or syscalls perform tasks that user processes cannot do well or at all.  

- **fork():** when called by a process, the kernel creates a nearly identical copy of the process.  
- **exec():** kernel loads and starts a program, replacing the current process.  

The kernel also supports user processes with features other than system calls, such as pseudodevices.

---

### 1.4 User Space

User space is the main memory allocated by the kernel for user processes. Despite user processes appearing the same from the kernel's point of view, we can divide them into different levels.

---

### 1.5 Users

The Linux kernel supports the traditional concept of a Unix user.  

- **User:** an entity that can run processes and own files.  
- **Groups:** sets of users. The primary goal of groups is to share file access among group members.

---

## 2. BASIC COMMANDS AND DIRECTORY HIERARCHY

The **Bourne Shell**: `/bin/sh` — standard shell developed at Bell Labs for early Unix versions.  
Every Unix system needs a version of `/bin/sh` to work.  
The **Bourne-again shell (bash)** is an enhanced version of the Bourne shell. On Linux, `/bin/sh` is usually a link to bash.

---

### Using the Shell

#### 2.1 The Shell Window - Terminal

#### 2.2 cat

#### 2.3 Standard Input and Standard Output

- **CTRL-D vs CTRL-C difference**

**Standard I/O streams:**  
- Standard input  
- Standard output  
- Standard error

---

### 2.4 Basic Commands

#### 2.4.1 ls

#### 2.4.2 cp

#### 2.4.3 mv

#### 2.4.4 touch

#### 2.4.5 rm

#### 2.4.6 echo

---

### 2.5 Navigating Directories

- `/` — root directory, absolute path  
- `..` — parent of a directory  
- `.` — current directory

#### 2.5.1 cd

#### 2.5.2 mkdir

#### 2.5.3 rmdir

#### 2.5.4 Shell Globbing ("Wildcards")

- Check connection to regexp

---

### 2.6 Intermediate Commands

#### 2.6.1 grep  
Prints lines from a file or input stream that match an expression.

#### 2.6.2 less  
Terminal pager program used to view file contents one page at a time.

#### 2.6.3 pwd  
*Symbolic links may obscure true full path, so use* `pwd -P` *to eliminate this confusion.*

#### 2.6.4 diff  
Shows differences between two text files.

#### 2.6.5 file  
Checks file format.

#### 2.6.6 find and locate

#### 2.6.7 head and tail

#### 2.6.8 sort

---

### 2.7 Changing Your Password and Shell

Use `passwd` command to change your password.

---

### 2.8 Dot files

Files starting with `.` are hidden (e.g., `.dir`, `.filename`).

---

### 2.9 Environment and Shell Variables

- Shell variables: temporary string values accessible from the programs the shell runs but not from commands run.
- Shell variables are local to the current terminal session and not visible to child processes unless exported.
- Use `export` to assign shell var value to environment var, e.g., `export myvar="good"` creates an environment variable.

---

### 2.10 The Command Path

- `PATH` is a special environment variable containing a list of system directories that the shell searches when locating a command.
- Append a directory: `export PATH=$PATH:/home/user/mybin`
- Prepend a directory: `export PATH=/home/user/mybin:$PATH`

---

### 2.11 Special Characters

---

### 2.12 Command-Line Editing

---

### 2.13 Text Editors

---

### 2.14 Getting Help

- `man` — manual pages  
- `man -k keyword` — search manuals  
- `command --help`  
- `info command`  
- `man 5 passwd` — section 5 for file formats

---

### 2.15 Shell Input and Output

- Redirect output to file: `command > file`  
- Append output to file: `command >> file`  
- Pipe output of one command to another: `command1 | command2`

#### 2.15.1 Standard Error

- Additional output stream for diagnostics/debugging.
- Redirect stderr: `ls /fffff > f 2> e`  
- `2` is the stream ID for standard error  
- Redirect stderr to stdout: `ls /fffff > f 2>&1`

#### 2.15.2 Standard Input Redirection

- Channel a file to a program’s standard input: `head < /proc/cpuinfo`

---

### 2.16 Understanding Error Messages

#### 2.16.1 Anatomy of a Unix Error Message

#### 2.16.2 Common Errors

---

### 2.17 Listing and Manipulating Processes

- `ps` — lists running processes  
- **PID:** process ID  
- **TTY:** terminal device where process runs  
- **STAT:** process status  
- **TIME:** CPU time used by process  
- **COMMAND:** command used to run the program (can change)

#### 2.17.1 ps Command Options

- `ps x` — show all your running processes  
- `ps ax` — show all processes on the system, not just those you own  
- `ps u` — detailed info on processes  
- `ps w` — show full command names

#### 2.17.2 Process Termination

- `kill pid` — sends TERM signal  
- `kill -STOP pid`  
- `kill -CONT pid`

#### 2.17.3 Job Control

Shell supports job control, allowing signals `TSTP` and `CONT` to be sent via commands and keystrokes.  
Commands: `fg`, `bg`  
Tools: `screen`, `tmux`

#### 2.17.4 Background Processes

- Run process in background with `&`:  
  `nohup command &`  
  `gunzip file.gz &`

---

### 2.18 File Modes and Permissions

- Every Unix file has permissions for reading, writing, and running.  
- `ls -l` shows permissions. Example:  
  `-rw-r--r-- juser somegroup 7041 Mar 26 19:34 endnotes.html`  
- `setuid` explained later.

#### 2.18.1 Modifying Permissions

- `umask 077` — default permission mask for new files.

#### 2.18.2 Working with Symbolic Links

- Create symlink: `ln -s target linkname`  
- View symlinks: `ls -l`  
- Get real path: `readlink` / `realpath link_to_file`  
- Remove symlink: `rm` / `unlink link_name`  
- Test if symlink: `test -L file`  
- *Hard links* also exist.

---

### 2.19 Archiving and Compressing Files

#### 2.19.1 gzip

- Compress: `gzip file`  
- Uncompress: `gunzip file.gz`

#### 2.19.2 tar

- Create archive: `tar cvf archive.tar file1 file2 ...`  
- Extract archive: `tar xvf archive.tar`  
- Options: `p` and `t` modes

#### 2.19.3 Compressed Archives

#### 2.19.4 zcat

- Equivalent to `gunzip -dc` which decompresses and sends output to stdout.  
- Example: `zcat file.tar.gz | tar xvf -`  
- `z` option automatically invokes gzip on archive.

#### 2.19.5 Other Compression Utilities

---

## Linux Directory Hierarchy Essentials

*(Not included in the excerpt)*

