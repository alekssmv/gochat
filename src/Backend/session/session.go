package session

import (
    "github.com/gorilla/sessions"
)

var (
    // Key must be 16, 24, or 32 bytes long for AES-128, AES-192, or AES-256
    Key   = []byte("super-secret-key")
    Store = sessions.NewCookieStore(Key)
    Name  = "gochat-session"
)
