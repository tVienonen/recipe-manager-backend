package auth

import (
	"log"

	"github.com/go-ozzo/ozzo-dbx"

	"github.com/kataras/iris"
	"github.com/tVienonen/recipe-manager-backend/resources/application"
	"github.com/tVienonen/recipe-manager-backend/resources/database"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/go-playground/validator.v9"
)

func RegisterUserAction(ctx iris.Context) {
	validate := ctx.Values().Get("validate").(*validator.Validate)

	user := new(User)
	ctx.ReadJSON(user)

	if err := validate.Struct(user); err != nil {
		ctx.StatusCode(400)
		return
	}
	db := database.GetDB(ctx)
	q := db.NewQuery("SELECT count(email) as userCount FROM users where email = {:email}")
	q.Bind(dbx.Params{"email": user.Email})
	var cnt = struct {
		UserCount int64 `db:"userCount"`
	}{}
	if err := q.One(&cnt); err != nil {
		log.Println(err)
		ctx.StatusCode(500)
		ctx.JSON(application.ApplicationErrorResponse(application.InternalServerError))
		return
	}
	if cnt.UserCount > 0 {
		ctx.StatusCode(400)
		ctx.JSON(application.ApplicationErrorResponse(application.EmailIsInUse))
		return
	}
	// Hash password
	hashed, err := bcrypt.GenerateFromPassword([]byte(*user.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Println(err)
		ctx.StatusCode(500)
		ctx.JSON(application.ApplicationErrorResponse(application.InternalServerError))
		return
	}
	*user.Password = string(hashed)
	q = db.Insert("users", map[string]interface{}{
		"first_name": user.FirstName,
		"last_name":  user.LastName,
		"email":      user.Email,
		"password":   user.Password})
	if res, err := q.Execute(); err != nil {
		log.Println(err)
		ctx.StatusCode(500)
		ctx.JSON(application.ApplicationErrorResponse(application.InternalServerError))
		return
	} else {
		id, err := res.LastInsertId()
		if err != nil {
			log.Println(err)
			ctx.StatusCode(500)
			ctx.JSON(application.ApplicationErrorResponse(application.InternalServerError))
			return
		}
		user.ID = &id
	}
	ctx.StatusCode(201)
	ctx.JSON(PublicUser{
		User: user})
}
