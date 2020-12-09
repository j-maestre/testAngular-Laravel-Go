package discotecas

import (

	"errors"
	"github.com/xema/testAngular-Laravel-Go/backend/go2/common"
	"github.com/gin-gonic/gin"
	"net/http"
	
)
// "strconv" para los coments
//fmt para debug


func DiscotecasRegister(router *gin.RouterGroup) {
	router.POST("/", DiscotecaCreate)
	router.PUT("/:id", DiscotecaUpdate)
	router.DELETE("/:id", DiscotecaDelete)
	// router.POST("/:slug/favorite", DiscotecaFavorite)
	// router.DELETE("/:slug/favorite", DiscotecaUnfavorite)
	// router.POST("/:slug/comments", DiscotecaCommentCreate)
	// router.DELETE("/:slug/comments/:id", DiscotecaCommentDelete)
}

func DiscotecasAnonymousRegister(router *gin.RouterGroup) {
	router.GET("/", DiscotecaList)
	router.GET("/:id", DiscotecaById)  
	
}

// router.GET("/:slug/comments", DiscotecaCommentList)

func DiscotecaCreate(c *gin.Context){
	var discoteca Discotecas
	c.BindJSON(&discoteca);

	
	err:=CreateDiscoteca(&discoteca)

	if err !=nil{
		c.AbortWithStatus(http.StatusNotFound)
	}else{
		c.JSON(http.StatusOK, gin.H{"discoteca":discoteca})
		return
	}
}


//List Discotecas
func DiscotecaList(c *gin.Context) {
	var discoteca []Discotecas

	//Busca las discotecas y mete el resultado en la var discoteca
	err := GetAllDiscotecas(&discoteca)
	
	if err != nil {
		c.JSON(http.StatusOK, "Not found")
		c.AbortWithStatus(http.StatusNotFound)
	}else{
		c.JSON(http.StatusOK, discoteca)
	}
}

///////// Find ONE
func DiscotecaById(c *gin.Context) {
	id := c.Params.ByName("id")	

	var discoteca Discotecas
	err := GetDiscotecaById(&discoteca, id)
	
	if err != nil {
		c.JSON(http.StatusOK, "Discoteca Not Found")
		c.AbortWithStatus(http.StatusNotFound)
		return
	}else{
		c.JSON(http.StatusOK, gin.H{"discoteca": discoteca})
		return
	}
	// serializer := DiscotecaSerializer{c, discotecaModel}     //Serializer para enviarlo a angular?
	// c.JSON(http.StatusOK, gin.H{"discoteca": serializer.Response()})
}
////////

//UPDATE discoteca

func DiscotecaUpdate(c *gin.Context){
	var discoteca Discotecas
	var newDiscoteca Discotecas
	c.BindJSON(&newDiscoteca);  //Aqui en teoria está la discoteca que le hemos pasado por postman

	id := c.Params.ByName("id")
	err := GetDiscotecaById(&discoteca, id) //Este es la discoteca que he pillao con ese id, ¿para que? para comprobar que existe ese id

	discoteca.Name = newDiscoteca.Name
	discoteca.Company = newDiscoteca.Company
	discoteca.Events = newDiscoteca.Events


	if err != nil { 
		c.JSON(http.StatusNotFound, "NOT FOUND")
	}else{ 
		c.BindJSON(&discoteca)
		err = UpdateDiscoteca(&discoteca)//&discoteca  Aqui hay que meterle la discoteca nueva, con el c.BingJSON pero no me hace el json de la nueva
		if err != nil {
			c.JSON(http.StatusOK, "Not found")
			c.AbortWithStatus(http.StatusNotFound)
		} else {
			c.JSON(http.StatusOK, gin.H{"discoteca": discoteca})
			return
		}
	}

}


//DELETE discoteca
func DiscotecaDelete(c *gin.Context){
	var discoteca Discotecas
	id := c.Params.ByName("id")

	err := DeleteDiscoteca(&discoteca, id)

	if err != nil{
		c.JSON(http.StatusNotFound, common.NewError("discotecas", errors.New("Invalid id")))
		return
	} 

	c.JSON(http.StatusOK, gin.H{"discoteca": "Delete Discoteca"})

}





