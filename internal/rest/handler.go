package rest

import (
	"github.com/RomanVas30/store/internal/rest/handlers"
	"github.com/RomanVas30/store/internal/rest/middlewares"
	"github.com/RomanVas30/store/internal/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", handlers.SignUp(h.services.Authorization))
		auth.POST("/sign-in", handlers.SignIn(h.services.Authorization))
		auth.POST("/change_password", handlers.ChangePassword(h.services.Authorization))
	}

	api := router.Group("/api", middlewares.UserIdentity(h.services.Authorization, true))
	{
		staff := api.Group("/staff")
		{
			staff.POST("/create", handlers.NewStaffer(h.services.Staff))
			staff.GET("/all", handlers.GetStaff(h.services.Staff))
			staff.DELETE("/", handlers.DeleteStaffer(h.services.Staff))
			staff.POST("/search", handlers.SearchStaff(h.services.Staff))
			staff.POST("/update", handlers.UpdateStaffer(h.services.Staff))
		}

		orgUnits := api.Group("/org_units")
		{
			orgUnits.POST("/create", handlers.NewOrgUnit(h.services.OrgUnits))
			orgUnits.GET("/all", handlers.GetOrgUnits(h.services.OrgUnits))
			orgUnits.DELETE("/", handlers.DeleteOrgUnit(h.services.OrgUnits))
			orgUnits.POST("/update", handlers.UpdateOrgUnit(h.services.OrgUnits))
		}

		products := api.Group("/products")
		{
			products.POST("/create", handlers.NewProduct(h.services.Products))
		}
	}

	store := router.Group("/store", middlewares.UserIdentity(h.services.Authorization, false))
	{
		products := store.Group("/products")
		{
			products.GET("/all", handlers.GetProducts(h.services.Products))
			products.GET("/:id", handlers.GetProductById(h.services.Products))
		}

		orders := store.Group("/orders")
		{
			orders.POST("/create", handlers.NewOrder(h.services.Orders))
			orders.GET("/all", handlers.GetOrders(h.services.Orders))
			orders.GET("/:id", handlers.GetOrderById(h.services.Orders))
			orders.GET("/payment/:id", handlers.OrderPayment(h.services.Orders))
			orders.POST("/add_product", handlers.AddProductToOrder(h.services.Orders))
		}
	}

	/*api := router.Group("/api", h.userIdentity)
	{
		lists := api.Group("/lists")
		{
			lists.POST("/", h.createList)
			lists.GET("/", h.getAllLists)
			lists.GET("/:id", h.getListById)
			lists.PUT("/:id", h.updateList)
			lists.DELETE("/:id", h.deleteList)

			items := lists.Group(":id/items")
			{
				items.POST("/", h.createItem)
				items.GET("/", h.getAllItems)
			}
		}

		items := api.Group("items")
		{
			items.GET("/:id", h.getItemById)
			items.PUT("/:id", h.updateItem)
			items.DELETE("/:id", h.deleteItem)
		}
	}*/

	return router
}
