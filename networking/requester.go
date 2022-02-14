// package main

// import (
//     "bytes"
//     "encoding/json"
//     "fmt"
//     "log"
//     "net/http"
// )

// func main() {

//     values := map[string]string{"name": "John Doe", "occupation": "gardener"}
//     json_data, err := json.Marshal(values)

//     if err != nil {
//         log.Fatal(err)
//     }

//     resp, err := http.Post("https://httpbin.org/post", "application/json",
//         bytes.NewBuffer(json_data))

//     if err != nil {
//         log.Fatal(err)
//     }

//     var res map[string]interface{}

//     json.NewDecoder(resp.Body).Decode(&res)

//     fmt.Println(res["json"])
// }

// package main

// import (
//     "fmt"
//     "net/http"
//     "time"
// )

// func main() {
//     request("https://httpbin.org/post", "GET")
// }

// func request(url, method string) error {
//     client := &http.Client{
//         Timeout: time.Second * 10,
//     }
//     req, err := http.NewRequest(method, url, nil)
//     if err != nil {
//         return fmt.Errorf("Got error %s", err.Error())
//     }
//     req.Header.Set("user-agent", "golang application")
//     req.Header.Add("foo", "bar1")
//     req.Header.Add("foo", "bar2")
//     response, err := client.Do(req)
//     if err != nil {
//         return fmt.Errorf("Got error %s", err.Error())
//     }
//     defer response.Body.Close()
//     return nil
// }

// func must(err error){
// 	if err != nil {
// 		panic(err)
// 	}
// }