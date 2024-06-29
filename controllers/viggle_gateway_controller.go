package controllers

import (
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
)

var urlforgateway = "http://localhost:10000"

func Viggle(c *gin.Context) {
	// Create a new request based on the original request
	req, err := http.NewRequest(c.Request.Method, urlforgateway+c.Request.URL.Path, c.Request.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Copy headers from the original request
	for key, values := range c.Request.Header {
		for _, value := range values {
			req.Header.Add(key, value)
		}
	}

	// Perform the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer resp.Body.Close()

	// Copy headers from the response
	for key, values := range resp.Header {
		for _, value := range values {
			c.Writer.Header().Add(key, value)
		}
	}

	// Copy the status code
	c.Writer.WriteHeader(resp.StatusCode)

	// Copy the response body
	_, err = io.Copy(c.Writer, resp.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
}

// UploadVideo godoc
// @Summary Upload a video
// @Description Upload a video to the server
// @Tags viggle
// @Produce  json
// @Param video formData file true "Video file"
// @Success 200 {string} string "Video uploaded successfully"
// @Failure 400 {string} string "Bad request"
// @Failure 500 {string} string "Internal server error"
// @Router /upload_video [post]
func UploadVideo(c *gin.Context) {
	Viggle(c)
}

// UploadImage godoc
// @Summary Upload an image
// @Description Upload an image to the server
// @Tags viggle
// @Produce  json
// @Param image formData file true "Image file"
// @Success 200 {string} string "Image uploaded successfully"
// @Failure 400 {string} string "Bad request"
// @Failure 500 {string} string "Internal server error"
// @Router /upload_image [post]
func UploadImage(c *gin.Context) {
	Viggle(c)
}

// SetVideoTask godoc
// @Summary Set a video processing task
// @Description Set a video processing task for the server
// @Tags viggle
// @Accept  json
// @Produce  json
// @Param data body models.SetVideoTask true "Task data"
// @Success 200 {string} string "Task set successfully"
// @Failure 400 {string} string "Bad request"
// @Failure 500 {string} string "Internal server error"
// @Router /set_video_task [post]
func SetVideoTask(c *gin.Context) {
	Viggle(c)
}

// GetTaskResult godoc
// @Summary Get the result of a task
// @Description Get the result of a task from the server
// @Tags viggle
// @Accept  json
// @Produce  json
// @Param data body models.TaskResult true "Task data"
// @Success 200 {string} string "Task result"
// @Failure 400 {string} string "Bad request"
// @Failure 500 {string} string "Internal server error"
// @Router /get_task_result [post]
func GetTaskResult(c *gin.Context) {
	Viggle(c)
}
