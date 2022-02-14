// // package main

// // import (
// //     "bytes"
// //     "encoding/json"
// //     "fmt"
// //     "log"
// //     "net/http"
// // )

// // func main() {

// //     values := map[string]string{"name": "John Doe", "occupation": "gardener"}
// //     json_data, err := json.Marshal(values)

// //     if err != nil {
// //         log.Fatal(err)
// //     }

// //     resp, err := http.Post("https://httpbin.org/post", "application/json",
// //         bytes.NewBuffer(json_data))

// //     if err != nil {
// //         log.Fatal(err)
// //     }

// //     var res map[string]interface{}

// //     json.NewDecoder(resp.Body).Decode(&res)

// //     fmt.Println(res["json"])
// // }

// // package main

// // import (
// //     "fmt"
// //     "net/http"
// //     "time"
// // )

// // func main() {
// //     call("https://people.zoho.com/people/api/attendance", "POST")
// // }
// // //https://people.zoho.com/people/api/attendance/getUserReport?sdate=<sdate>&edate=<edate>&empId=767099910&emailId=dromero@edrans.com&mapId=<mapId>&dateFormat=<dateFormat>

// // func call(url, method string) error {
// //     client := &http.Client{
// //         Timeout: time.Second * 10,
// //     }
// //     req, err := http.NewRequest(method, url, nil)

// //     if err != nil {
// //         return fmt.Errorf("Got error %s", err.Error())
// //     }

// //     //req.Header.Set("user-agent", "golang application")
// // 	req.Header.Set("Authorization:", "Zoho-oauthtoken 1000.b126f3303e821720f1477f92be351b3c.afdc68b48b67186d79109510bdf6f787")
// // 	req.Header.Add("dateFormat","dd/MM/yyyy HH:mm:ss");
// // 	req.Header.Add("checkIn","11/02/2022 09:30:45");
// // 	//req.Header.set("checkOut","09/09/2013 18:45:13");
// // 	req.Header.Add("empId","2051");

// //     response, err := client.Do(req)

// //     if err != nil {
// //         return fmt.Errorf("Got error %s", err.Error())
// //     }
// //     defer response.Body.Close()
// //     return nil
// // }

// package main

// import (
//     "io/ioutil"
//     "log"
//     "net/http"
// )

// func main() {
//     url := "https://people.zoho.com/people/api/attendance"

//     // Create a Bearer string by appending string access token
//     var bearer = "Authorization: " + "Zoho-oauthtoken Zoho-oauthtoken 1000.8cb99dxxxxxxxx9be93.9b8xxxxxxf"

//     // Create a new request using http
//     req, err := http.NewRequest("GET", url, nil)

//     // add authorization header to the req
//     req.Header.Add("Authorization", bearer)

//     // Send req using http Client
//     client := &http.Client{}
//     resp, err := client.Do(req)

//     if err != nil {
//         log.Println("Error on response.\n[ERROR] -", err)
//     }
//     defer resp.Body.Close()

//     body, err := ioutil.ReadAll(resp.Body)
//     if err != nil {
//         log.Println("Error while reading the response bytes:", err)
//     }
//     log.Println(string([]byte(body)))
// }