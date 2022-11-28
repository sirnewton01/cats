Secure CAT (scat)
=================

This is an enhanced version of the venerable cat command from Unix. You can use it to read local or remote files over SSH.

```
scat nobody@supertxt.net:browsing.s.txt
```

If the file provided is remote then it will connect using ssh to read the file. Local files are opened with the regular cat command. It works much like other [sshla](http://supertxt.net/whats-sshla.s.txt) commands, such as scp or git.

There are special reflective capabilities so that scat works with well with [SuperTXT browser](http://supertxt.net/browsing.s.txt) in conjunction with SuperTXT documents.


Installation
============

Scat is written in Go. With a recent Go SDK installed you can install scat like this.

```
cd scat
go install ./...
```

Be sure to add the Go binaries directory to your PATH.

Legal
-----

See the [license](LICENSE.txt) for more information.
