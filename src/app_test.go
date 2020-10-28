package main

//	"encoding/json"
//	"net/http"
//	"net/http/httptest"
//	"testing"

//func executeRequest(req *http.Request) *httptest.ResponseRecorder {
//    rr := httptest.NewRecorder()
//	handler := http.HandlerFunc(getBook(repo))
//    handler.ServeHTTP(rr, req)
//
//    return rr
//}
//
//func checkResponseCode(t *testing.T, expected, actual int) {
//    if expected != actual {
//        t.Errorf("Expected response code %d. Got %d\n", expected, actual)
//    }
//}
//
//func TestGetBookThatIsNOTInLibrary(t *testing.T) {
//
//    req, _ := http.NewRequest("GET", "/book/4", nil)
//    response := executeRequest(req)
//
//    checkResponseCode(t, http.StatusNotFound, response.Code)
//
//    var m map[string]string
//    json.Unmarshal(response.Body.Bytes(), &m)
//    if m["error"] != "book not found" {
//        t.Errorf("Expected the 'error' key of the response to be set to 'book not found'. Got '%s'", m["error"])
//    }
//}

//func TestCreateBook(t *testing.T) {
//
//
//	tests := []struct{
//
//		jsonStr []byte
//		jsonStr : []byte(`{"id":4,"name":"a name", "author":"some author","isbn":"1234567890-90897-12912021"}`),
//    }}
//
//    for _, test := range tests {
//	expected := `{"id":4,"name":"a name", "author":"some author","isbn":"1234567890-90897-12912021"}`
//
//    }
//}