# Summer (is here)

The plan is to build something that loosely follows jsonrpc 2.0. That still allows for extensions in all directions.

Since im a huge Gin fanboy, it will be modelled as a gin.HandlerFunc that you mount where you want it.

## TODO

* Should each request spawn it's own registry with the methods that its identity have access to?
  > and then figure out a way to cache it for http-requests.
