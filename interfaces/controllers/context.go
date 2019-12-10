package controllers

type Context interface {
	Bind(interface{}) error
	JSON(int, interface{}) error
	FormValue(string) string
	Get(string) interface{}
}
