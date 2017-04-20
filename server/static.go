package server

import (
	"bytes"
	"log"
	"net/http"
	"path"
	"path/filepath"

	"go.ua-ecm.com/chaki/static"

	"github.com/labstack/echo"
)

const index = "index.html"

func serveStatic(root string) echo.HandlerFunc {
	return func(c echo.Context) error {
		serveFile := func(name string) error {
			fi, err := static.AssetInfo(name)
			if err != nil {
				return err
			}

			data, err := static.Asset(name)
			if err != nil {
				return err
			}

			r := bytes.NewReader(data)

			http.ServeContent(c.Response(), c.Request(), fi.Name(), fi.ModTime(), r)
			return nil
		}

		_, err := static.AssetInfo(root)
		if err == nil {
			return serveFile(root)
		}

		p := c.Request().URL.Path
		name := filepath.Join(root, path.Clean("/"+p))

		log.Println("serving static file", name, p)

		err = serveFile(name)
		if err != nil {
			ip := filepath.Join(name, index)
			err = serveFile(ip)
		}
		return err
	}
}
