package mediaUpload

import (
	"bytes"
	"io"
	"net/http"
	"os"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"github.com/vishalrana9915/demo_app/pkg/constant"
)

// Api handler to upload media to cloud
func UploadMedia(c *gin.Context) {

	// get file from
	file, _ := c.FormFile("file")

	src, _ := file.Open()
	// close file
	defer src.Close()

	// creating buffer for teh file
	fileBuffer := bytes.NewBuffer(nil)
	io.Copy(fileBuffer, src)

	cloud_name := os.Getenv("CLOUDINARY_CLOUD_NAME")
	api_key := os.Getenv("CLOUDINARY_API_KEY")
	api_secret := os.Getenv("CLOUDINARY_API_SECRET")

	// Initialize the Cloudinary client
	cld, _ := cloudinary.NewFromParams(cloud_name, api_key, api_secret)

	// generate a unique file name
	fileName := uuid.NewV4().String() + ".jpg"

	// Upload the file to Cloudinary
	resp, err := cld.Upload.Upload(c, fileBuffer, uploader.UploadParams{PublicID: fileName})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": constant.UPLOAD_FAILED,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"message": constant.SUCCESS,
		"url":     resp.SecureURL,
	})

}
