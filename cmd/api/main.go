package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/marifsulaksono/go-echo-boilerplate/internal/api"
	"github.com/marifsulaksono/go-echo-boilerplate/internal/config"
	"github.com/marifsulaksono/go-echo-boilerplate/internal/contract"
	"github.com/marifsulaksono/go-echo-boilerplate/internal/model"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// @title TokoKita API
// @version 1.0
// @description Backend untuk sistem penjualan FIFO

// @host localhost:8080
// @BasePath /

// @securityDefinitions.basic BasicAuth
func main() {
	var (
		ctx       = context.Background()
		isEnvFile = true
	)

	isSeed := flag.Bool("seed", false, "Run user seeder and exit")
	flag.Parse()

	err := config.Load(ctx, isEnvFile)
	if err != nil {
		log.Fatalf("Error load configuration with value isEnvFile = true: %v", err)
	}

	contract, err := contract.NewContract(ctx)
	if err != nil {
		log.Fatalf("Error setup contract / dependecy injection: %v", err)
	}
	contract.Common.AutoMigrate()

	if *isSeed {
		err := seedUserSuperAdmin(contract)
		if err != nil {
			log.Fatalf("Error seed user super admin: %v", err)
			return
		}
		log.Println("Seeder selesai dijalankan.")
		return
	}

	e := api.NewHTTPServer(contract)
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGTERM, syscall.SIGINT)
	go func() {
		e.RunHTTPServer()
	}()

	<-sig

	log.Println("Shutting down....")
	// put all processes to be stopped before successful termination here
	log.Println("Server gracefully terminated.")
}

// this function will create a new role and user data from environment variables. if data already exists, it will skip the seeding process.
func seedUserSuperAdmin(contract *contract.Contract) error {
	db := contract.Common.DB

	roleName := os.Getenv("ROLE_NAME_SEEDER")
	if roleName == "" {
		return fmt.Errorf("ROLE_NAME_SEEDER not set")
	}

	var role model.Role
	err := db.Where("name = ?", roleName).First(&role).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			newRole := model.Role{
				Name: roleName,
			}
			if err := db.Create(&newRole).Error; err != nil {
				return fmt.Errorf("failed to create role: %w", err)
			}
			role = newRole // assign ke variabel role untuk dipakai nanti
			log.Println("Role Super Admin created successfully.")
		} else {
			return fmt.Errorf("failed to check role: %w", err)
		}
	} else {
		log.Println("Role Super Admin already exists, skipping seed role.")
	}

	userName := os.Getenv("USER_NAME_SEEDER")
	if userName == "" {
		return fmt.Errorf("USER_NAME_SEEDER not set")
	}

	userEmail := os.Getenv("USER_EMAIL_SEEDER")
	if userEmail == "" {
		return fmt.Errorf("USER_EMAIL_SEEDER not set")
	}

	userPassword := os.Getenv("USER_PASSWORD_SEEDER")
	if userPassword == "" {
		return fmt.Errorf("USER_PASSWORD_SEEDER not set")
	}

	var count int64
	db.Model(&model.User{}).Where("email = ?", userEmail).Count(&count)
	if count > 0 {
		log.Println("Super admin already exists, skipping user seed.")
		return nil
	}

	hashedPwd, err := bcrypt.GenerateFromPassword([]byte(userPassword), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("failed to hash password: %w", err)
	}

	admin := model.User{
		Name:     userName,
		Email:    userEmail,
		Password: string(hashedPwd),
		RoleID:   role.ID,
	}

	if err := db.Create(&admin).Error; err != nil {
		return fmt.Errorf("failed to create super admin: %w", err)
	}

	log.Println("Super admin seeded successfully.")
	return nil
}
