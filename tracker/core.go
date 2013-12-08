
package tracker

import(
	_ "github.com/lib/pq"
	"database/sql"
	"encoding/json"
	"net/http"
)

//= Global DB Pointer initialised in main.go
var Db *sql.DB

//-------------------------------------------------------------------------
// == Ajax Helpers ==
// 
// Returns a payload with  {error: "My Error", success: true}
func CreateAjaxErrorPayload(description string, err error) string{

	err_str := description
	if err != nil {
		err_str = err_str + " " + err.Error()
	}
	payload := map[string]interface{}{
        "error": err_str,
        "success":  true, //This is a quirk for extjs
    }
	s, _ := json.MarshalIndent(payload, "" , "  ")
	return string(s)
}

func SetAjaxHeaders(w http.ResponseWriter){
	
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	//w.Header().Set("Access-Control-Allow-Origin", "*") // TODO cache control ?
}