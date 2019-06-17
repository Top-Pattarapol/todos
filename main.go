package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

type Todo struct {
	Id     int    `json:"id"`
	Title  string `json:"title"`
	Status string `json:"status"`
}

type Status struct {
	Status string `json:"status"`
}

var todoList = map[int]Todo{}
var idCount = 0

func main() {
	path := "api/todos"
	pathParam := "/:id"
	e := echo.New()
	e.POST(path, postHandler)
	e.GET(path, getHandler)
	e.GET(path+pathParam, getIdHandler)
	e.PUT(path+pathParam, putHandler)
	e.DELETE(path+pathParam, deleteHandler)
	e.Logger.Fatal(e.Start())
}

func postHandler(c echo.Context) error {
	todo := new(Todo)
	if err := c.Bind(todo); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	idCount++
	todoList[idCount] = Todo{Id: idCount, Title: todo.Title, Status: todo.Status}
	return c.JSON(http.StatusCreated, todoList[idCount])
}

func getHandler(c echo.Context) error {
	arr := []Todo{}
	for _, v := range todoList {
		arr = append(arr, v)
	}
	return c.JSON(http.StatusOK, arr)
}

func getIdHandler(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	output := todoList[id]
	return c.JSON(http.StatusOK, output)
}

func putHandler(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	todo := new(Todo)
	if err := c.Bind(todo); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	todoList[id] = Todo{Id: id, Title: todo.Title, Status: todo.Status}
	return c.JSON(http.StatusOK, todoList[id])
}

func deleteHandler(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	delete(todoList, id)
	return c.JSON(http.StatusOK, Status{Status: "success"})
}
