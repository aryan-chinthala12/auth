# Configuration pattern

This package does not rely on global initialization functions. The current API is constructor-based and is configured through the options passed to `auth.New(...)`.

Use the library like this:

```go
ctx := context.Background()
authClient, err := auth.New(ctx,
    auth.WithStorage(storage),
    auth.WithPepper([]byte("super-secret-pepper")),
    auth.WithJWT([]byte("jwt-secret"), 24*time.Hour),
    auth.WithRefreshToken(7*24*time.Hour, 32),
    auth.WithOTP(6, 5*time.Minute),
)
if err != nil {
    panic(err)
}
defer authClient.Close()
```

## Common configuration options

- `auth.WithStorage(storage)`
  - Required for storage-backed operations such as user registration, login, OTPs, and refresh tokens.

- `auth.WithPepper([]byte("..."))`
  - Recommended when OTPs are enabled or when you want extra server-side secret protection.

- `auth.WithJWT(secret, expiry)`
  - Configures the JWT signing secret and access-token lifetime.

- `auth.WithRefreshToken(expiry, length)`
  - Configures refresh-token lifetime and token size.

- `auth.WithOTP(length, expiry)`
  - Enables OTP support and defines its length and validity window.

- `auth.WithRedis(client)`
  - Enables Redis-backed rate limiting and token caching.

## Hash tuning

If you need to tune the Argon2id parameters, call the instance method on the configured `*auth.Auth` value before using the library for password hashing:

```go
if err := authClient.DefaultSaltParameters(3, 64*1024, 1, 32); err != nil {
    panic(err)
}
```

This is not a global function; it is configured per `Auth` instance.
