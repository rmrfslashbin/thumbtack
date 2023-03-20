[![Go Reference](https://pkg.go.dev/badge/github.com/rmrfslashbin/thumbtack.svg)](https://pkg.go.dev/github.com/rmrfslashbin/thumbtack)

# thumbtack
Thumbtack is a Go client for https://pinboard.in.

## Pinboard
Find out more about Pinboard at https://pinboard.in/about/.
```
Pinboard is a fast bookmarking website for privacy-minded people. It helps you keep track of things you find online and manage your tab clutter.
```

## Pinboard API
The Pinboard API is a RESTful API that allows you to interact with a Pinboard account. The API is documented at https://pinboard.in/api/.

## API Coverage
This client supports the following API endpoints (https://api.pinboard.in/v1/):
- update
    - posts/update
- posts
    - posts/add
    - posts/delete
    - posts/get
    - posts/dates
    - posts/recent
    - posts/all
    - posts/suggest
- tags
    - tags/get
    - tags/delete
    - tags/rename
- user
    - user/secret
- notes
    - notes/list
    - notes/ID

## CLI
This repo provides a CLI as a reference implementation of the client. The CLI is not intended to be a full-featured client, but rather a simple example of how to use the client. Most of the CLI's functionality is implemented in the `cmd` package.

### Building
To build the CLI, run `make build`. This will build the CLI and place it in the `bin` directory.

## Pinboard Authentication and User Tokens
This client only supports `API authentication tokens` for authentication. The client does not support `Regular HTTP Auth`. Users can find their API token on their settings page: https://pinboard.in/settings/password.

As stated by the Pinboard API documentation:
```
The Pinboard v1 API requires you to use HTTPS. There are two ways to authenticate:

Regular HTTP Auth:
https://user:password@api.pinboard.in/v1/method

API authentication tokens:
https://api.pinboard.in/v1/method?auth_token=user:NNNNNN

An authentication token is a short opaque identifier in the form "username:TOKEN".

Users can find their API token on their settings page. They can request a new token at any time; this will invalidate their previous API token.

Any third-party sites making API requests on behalf of Pinboard users from an outside server MUST use this (API authentication tokens) authentication method instead of storing the user's password. Violators will be blocked from using the API.
```

## Rate Limiting
This client will return http status codes and error codes but does not implement any backoff logic. It is up to the user to implement backoff logic.

As stated by the Pinboard API documentation:
```
API requests are limited to one call per user every three seconds, except for the following:

posts/all - once every five minutes
posts/recent - once every minute
If you need to make unusually heavy use of the API, please consider discussing it with me first, to avoid unhappiness.

Make sure your API clients check for 429 Too Many Requests server errors and back off appropriately. If possible, keep doubling the interval between requests until you stop receiving errors.
```

## Error Handling
This client will return http status codes and error codes as derived from the Pinboard API documentation. It is up to the user to handle these errors. The client provides a number of error types that can be used to determine the type of error that was returned.

## User Agent
This client provides a default user agent that consists of the repo/package name and the version of the client. The user agent can be overridden by the user/client implementation. The user agent is used to identify the client to the Pinboard API.

From the Pinboard API documentation (https://pinboard.in/api/v2/overview)
```
Don't mislead the API server about who you are (by using a fake User-Agent string, for example). If you develop app, please register and use an app identifier, and update it with each majore release.
```

## Logging
This client uses https://github.com/rs/zerolog for logging. If desired, Zerolog output can be effectively silenced by setting the log level to `zerolog.SetGlobalLevel(zerolog.PanicLevel)`.