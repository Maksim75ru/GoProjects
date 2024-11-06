package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// func calculateCircle(radius float64) float64 {
// 	return radius * radius * math.Pi
// }
// func main() {
// 	http.HandleFunc(
// 		"/",
// 		func(w http.ResponseWriter, r *http.Request) {
// 			fmt.Fprint(w, "Привет Интернет")
// 		},
// 	)
// 	http.HandleFunc(
// 		"/circle_area",
// 		func(w http.ResponseWriter, r *http.Request) {
// 			radiusQuery := r.URL.Query().Get("radius")
// 			number, err := strconv.ParseFloat(radiusQuery, 64)
// 			if err != nil {
// 				fmt.Fprint(w, "Ты кусок гавна, введи число!")
// 			} else {
// 				circleArea := calculateCircle(number)
// 				fmt.Fprint(w, circleArea)
// 			}
// 		},
// 	)
// 	http.ListenAndServe(":8080", nil)
// }

func main() {
	r := gin.Default()
	r.GET(
		"/ping",
		func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"message": "pong"})
		},
	)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
