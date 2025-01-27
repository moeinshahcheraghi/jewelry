package controllers

import (
    "log"
    "net/http"
    "os"
    "time"

    "github.com/gin-gonic/gin"
    "github.com/golang-jwt/jwt/v4"
    "golang.org/x/crypto/bcrypt"
    "jewelry/backend/config"
    "jewelry/backend/models"
)

func generateToken(userID uint, role string) (string, error) {
    secretKey := os.Getenv("JWT_SECRET")
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "user_id": userID,
        "role":    role,
        "exp":     time.Now().Add(24 * time.Hour).Unix(),
    })

    return token.SignedString([]byte(secretKey))
}

func Register(c *gin.Context) {
    var input struct {
        Name     string `json:"name" binding:"required"`
        Email    string `json:"email" binding:"required,email"`
        Password string `json:"password" binding:"required,min=6"`
        Role     string `json:"role"` 
    }

    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    var existingUser models.User
    if err := config.DB.Where("email = ?", input.Email).First(&existingUser).Error; err == nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Email is already registered"})
        return
    }

    if input.Role == "" {
        input.Role = "user" 
    }

    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
    if err != nil {
        log.Println("Error hashing password:", err)
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not hash password"})
        return
    }

    user := models.User{
        Name:     input.Name,
        Email:    input.Email,
        Password: string(hashedPassword),
        Role:     input.Role,
    }

    if err := config.DB.Create(&user).Error; err != nil {
        log.Println("Error creating user:", err)
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create user"})
        return
    }

    c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})
}
func Login(c *gin.Context) {
    var input struct {
        Email    string `json:"email" binding:"required,email"`
        Password string `json:"password" binding:"required"`
    }

    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    var user models.User
    if err := config.DB.Where("email = ?", input.Email).First(&user).Error; err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
        return
    }

    if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
        return
    }

    token, err := generateToken(user.ID, user.Role)  // نقش کاربر را در توکن قرار می‌دهیم
    if err != nil {
        log.Println("Error generating token:", err)
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate token"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"token": token})
}

