package main

import "time"

type News struct {
	Id     int       `json:"id"`
	Title  string    `json:"title"`
	Author string    `json:"author"`
	Date   time.Time `json:"date"`
}

type NewsList []News
