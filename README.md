# Sprite

Sprite is a super-lightweight HTTP server, meant to be used when Apache (or even
nginx) is overly cumbersome. It is distributed as a single binary with a single
example config file. The example config file will be useable with modification
by most users.

All that is required to run, then, is to put the binary somewhere on your
`PATH`, put the config file in a place like `/etc/sprite.conf`, and start the
server by typing `sprite -f /etc/sprite.conf`. That's all there is to deploying
your own web server.

# Status

Sprite is nowhere near complete. Most of the description above is a goal, not a
current status. Check back in for updates, though.
