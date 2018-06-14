package main

import (
	"fmt"
	"image"
	"image/color"
	"io/ioutil"
	"log"
	"time"

	"github.com/golang/freetype"
)

type JSONTime struct {
	time.Time
}

const (
	dx       = 100         // 图片的大小 宽度
	dy       = 40          // 图片的大小 高度
	fontFile = "RAVIE.TTF" // 需要使用的字体文件
	fontSize = 20          // 字体尺寸
	fontDPI  = 72          // 屏幕每英寸的分辨率
)

const CustomTimeFormat = "2006-01-02T15:04:05"

func (ct *JSONTime) UnmarshalJSON(b []byte) (err error) {
	if b[0] == '"' && b[len(b)-1] == '"' {
		b = b[1 : len(b)-1]
	}
	loc := time.Now().Local().Location()
	ct.Time, err = time.ParseInLocation(CustomTimeFormat, string(b), loc)
	return
}

func (ct *JSONTime) MarshalJSON() ([]byte, error) {
	return []byte(ct.Time.Format(CustomTimeFormat)), nil
}

func StringToImage(planText string) (image.Image, error) {
	// 新建一个 指定大小的 RGBA位图
	img := image.NewNRGBA(image.Rect(0, 0, dx, dy))
	// 画背景
	for y := 0; y < dy; y++ {
		for x := 0; x < dx; x++ {
			// 设置某个点的颜色，依次是 RGBA
			img.Set(x, y, color.RGBA{uint8(x), uint8(y), 0, 255})
		}
	}
	// 读字体数据
	fontBytes, err := ioutil.ReadFile(fontFile)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	font, err := freetype.ParseFont(fontBytes)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	c := freetype.NewContext()
	c.SetDPI(fontDPI)
	c.SetFont(font)
	c.SetFontSize(fontSize)
	c.SetClip(img.Bounds())
	c.SetDst(img)
	c.SetSrc(image.White)

	pt := freetype.Pt(10, 0) // 字出现的位置

	_, err = c.DrawString(planText, pt)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return img, nil

}
