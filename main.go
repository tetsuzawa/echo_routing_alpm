package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"
)

//e.Static("/", "public")
//e.GET("/", func(c echo.Context) error {
//e.GET("/initialize", func(c echo.Context) error {
//e.POST("/api/users", func(c echo.Context) error {
//e.GET("/api/users/:id", func(c echo.Context) error {
//e.POST("/api/actions/login", func(c echo.Context) error {
//e.POST("/api/actions/logout", func(c echo.Context) error {
//e.GET("/api/events", func(c echo.Context) error {
//e.GET("/api/events/:id", func(c echo.Context) error {
//e.POST("/api/events/:id/actions/reserve", func(c echo.Context) error {
//e.DELETE("/api/events/:id/sheets/:rank/:num/reservation", func(c echo.Context) error {
//e.GET("/admin/", func(c echo.Context) error {
//e.POST("/admin/api/actions/login", func(c echo.Context) error {
//e.POST("/admin/api/actions/logout", func(c echo.Context) error {
//e.GET("/admin/api/events", func(c echo.Context) error {
//e.POST("/admin/api/events", func(c echo.Context) error {
//e.GET("/admin/api/events/:id", func(c echo.Context) error {
//e.POST("/admin/api/events/:id/actions/edit", func(c echo.Context) err
//e.GET("/admin/api/reports/events/:id/sales", func(c echo.Context) err
//e.GET("/admin/api/reports/sales", func(c echo.Context) error {

func main() {
	lines := readlines(os.Stdin)
	r, err := regexp.Compile(`"/[^"]*"`)
	if err != nil {
		fmt.Fprint(os.Stderr, "failed to compile regex")
	}
	parameterizedPaths := make([]string, 0, len(lines))
	for _, line := range lines {
		path := r.FindString(line)
		if strings.Contains(path, "/:") {
			parameterizedPaths = append(parameterizedPaths, path)
		}
	}

	builder := &strings.Builder{}
	for i, path := range parameterizedPaths {
		builder.WriteString(strings.Replace(path, `"`, "", -1))
		if i != len(parameterizedPaths)-1 {
			builder.WriteString(",")
		}
	}
	fmt.Println(builder.String())
}

func readlines(reader io.Reader) []string {
	scanner := bufio.NewScanner(reader)
	lines := make([]string, 0)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}
