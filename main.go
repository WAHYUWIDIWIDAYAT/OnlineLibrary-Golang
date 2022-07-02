package main

import (
  	"github.com/gin-gonic/gin"
	"web/models"
	"web/controllers"
	"web/middlewares"
)

func main() {
	//connectDatabase
	models.ConnectDataBase()
	
	r := gin.Default()

	public := r.Group("/api")

	//register for admin
	public.POST("/register", controllers.Register)
	//login for admin
	public.POST("/login",controllers.Login)
	//register for client
	public.POST("/client/register",controllers.RegisterClient)
	//register for client
	public.POST("/client/login",controllers.LoginClient)

	//Group for protected routes in client 
	public = r.Group("/api/client")
	public.Use(middlewares.JwtAuthMiddleware())
	//Show all books in page
	public.GET("/data", controllers.FindBooks)
	//To show information about data user client (only for client)
	public.GET("/user",controllers.CurrentClient)
	//To input information in cart (book) to database (data in this cart only can access by client who have data)
	public.POST("/cart",controllers.AddCart)
	//To show information about data in cart (book) (data in this cart only can access by client who have data)
	public.GET("/cart",controllers.ViewCart)



	//Group for protected routes in admin
	protected := r.Group("/api/admin")
	protected.Use(middlewares.JwtAuthMiddlewareAdmin())
	//To show information about data in database
	protected.GET("/user",controllers.CurrentUser)
	//To show information about data book in database
	protected.GET("/books", controllers.FindBooks)
	//To show specific information about data book in database
	protected.GET("/books/:id", controllers.FindBook)
	//To input information about data book in database
	protected.POST("/books", controllers.CreateBook)
	//To update information about data book in database
	protected.PATCH("/books/:id", controllers.UpdateBook)
	//To delete information about data book in database
	protected.DELETE("/books/:id", controllers.DeleteBook)

	//run server
	r.Run(":8080")

	//localhost:8080/api/register
		// method : POST
		// {username:admin,password:admin}
		// to add user to database
	//localhost:8080/api/login
		// method : POST
		// will return json token to auth
	//localhost:8080/api/admin/user
		// method : GET
		// will return json data user admin with authentication bearer using json token 


	//	client
	//	{   
	//		"username":"april",
	//		"email":"wahyu@gmail.com",
	//		"password":"April598@gmail.com"
	//	}

	//"username":"april123",
	//"email":"wahyu@gmail.com",
	//"password":"wahyuwidi"

}