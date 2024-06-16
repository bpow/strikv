# Silly Trivial REST-like Internet-accessible Key-Value store

Written in an evening with very little residual golang knowledge from having written
something with it months ago... shows of the power of some golang libraries and of
being able to search for things on the net.

*USE AT YOUR OWN RISK* Did I mention how little golang experience I have?

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

# Example usage

With an appropriate docker and docker compose setup, start with:

```
docker compose up -d
```

Of course, you could `docker compose logs -f` to see some logging info...

Although the container listens to 8080 by default, the `docker-compose.yml` listens
on 9080 instead. This is clearly because I wanted to give and example of how to
do this, and not just because I already had something listening on 8080 when I wrote this...

Using the very-useful HTTPie to test things:

```
http http://localhost:9080/important-advice
```

Gives response:
```
HTTP/1.1 404 Not Found
Content-Length: 0
Date: Sun, 16 Jun 2024 15:29:23 GMT
```

That's because there's nothing there yet. We can easily write some JSON data:

```
http PUT http://localhost:9080/important-advice get-plenty-of-sleep=true
```

And we get:

```
HTTP/1.1 201 Created
Content-Length: 0
Date: Sun, 16 Jun 2024 15:31:41 GMT
```

Let's try that again:

```
http http://localhost:9080/important-advice
```

Now yields...

```
HTTP/1.1 200 OK
Content-Length: 31
Content-Type: text/plain; charset=utf-8
Date: Sun, 16 Jun 2024 15:31:54 GMT

{
    "get-plenty-of-sleep": "true"
}
```