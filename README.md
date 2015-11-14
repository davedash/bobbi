# Bobbi

An agent that sits in a server that takes orders from the outside and executes them.

At least that'll be our intent.

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
