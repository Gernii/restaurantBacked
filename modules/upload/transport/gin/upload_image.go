package ginupload

import (
	"net/http"
	"restaurantBacked/common"
	"restaurantBacked/component/appctx"
	bizupload "restaurantBacked/modules/upload/biz"

	"github.com/gin-gonic/gin"
)

func UploadImage(appCtx appctx.AppContext) func(*gin.Context) {
	return func(c *gin.Context) {
		fileHeader, err := c.FormFile("file")
		if err != nil {
			panic(common.ErrInvalidRequest(err))

		}
		folder := c.DefaultPostForm("folder", "img")
		file, err := fileHeader.Open()
		if err != nil {
			panic(common.ErrInvalidRequest(err))

		}
		defer file.Close()

		dataByte := make([]byte, fileHeader.Size)
		if _, err := file.Read(dataByte); err != nil {
			panic(common.ErrInvalidRequest(err))

		}

		biz := bizupload.NewUploadBiz(appCtx.UploadProvider())
		img, err := biz.Upload(c.Request.Context(), dataByte, folder, fileHeader.Filename)
		if err != nil {
			panic(err)
		}
		img.Fulfill(appCtx.UploadProvider().GetDomain())
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(img))
	}
}
