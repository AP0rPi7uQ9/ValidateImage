# ValidateImage
Golang -  Verify that the file is a image type file

Support git,png,jpg

Usage

Install
```
go get github.com/AP0rPi7uQ9/ValidateImage
```

Import
```
import	"github.com/AP0rPi7uQ9/ValidateImage"
```

Validate File
```
//Validate If it is a picture
tp, err := validateImage.ValidateImage(&src)
if err != nil {
fmt.Println(err)
}
```


