package main

import "fmt"

var currentId int

var newsList NewsList
var news News

// Give us some seed data
func init() {
	RepoCreateNews(News{Title: "Write presentation", Author: "Raka Westu"})
	RepoCreateNews(News{Title: "Host meetup", Author: "Kuncara Adi"})
}

func RepoFindNews(id int) News {
	for _, t := range newsList {
		if t.Id == id {
			news = t
			return t
		}
	}
	// return empty Todo if not found
	news = News{}
	return News{}
}

func RepoCreateNews(t News) News {
	currentId += 1
	t.Id = currentId
	newsList = append(newsList, t)
	return t
}

func RepoDestroyNews(id int) error {
	for i, t := range newsList {
		if t.Id == id {
			newsList = append(newsList[:i], newsList[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("Could not find Todo with id of %d to delete", id)
}
