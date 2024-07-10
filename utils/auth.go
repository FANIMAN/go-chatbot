package utils

import (
    "net/http"
)

func AuthMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        cookie, err := r.Cookie("token")
        if err != nil {
            if err == http.ErrNoCookie {
                http.Error(w, "Unauthorized", http.StatusUnauthorized)
                return
            }
            http.Error(w, "Unauthorized", http.StatusUnauthorized)
            return
        }

        tokenString := cookie.Value
        claims, err := ValidateJWT(tokenString)
        if err != nil {
            http.Error(w, "Unauthorized", http.StatusUnauthorized)
            return
        }

        r.Header.Set("username", claims.Username)
        next.ServeHTTP(w, r)
    })
}
