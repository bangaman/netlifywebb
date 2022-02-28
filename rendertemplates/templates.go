package templates


import (
    "net/http"
    "bytes"
	"html/template"
	"io/ioutil"
	"fmt"
)



func HTML(data interface{}, link string)string{
    url := link
    
    resp, err := http.Get(url)
    // handle the error if there is one
    if err != nil {
        panic(err)
	}
    
    // close 
    defer resp.Body.Close()

	// reads html as a slice of bytes
	html, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
    
    t, _ := template.New("todos").Parse(string(html[:]))
    
    var tpl bytes.Buffer
    if err := t.Execute(&tpl, data); err != nil {
		fmt.Println(err)
	}	

	result := tpl.String()

    return result
}