package app

import (
	"net/http"

	"forum/internal/model"
	"forum/pkg"
)

func (app *App) FilterWelcomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		pkg.ErrorHandler(w, http.StatusMethodNotAllowed)
		return
	}
	category := (r.URL.Query().Get("category"))
	switch category {
	case "most-like":
		mostLikePost, err := app.postService.GetFilterPosts("most-like")
		if err != nil {
			pkg.ErrorHandler(w, http.StatusInternalServerError)
			return
		}
		data := model.Data{
			Posts: mostLikePost,
		}
		pkg.RenderTemplate(w, "unauth.html", data)
	case "most-dislike":
		mostDislikePost, err := app.postService.GetFilterPosts("most-dislike")
		if err != nil {
			pkg.ErrorHandler(w, http.StatusInternalServerError)
			return
		}
		data := model.Data{
			Posts: mostDislikePost,
		}
		pkg.RenderTemplate(w, "unauth.html", data)
	case "fantazy":
		fantazyPost, err := app.postService.GetFilterPosts("fantazy")
		if err != nil {
			pkg.ErrorHandler(w, http.StatusInternalServerError)
			return
		}
		data := model.Data{
			Posts: fantazyPost,
		}
		pkg.RenderTemplate(w, "unauth.html", data)
	case "drama":
		dramaPost, err := app.postService.GetFilterPosts("drama")
		if err != nil {
			pkg.ErrorHandler(w, http.StatusInternalServerError)
			return
		}
		data := model.Data{
			Posts: dramaPost,
		}
		pkg.RenderTemplate(w, "unauth.html", data)
	case "comedy":
		comedyPost, err := app.postService.GetFilterPosts("comedy")
		if err != nil {
			pkg.ErrorHandler(w, http.StatusInternalServerError)
			return
		}
		data := model.Data{
			Posts: comedyPost,
		}
		pkg.RenderTemplate(w, "unauth.html", data)
	case "adventure":
		adventurePost, err := app.postService.GetFilterPosts("adventure")
		if err != nil {
			pkg.ErrorHandler(w, http.StatusInternalServerError)
			return
		}
		data := model.Data{
			Posts: adventurePost,
		}
		pkg.RenderTemplate(w, "unauth.html", data)
	case "romance":
		romancePost, err := app.postService.GetFilterPosts("romance")
		if err != nil {
			pkg.ErrorHandler(w, http.StatusInternalServerError)
			return
		}
		data := model.Data{
			Posts: romancePost,
		}
		pkg.RenderTemplate(w, "unauth.html", data)
	default:
		pkg.ErrorHandler(w, http.StatusBadRequest)
		return
	}
}
