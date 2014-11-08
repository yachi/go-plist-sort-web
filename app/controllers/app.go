package controllers

import "github.com/revel/revel"
import "howett.net/plist"
import "bytes"

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	return c.Render()
}

func (c App) Plist(plist_input string) revel.Result {
	// string to bytes and unmarshal
	plist_bytes := []byte(plist_input)
	var plist_dict map[string]map[string]interface{}
	plist.Unmarshal(plist_bytes, &plist_dict)

	// encode to plist xml string
	var plist_buf bytes.Buffer
	encoder := plist.NewEncoder(&plist_buf)
	encoder.Indent("  ")
	encoder.Encode(plist_dict)
	plist_result := plist_buf.String()

	return c.Render(plist_result)
}
