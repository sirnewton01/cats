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

Legal
-----

See the [license](LICENSE.txt) for more information.
