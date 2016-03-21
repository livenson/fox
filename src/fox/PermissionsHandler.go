package fox

import(
	"strings"
	"authn"
	"authz"
	"net/http"

)


// PermissionHandler validates the permissions of a user before further handling
func PermissionHandler(inner http.Handler) http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		var user string
		 
		t := r.Header.Get("Authorization")
		if strings.HasPrefix(t, "Bearer "){
			user, _ = authn.Validate(strings.SplitAfter(t, "Bearer ")[1])
			log.Debugf("Getting user %s from %s", user, t)	
		}else{
			user = ""
		}
		
		if authz.GetProvider().IsAuthorized(user, r.Method, r.URL.RequestURI()){
			sw := makeLogger(w)
			inner.ServeHTTP(sw, r)		
		} else {
				w.WriteHeader(http.StatusUnauthorized)
		}
						
	})
}
