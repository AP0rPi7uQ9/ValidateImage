package validateImage

import (
	"errors"
	"image"
	"mime/multipart"
)

func ValidateImage(f multipart.File) (string, error) {
	// 识别图片类型
	_, image_type, _ := image.Decode(f)

	// 获取图片的类型
	switch image_type {
	case `jpeg`:
		return "jpeg", nil
	case `png`:
		return "png", nil
	case `gif`:
		return "git", nil
	case `bmp`:
		return "bmp", nil
	default:
		err := errors.New("This is not an image file, or the image file format is not supported!")
		return "", err
	}

}