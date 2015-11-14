# Bobbi

An agent that sits in a server that takes orders from the outside and executes them.

The idea is you might be running ChatOps via Hubot or something, and inevitably you'll want to run some crazy commands on some other host.  E.g. you might want to reboot a host from hubot.

You could have hubot SSH to the machine and run the right commands, but that involves creating an account for Hubot, and setting up SSH keys.

My proposal is to have Bobbi running on any server you'd like.  You can make HTTP requests from your Hubot scripts to Bobbi.  Bobbi will happily run them on the server it's on.  This can all be locked down via Security Groups.  

Note, I've undoubtedly made some mistakes or silly assumptions.  Feel free to give me feedback.

## TODO

- Bobbi will be an HTTP server.  If you'd rather use HTTPS, you can front it with `nginx`.  If you want authentication, also look at `nginx`.
- It'll run on port 8000 by default, but you can configure that.
- Listens to end-points defined in a file e.g. `/reboot`
- Runs a task, e.g. `reboot` as a specific user (e.g. `root`).
  - Task run in Go routine
  - POST to that endpoint gives a 202 Accepted and maybe some JSON status
  - Repeated POST will just give a 200 with current status, until the job is done
- GET to the endpoint will give you the current status: not running, running, success, fail
- Maybe list the last few lines of output.
