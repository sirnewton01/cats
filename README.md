Secure CAT (cats)
=================

This is an enhanced version of the venerable cat command from Unix. You can use it to read local or remote files over SSH.

```
cats nobody@supertxt.net:browsing.s.txt
```

If the file provided is remote then it will connect using ssh to read the file. Local files are opened with the regular cat command. It works much like other [sshla](http://supertxt.net/whats-sshla.s.txt) commands, such as scp or git.

There are special reflective capabilities so that cats works with well with [SuperTXT browser](http://supertxt.net/browsing.s.txt) in conjunction with SuperTXT documents.


Installation
============

Cats is written in Go. With a recent Go SDK installed you can install cats like this.

```
cd cats
go install ./...
```

Be sure to add the Go binaries directory to your PATH.


Reflection
==========

The cats command supports two kinds of command-line reflection through the '--srefl' option. It's often difficult for an external program to determine what sort of output to expect from a command-line, but commands can often determine this for themselves. Agent programs like the shell could do more with program output if they are aware of the type of data going to the output, such as syntax highlighting, or even rendering images on the screen. You can ask cats to output the file extension of a command without actually running it like this:

``` .sh
cats --srefl=extension some/path/to/file.jpg
```

If cats can determine the file extension from the path provided then it prints it to stdout.

```
.jpg
```

Another kind of reflection is available for redirection. An agent program (eg. a shell or browser) may not know how to rewrite the command line arguments for a program given relative path or other kind of link fragment shown in its output. A program like cats can reconstruct its own command-line options and arguments using the command that was run before, with the redirect option, and the link fragment from stdin like this:

``` .sh
echo 'browsing.s.txt' | cats --srefl=redirect supertxt.net:00-intro.s.txt
```

Cats then outputs to stdout what the new command should be for this redirect.

``` .sh
cats supertxt.net:browsing.s.txt
```

You can read more about reflection and the srefl option on the superxt.net website in the [browsing](http://supertxt.net/browsing.s.txt) section.

Legal
-----

See the [license](LICENSE.txt) for more information.
