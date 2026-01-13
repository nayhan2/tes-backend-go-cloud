package main

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"tes-database-pq/config"
	"tes-database-pq/models"
	"tes-database-pq/routes"
)

// @title User CRUD API
// @version 1.0
// @description API CRUD sederhana untuk manajemen User dengan Gin dan GORM
// @host localhost:8080
// @BasePath /

func main() {
	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "migrate":
			runMigration()
			return
		case "benchmark":
			iterations := 100
			if len(os.Args) > 2 {
				if n, err := strconv.Atoi(os.Args[2]); err == nil {
					iterations = n
				}
			}
			runBenchmark(iterations)
			return
		case "seed":
			count := 100
			if len(os.Args) > 2 {
				if n, err := strconv.Atoi(os.Args[2]); err == nil {
					count = n
				}
			}
			seedUsers(count)
			return
		}
	}

	// Start server
	config.InitDB()
	r := routes.SetupRoutes()

	fmt.Println("🚀 Server berjalan di http://localhost:8080")
	fmt.Println("📚 Swagger UI: http://localhost:8080/swagger/index.html")
	r.Run(":8080")
}

func runMigration() {
	fmt.Println("🔄 Menjalankan migrasi database...")
	config.InitDB()

	err := config.GetDB().AutoMigrate(&models.User{})
	if err != nil {
		fmt.Println("❌ Gagal migrasi:", err)
		return
	}

	fmt.Println("✅ Migrasi berhasil! Tabel 'users' telah dibuat.")
}

func seedUsers(count int) {
	fmt.Printf("🌱 Membuat %d users untuk testing...\n", count)
	config.InitDB()

	for i := 1; i <= count; i++ {
		user := models.User{
			Name:  fmt.Sprintf("User %d", i),
			Email: fmt.Sprintf("user%d@example.com", i),
		}
		config.GetDB().Create(&user)
	}

	fmt.Printf("✅ Berhasil membuat %d users!\n", count)
}

func runBenchmark(iterations int) {
	fmt.Println("⏱️  BENCHMARK: Testing GET All Users Speed")
	fmt.Println("==========================================")
	config.InitDB()

	// Count users first
	var count int64
	config.GetDB().Model(&models.User{}).Count(&count)
	fmt.Printf("📊 Jumlah users di database: %d\n", count)
	fmt.Printf("🔄 Iterasi: %d\n\n", iterations)

	var users []models.User
	var totalDuration time.Duration
	var minDuration time.Duration = time.Hour
	var maxDuration time.Duration

	fmt.Println("Running benchmark...")

	for i := 0; i < iterations; i++ {
		start := time.Now()
		config.GetDB().Find(&users)
		duration := time.Since(start)

		totalDuration += duration
		if duration < minDuration {
			minDuration = duration
		}
		if duration > maxDuration {
			maxDuration = duration
		}

		// Progress indicator every 10%
		if (i+1)%(iterations/10) == 0 {
			fmt.Printf("  Progress: %d%%\n", (i+1)*100/iterations)
		}
	}

	avgDuration := totalDuration / time.Duration(iterations)

	fmt.Println("\n==========================================")
	fmt.Println("📈 HASIL BENCHMARK:")
	fmt.Println("==========================================")
	fmt.Printf("  Total waktu     : %v\n", totalDuration)
	fmt.Printf("  Rata-rata       : %v\n", avgDuration)
	fmt.Printf("  Tercepat        : %v\n", minDuration)
	fmt.Printf("  Terlambat       : %v\n", maxDuration)
	fmt.Printf("  Queries/detik   : %.2f\n", float64(iterations)/totalDuration.Seconds())
	fmt.Println("==========================================")
}
