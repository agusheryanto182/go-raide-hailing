package main

import (
	"context"
	"log"

	"github.com/gofiber/fiber/v2/middleware/cors"
	fiberLogger "github.com/gofiber/fiber/v2/middleware/logger"

	"github.com/agusheryanto182/go-raide-hailing/config"
	imageController "github.com/agusheryanto182/go-raide-hailing/module/feature/image/controller"
	imageService "github.com/agusheryanto182/go-raide-hailing/module/feature/image/service"
	"github.com/go-playground/validator/v10"

	purchaseController "github.com/agusheryanto182/go-raide-hailing/module/feature/purchase/controller"
	purchaseRepo "github.com/agusheryanto182/go-raide-hailing/module/feature/purchase/repository"
	purchaseService "github.com/agusheryanto182/go-raide-hailing/module/feature/purchase/service"

	merchantController "github.com/agusheryanto182/go-raide-hailing/module/feature/merchant/controller"
	merchantRepo "github.com/agusheryanto182/go-raide-hailing/module/feature/merchant/repository"
	merchantService "github.com/agusheryanto182/go-raide-hailing/module/feature/merchant/service"

	userController "github.com/agusheryanto182/go-raide-hailing/module/feature/user/controller"
	userRepo "github.com/agusheryanto182/go-raide-hailing/module/feature/user/repository"
	userService "github.com/agusheryanto182/go-raide-hailing/module/feature/user/service"
	"github.com/agusheryanto182/go-raide-hailing/module/middleware"
	"github.com/agusheryanto182/go-raide-hailing/routes"
	"github.com/agusheryanto182/go-raide-hailing/utils/database"
	"github.com/agusheryanto182/go-raide-hailing/utils/hash"
	"github.com/agusheryanto182/go-raide-hailing/utils/jwt"
	"github.com/agusheryanto182/go-raide-hailing/utils/logging"
	statementutil "github.com/agusheryanto182/go-raide-hailing/utils/statementUtils"
	"github.com/agusheryanto182/go-raide-hailing/utils/validation"
	awsCfg "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	app := fiber.New(fiber.Config{
		ErrorHandler: middleware.ErrorHandler,
		AppName:      "Project Sprint Week 4 - Raide Hailing",
	})
	cfg := config.NewConfig()

	hash := hash.NewHash(cfg)
	jwt := jwt.NewJWTService(cfg)
	valid := validator.New()

	// register validation
	if err := valid.RegisterValidation("imageUrl", validation.ValidateImageURL); err != nil {
		logging.GetLogger("validator").Error(err.Error())
	}

	// AWS Config
	awsConfig, err := awsCfg.LoadDefaultConfig(
		context.Background(),
		awsCfg.WithRegion("ap-southeast-1"),
		awsCfg.WithCredentialsProvider(
			credentials.NewStaticCredentialsProvider(
				cfg.AWS.ID,
				cfg.AWS.SecretKey,
				"",
			),
		),
	)

	if err != nil {
		logging.GetLogger().Error(err.Error())
	}

	db, err := database.InitDatabase(cfg)
	if err != nil {
		logging.GetLogger("database").Error(err.Error())
	}

	statementutil.SetUp(db)
	defer statementutil.CleanUp()

	defer db.Close()

	app.Use(recover.New())
	app.Use(fiberLogger.New())
	app.Use(cors.New())

	// repo
	userRepo := userRepo.NewUserRepository()
	merchantRepo := merchantRepo.NewMerchantRepository(db)
	purchaseRepo := purchaseRepo.NewPurchaseRepository(db)

	// service
	userService := userService.NewUserService(userRepo, jwt, hash)
	imageService := imageService.NewImageService(awsConfig, cfg.AWS.BucketName)
	merchantService := merchantService.NewMerchantService(merchantRepo)
	purchaseService := purchaseService.NewPurchaseService(purchaseRepo)

	// controller
	userController := userController.NewUserController(userService, valid)
	imageController := imageController.NewImageController(imageService)
	merchantController := merchantController.NewMerchantController(merchantService, valid)
	purchaseController := purchaseController.NewPurchaseController(purchaseService, valid)

	// route
	routes.UserRoute(app, userController, jwt, userService)
	routes.ImageRoute(app, imageController, jwt, userService)
	routes.MerchantRoute(app, merchantController, jwt, userService)
	routes.PurchaseRoute(app, purchaseController, jwt, userService)

	log.Fatal(app.Listen(":8080"))
}
