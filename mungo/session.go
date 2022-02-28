
package mungo

import (
	"github.com/gorilla/sessions"
)


var (
	
	Store = sessions.NewCookieStore([]byte("top-name"))
)