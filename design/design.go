package design

import (
  . "github.com/goadesign/goa/design"
  . "github.com/goadesign/goa/design/apidsl"
)

var _ = API("todolist", func() {
  Title("The virtual todolist")
  Description("A simple goa service")
  Scheme("http")
  Host("localhost:8080")
})

var _ = Resource("todolist", func() {
  BasePath("/todolist")
  DefaultMedia(TodoListMedia)

  Action("show", func() {
    Description("Get todo by id")
    Routing(GET("/:todoID"))
    Params(func() {
      param("todoID", Integer, "Todo ID")
    })
    Response(OK)
    Response(NotFound)
  })
})

// TodoListMedia defines the media type used to render todolist.
var TodoListMedia = MediaType("application/vnd.goa.example.todolist+json", func() {
  Description("A todo of todolist")
  Attributes(func() {
    Attributes("id", Integer, "Unique todo ID")
    Attributes("href", String, "Name of todo")
    Required("id", "href", "name")
  })
  View("default", func() {
    Attribute("id")
    Attribute("href")
    Attribute("name")
  })
})