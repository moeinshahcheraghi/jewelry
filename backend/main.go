package main

import (
    "github.com/gin-gonic/gin"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
    "log"
    "os"
)

type Post struct {
    ID      uint   `json:"id" gorm:"primaryKey"`
    Title   string `json:"title"`
    Content string `json:"content"`
}

var db *gorm.DB

// بارگذاری تنظیمات و اتصال به پایگاه داده
func connectDatabase() {
    var err error

    // بارگذاری اطلاعات از متغیرهای محیطی
    host := os.Getenv("DB_HOST")
    user := os.Getenv("DB_USER")
    password := os.Getenv("DB_PASSWORD")
    dbname := os.Getenv("DB_NAME")
    port := os.Getenv("DB_PORT")

    // ساخت رشته اتصال به پایگاه داده
    dsn := "host=" + host + " user=" + user + " password=" + password + " dbname=" + dbname + " port=" + port + " sslmode=disable"
    
    db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatalf("Could not connect to the database: %v", err)
    } else {
        log.Println("Connected to the database successfully!")
    }
}

// تنظیمات روت‌ها
func setupRouter() *gin.Engine {
    r := gin.Default()

    // Endpoint برای دریافت همه پست‌ها
    r.GET("/api/posts", func(c *gin.Context) {
        var posts []Post
        if err := db.Find(&posts).Error; err != nil {
            log.Printf("Error retrieving posts: %v", err)
            c.JSON(500, gin.H{"error": "Could not retrieve posts"})
            return
        }
        c.JSON(200, posts)
    })

    // Endpoint برای دریافت یک پست خاص
    r.GET("/api/posts/:id", func(c *gin.Context) {
        var post Post
        if err := db.First(&post, c.Param("id")).Error; err != nil {
            log.Printf("Post not found: %v", err)
            c.JSON(404, gin.H{"error": "Post not found"})
            return
        }
        c.JSON(200, post)
    })

    // Endpoint برای ایجاد پست جدید
    r.POST("/api/posts", func(c *gin.Context) {
        var newPost Post
        if err := c.ShouldBindJSON(&newPost); err != nil {
            log.Printf("Error binding JSON: %v", err)
            c.JSON(400, gin.H{"error": "Invalid request"})
            return
        }

        if err := db.Create(&newPost).Error; err != nil {
            log.Printf("Error creating post: %v", err)
            c.JSON(500, gin.H{"error": "Could not create post"})
            return
        }

        c.JSON(201, newPost)
    })

    // Endpoint برای به‌روزرسانی پست
    r.PUT("/api/posts/:id", func(c *gin.Context) {
        var post Post
        if err := db.First(&post, c.Param("id")).Error; err != nil {
            log.Printf("Post not found: %v", err)
            c.JSON(404, gin.H{"error": "Post not found"})
            return
        }

        var updatedPost Post
        if err := c.ShouldBindJSON(&updatedPost); err != nil {
            log.Printf("Error binding JSON: %v", err)
            c.JSON(400, gin.H{"error": "Invalid request"})
            return
        }

        post.Title = updatedPost.Title
        post.Content = updatedPost.Content

        if err := db.Save(&post).Error; err != nil {
            log.Printf("Error updating post: %v", err)
            c.JSON(500, gin.H{"error": "Could not update post"})
            return
        }

        c.JSON(200, post)
    })

    // Endpoint برای حذف پست
    r.DELETE("/api/posts/:id", func(c *gin.Context) {
        if err := db.Delete(&Post{}, c.Param("id")).Error; err != nil {
            log.Printf("Error deleting post: %v", err)
            c.JSON(500, gin.H{"error": "Could not delete post"})
            return
        }
        c.Status(204)
    })

    return r
}

// تابع اصلی
func main() {
    connectDatabase()
    r := setupRouter()
    if err := r.Run(":8080"); err != nil {
        log.Fatalf("Failed to run server: %v", err)
    }
}

