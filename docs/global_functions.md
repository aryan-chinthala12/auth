# Configuration pattern

Before you use this library, there are a few functions that should be called globally.

Luckily, most of these aren’t functions that should be configured by everyone, and are not recommended for all users to configure, unless the person has good technical understanding of what they are doing.

The functions are these:

---

1. **auth.Init()**

* Purpose: Initializes the library, sets all global variables, and validates the database schemas.
* Usage Notes:

  * Should be called globally **before using any other function**.
  * Failing to call this first may result in unexpected errors.
* Recommended For: All users. This is mandatory.

---

2. **auth.PepperInit(pep string)**

* Purpose: Sets the global secret pepper used in password hashing.
* Usage Notes:

  * Should be called **once at startup**, immediately after `auth.Init()`.
  * Only needed if you want to enable the “pepper” feature.
  * The pepper is stored **in memory only**, never in the database.
  * Losing or changing the pepper will make all existing password hashes invalid.
* Recommended For: Advanced users who understand the risks and want extra server-side secret protection.
* Note: Not recommended for everyone.

---

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

* Recommended For: Advanced users who want to tweak performance and security settings.
* Note: Not recommended for everyone.
