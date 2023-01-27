package app

import (
	"fmt"
	"net/http"

	"forum/internal/model"
	"forum/pkg"
)

func (app *App) FilterHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		pkg.ErrorHandler(w, http.StatusMethodNotAllowed)
		return
	}
	category := (r.URL.Query().Get("category"))
	c, _ := r.Cookie("session_token")
	session, err := app.sessionService.GetSessionByToken(c.Value)
	if err != nil {
		http.Redirect(w, r, "/sign-in-form", http.StatusFound)
		return
	}
	user, err := app.userService.GetUserByEmail(session.Email)
	if err != nil {
		pkg.ErrorHandler(w, http.StatusInternalServerError)
		return
	}
	user.Password = ""
	switch category {
	case "most-like":
		mostLikePost, err := app.postService.GetFilterPosts("most-like")
		if err != nil {
			pkg.ErrorHandler(w, http.StatusInternalServerError)
			return
		}
		data := model.Data{
			Posts: mostLikePost,
			User:  user,
			Genre: "most-like",
		}
		pkg.RenderTemplate(w, "filter.html", data)
		return
	case "most-dislike":
		mostDislikePost, err := app.postService.GetFilterPosts("most-dislike")
		if err != nil {
			pkg.ErrorHandler(w, http.StatusInternalServerError)
			return
		}
		data := model.Data{
			Posts: mostDislikePost,
			User:  user,
			Genre: "most-dislike",
		}
		pkg.RenderTemplate(w, "filter.html", data)
		return
	case "liked-post":
		postId, _ := app.postService.GetLikedPostIdByUser(user)
		var allPost []model.Post
		for _, v := range postId {
			post, err := app.postService.GetPostById(v)
			if err != nil {
				fmt.Println(err)
				pkg.ErrorHandler(w, http.StatusInternalServerError)
				return
			}
			allPost = append(allPost, post)
		}
		data := model.Data{
			Posts: allPost,
			User:  user,
			Genre: "liked-post",
		}
		pkg.RenderTemplate(w, "filter.html", data)
		return
	case "created-post":
		all, _ := app.postService.GetAllPosts()
		var allPost []model.Post
		for _, v := range all {
			if v.Author.Username == user.Username {
				allPost = append(allPost, v)
			}
		}
		data := model.Data{
			Posts: allPost,
			User:  user,
			Genre: "created-post",
		}
		pkg.RenderTemplate(w, "filter.html", data)
		return
	case "fantazy":
		fantazyPost, err := app.postService.GetFilterPosts("fantazy")
		if err != nil {
			pkg.ErrorHandler(w, http.StatusInternalServerError)
			return
		}
		user.Password = ""
		data := model.Data{
			Posts: fantazyPost,
			User:  user,
			Genre: "fantazy",
		}
		pkg.RenderTemplate(w, "filter.html", data)
		return
	case "drama":
		dramaPost, err := app.postService.GetFilterPosts("drama")
		if err != nil {
			pkg.ErrorHandler(w, http.StatusInternalServerError)
			return
		}
		user.Password = ""
		data := model.Data{
			Posts: dramaPost,
			User:  user,
			Genre: "drama",
		}
		pkg.RenderTemplate(w, "filter.html", data)
		return
	case "comedy":
		comedyPost, err := app.postService.GetFilterPosts("comedy")
		if err != nil {
			pkg.ErrorHandler(w, http.StatusInternalServerError)
			return
		}
		user.Password = ""
		data := model.Data{
			Posts: comedyPost,
			User:  user,
			Genre: "comedy",
		}
		pkg.RenderTemplate(w, "filter.html", data)
		return
	case "adventure":
		adventurePost, err := app.postService.GetFilterPosts("adventure")
		if err != nil {
			pkg.ErrorHandler(w, http.StatusInternalServerError)
			return
		}
		data := model.Data{
			Posts: adventurePost,
			User:  user,
			Genre: "adventure",
		}
		pkg.RenderTemplate(w, "filter.html", data)
		return
	case "romance":
		romancePost, err := app.postService.GetFilterPosts("romance")
		if err != nil {
			pkg.ErrorHandler(w, http.StatusInternalServerError)
			return
		}
		data := model.Data{
			Posts: romancePost,
			User:  user,
			Genre: "romance",
		}
		pkg.RenderTemplate(w, "filter.html", data)
		return
	default:
		pkg.ErrorHandler(w, http.StatusBadRequest)
		return
	}
}
