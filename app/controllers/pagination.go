package controllers

import "tms/app/models"

func MakePageStruct(length int, currentPage int) models.Page {
	listsize := models.ListSize
	var Page models.Page
	for i := 1; i < length/listsize+1; i++ {
		Page.TotalPage = append(Page.TotalPage, i)
	}
	Page.CurrentPage = currentPage
	return Page
}
