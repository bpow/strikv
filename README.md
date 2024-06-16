# Silly Trivial REST-like Internet-accessible Key-Value store

Written in an evening with very little residual golang knowledge from having written
something with it months ago... shows of the power of some golang libraries and of
being able to search for things on the net.

# What it does

Runs a simple http server using golang standard library. PUT requests will treat the url
path as a key and store the body content in a [Badger](https://github.com/dgraph-io/badger)
database. GET requests at that path will return that content, or error 404 if nothing
was ever stored there.

## parameters that can be provided

As environment variables:

- `STRIKV_PATH` - the location where the Badger database will be stored
- `STRIKV_PORT` - port on which to listen (use `0` to get an available port selected
  for you)

# What it doesn't do

Much of anything else...

The following are left as exercises to the reader:

1. Authentication or any sort of access control-- write something yourself or maybe put
   this behind something like an oauth procy...
2. Much by way of useful error messages. Lots could probably be included here.
3. Prevention of overwriting of old values
4. Logging of usage
5. Rate limiting

Golly, readers will be so healthy after all of those exercises!
