# GO-GETSTREAM

[![wercker status](https://app.wercker.com/status/adc2bf440cb3e5b8f4fa3abf9244624d/s/master "wercker status")](https://app.wercker.com/project/byKey/adc2bf440cb3e5b8f4fa3abf9244624d)
[![godoc](https://godoc.org/github.com/mrhenry/go-getstream?status.svg)](https://godoc.org/github.com/mrhenry/go-getstream)
[![codecov](https://codecov.io/gh/mrhenry/go-getstream/branch/master/graph/badge.svg)](https://codecov.io/gh/mrhenry/go-getstream)

Golang pkg for [getstream.io](getstream.io). The goal of this package is to provide server-side support for getstream and to generate client-side tokens.

### Supported :
- Flat Feed
  - Add
  - Add Multiple
  - Remove
  - Remove By ForeignID
  - List
  - Follow
  - UnFollow
  - Followers
  - Following

- Aggregated Feed
  - Add
  - Add Multiple
  - Remove
  - Remove By ForeignID
  - List
  - Follow
  - UnFollow
  - Following

- Notification Feed
  - Add
  - Add Multiple
  - Remove
  - Remove By ForeignID
  - List
  - Follow
  - UnFollow
  - Following
  - Mark Read
  - Mark Seen

- Untested :
  - Client Side Scope Tokens

### Structure :
- Follows getstream API standards for all request payloads
- `data` : Statically typed payloads as `json.RawMessage`
- `metadata` : Top-level key/value pairs

You can/should use `data` to send golang structures through getstream. This will give you the benefit of golang's static type system.
If you can't know the contents of an Activity you can use metadata which is a `map[string]string`, encoding to json will move this values to the top-level. This means that keys which conflict with standard getstream keys will be overwritten. The benefit of this structure is that these key/value pairs will be exposes to getstream internals such as ranking,...

### Design Choices :

- Flat / Aggregated / Notification Feeds have separate structures and methods to prevent the impact of future getstream changes. If two types of feeds grow farther apart this can be incorporated in this client without breaking everything.

### Credits :

This pkg started out as a fork from https://github.com/hyperworks/go-getstream and still borrows snippets (token & errors) from the original. I decided to make this a separate repo entirely because drastic changes were made to the interface and internal workings.
