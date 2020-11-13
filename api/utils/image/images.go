package image

import (
	"strconv"
	"strings"
)

func ImageSize(imageURL string, width, height int) string {
	if strings.Contains(imageURL, "imageView") || !strings.Contains(imageURL, "img.hacpai.com") {
		return imageURL
	}

	return imageURL + "?imageView2/1/w/" + strconv.Itoa(width) + "/h/" + strconv.Itoa(height) + "/interlace/1/q/100"
}
