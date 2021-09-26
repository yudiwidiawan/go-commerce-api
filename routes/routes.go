package routes

import (
	"fp-jcc-go-2021-commerce/controllers"
	"fp-jcc-go-2021-commerce/middlewares"

	_ "fp-jcc-go-2021-commerce/docs/commerce"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRoutes(db *gorm.DB) *gin.Engine {
	r := gin.Default()
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
	})
	url := ginSwagger.URL("http://localhost:8080/swagger/doc.json") // The url pointing to API definition
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	// Routes
	public := r.Group("/api")
	public.POST("/register", controllers.CreateUser)
	public.POST("/login", controllers.LoginUser)
	public.GET("/shops/search", controllers.GetShopByKeyword)
	public.GET("/products/search", controllers.GetProductsByKeyword)
	public.GET("/categories", controllers.GetCategories)

	protected := r.Group("/api/admin")
	protected.Use(middlewares.JwtAuthMiddlewareAdmin())
	protected.GET("/users", controllers.FindUsers)
	protected.GET("/users/:user_id", controllers.FindUser)
	protected.PATCH("/users/:user_id", controllers.UpdateUser)
	protected.DELETE("/users/:user_id", controllers.DeleteUser)
	protected.GET("/shops", controllers.GetShops)
	protected.POST("/shops/create", controllers.CreateShop)
	protected.GET("/shops/:shop_id", controllers.FindShop)
	protected.PATCH("/shops/:shop_id", controllers.UpdateShop)
	protected.DELETE("/shops/:shop_id", controllers.DeleteShop)
	protected.GET("/category/:category_id", controllers.GetCategory)
	protected.POST("/category/create", controllers.CreateCategory)
	protected.PATCH("/category/:category_id", controllers.UpdateCategory)
	protected.DELETE("/category/:category_id", controllers.DeleteCategory)

	protectedUser := r.Group("/api/user")
	protectedUser.Use(middlewares.JwtAuthMiddleware())
	// User Profile API
	protectedUser.GET("/details", controllers.SelfDetailUser)
	protectedUser.PATCH("/update", controllers.UpdateSelfUser)
	// User Shops API
	protectedUser.GET("/shops", controllers.GetUserShops)
	protectedUser.POST("/shops/create", controllers.CreateShopUser)
	protectedUser.GET("/shop/:shop_id", controllers.ShopDetailUser)
	protectedUser.PATCH("/shop/:shop_id", controllers.UpdateShopUser)
	protectedUser.DELETE("/shop/:shop_id", controllers.DeleteShopUser)
	// User Shops Etalase API
	protectedUser.GET("/shop/:shop_id/etalase", controllers.GetEtalases)
	protectedUser.POST("/shop/:shop_id/etalase/create", controllers.CreateEtalase)
	protectedUser.GET("/shop/:shop_id/etalase/:etalase_id", controllers.FindEtalase)
	protectedUser.PATCH("/shop/:shop_id/etalase/:etalase_id", controllers.UpdateEtalase)
	protectedUser.DELETE("/shop/:shop_id/etalase/:etalase_id", controllers.DeleteEtalase)
	// User Shops Etalase Product
	protectedUser.GET("/shop/:shop_id/products", controllers.GetShopProducts)
	protectedUser.POST("/etalase/:etalase_id/product/create", controllers.CreateProduct)
	protectedUser.GET("/etalase/:etalase_id/products", controllers.GetProducts)
	protectedUser.POST("/product/create", controllers.CreateProduct)
	protectedUser.PATCH("/product/:product_id", controllers.UpdateProduct)
	protectedUser.DELETE("/product/:product_id", controllers.DeleteProduct)
	// User Address
	protectedUser.GET("/addresses", controllers.GetAddresses)
	protectedUser.POST("/address/create", controllers.CreateAddress)
	protectedUser.PATCH("/address/:address_id", controllers.UpdateAddress)
	protectedUser.DELETE("/address/:address_id", controllers.DeleteAddress)
	// User Product Category
	protectedUser.GET("/product/:product_id/categories", controllers.GetProductCategories)
	protectedUser.POST("/product/category/create", controllers.CreateProductCategory)
	protectedUser.PATCH("/product/category/:product_category_id", controllers.UpdateProductCategory)
	protectedUser.DELETE("/product/category/:product_category_id", controllers.DeleteProductCategory)
	// User Cart Product
	protectedUser.POST("/cart/create", controllers.CreateCart)
	protectedUser.POST("/cart/add", controllers.AddProductCart)
	protectedUser.PATCH("/cart/:cart_id", controllers.UpdateCart)
	protectedUser.PATCH("/cart/product/:cart_product_id", controllers.UpdateProductCart)
	protectedUser.DELETE("/cart/product/:cart_product_id", controllers.DeleteProductCart)
	protectedUser.DELETE("/cart/:cart_id", controllers.DeleteCart)
	return r
}
