package auth

import "errors"

var (
	/* ErrNotInitialized indicates the auth package has not been configured. */
	ErrNotInitialized = errors.New("auth package not initialized")
	/* ErrDatabaseUnavailable indicates the configured storage backend is unavailable. */
	ErrDatabaseUnavailable = errors.New("database connection unavailable")
	/* ErrInvalidToken indicates a JWT could not be parsed or validated. */
	ErrInvalidToken = errors.New("invalid jwt token")
	/* ErrTokenExpired indicates a JWT exceeded its configured expiration time. */
	ErrTokenExpired = errors.New("jwt token expired")
	/* ErrOTPExpired indicates an OTP has expired and can no longer be used. */
	ErrOTPExpired = errors.New("otp expired")
	/* ErrOTPNotFound indicates no OTP record exists for a user. */
	ErrOTPNotFound = errors.New("otp not found")
	/* ErrInvalidOTP indicates the supplied OTP code does not match the stored value. */
	ErrInvalidOTP = errors.New("invalid otp code")
	/* ErrUserNotFound indicates the requested user does not exist. */
	ErrUserNotFound = errors.New("user not found")
	/* ErrInvalidCredentials indicates username/password or role validation failed. */
	ErrInvalidCredentials = errors.New("invalid credentials")
	/* ErrSMTPNotInitialized indicates no email sender has been configured. */
	ErrSMTPNotInitialized = errors.New("smtp not initialized")
	/* ErrJWTSecretMissing indicates the JWT signing secret is not configured. */
	ErrJWTSecretMissing = errors.New("jwt secret not initialized")
	/* ErrInvalidInput indicates a supplied value violates expected constraints. */
	ErrInvalidInput = errors.New("invalid input provided")
	/* ErrInvalidEmail indicates a supplied email address does not parse correctly. */
	ErrInvalidEmail = errors.New("invalid email format")
	/* ErrEmptyInput indicates a required field has an empty value. */
	ErrEmptyInput = errors.New("required field cannot be empty")
	/* ErrRateLimitExceeded indicates a request is over the allowed threshold. */
	ErrRateLimitExceeded = errors.New("rate limit exceeded")
	/* ErrRefreshTokenInvalid indicates the refresh token does not exist or is malformed. */
	ErrRefreshTokenInvalid = errors.New("invalid refresh token")
	/* ErrRefreshTokenRevoked indicates the refresh token was explicitly revoked. */
	ErrRefreshTokenRevoked = errors.New("refresh token has been revoked")
	/* ErrRefreshTokenExpired indicates the refresh token has expired. */
	ErrRefreshTokenExpired = errors.New("refresh token has expired")
	/* ErrOAuthNotInitialized indicates Google OAuth has not been configured. */
	ErrOAuthNotInitialized = errors.New("oauth not initialized")
	/* ErrOAuthExchangeFailed indicates OAuth token exchange failed. */
	ErrOAuthExchangeFailed = errors.New("failed to exchange oauth code")
	/* ErrOAuthProfileFetchFailed indicates the OAuth userinfo endpoint failed or returned an invalid response. */
	ErrOAuthProfileFetchFailed = errors.New("failed to fetch user profile from provider")
	/* ErrRedisUnavailable indicates the configured Redis client is unusable. */
	ErrRedisUnavailable = errors.New("redis connection unavailable")
	/* ErrRateLimitBackendDown indicates the Redis-backed rate limiter backend is unavailable. */
	ErrRateLimitBackendDown = errors.New("rate limit backend is down")
)
