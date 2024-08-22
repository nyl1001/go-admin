package apis

import (
	"encoding/base64"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/xuri/excelize/v2"
)

type Vehicle struct {
	Name   string   `json:"name"`
	Model  string   `json:"model"`
	Year   int      `json:"year"`
	Images []string `json:"images"` // 图片的 Base64 编码
}

// chatgpt 提示词： 请使用gin post接口 及github.com/xuri/excelize/v2包实现 excel表格数据上传功能，将表格数据解析到一个对象列表中，需要注意表格中车辆图片栏中可能包含多张车辆图片，这多张图片在同一个单元格中
func upload(c *gin.Context) {
	// 从请求中获取文件
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to upload file"})
		return
	}

	// 打开上传的文件
	src, err := file.Open()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to open file"})
		return
	}
	defer src.Close()

	// 使用 excelize 打开 Excel 文件
	f, err := excelize.OpenReader(src)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read Excel file"})
		return
	}

	// 假设数据在第一个工作表中，从第二行开始（跳过标题行）
	rows, err := f.GetRows(f.GetSheetName(1))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get rows from Excel file"})
		return
	}

	var vehicles []Vehicle

	for i, row := range rows {
		if i == 0 {
			// 跳过标题行
			continue
		}

		if len(row) < 4 {
			continue // 如果某行数据不足，跳过
		}

		// 读取车辆信息
		year, err := strconv.Atoi(row[2])
		if err != nil {
			continue // 如果转换失败，跳过该行
		}

		vehicle := Vehicle{
			Name:  row[0],
			Model: row[1],
			Year:  year,
		}

		// 处理车辆图片
		imagesCell := row[3] // 假设第4列为车辆图片列，可能包含多个图片路径
		if imagesCell != "" {
			imagePaths := strings.Split(imagesCell, ",") // 使用逗号作为分隔符
			for _, imagePath := range imagePaths {
				imagePath = strings.TrimSpace(imagePath)
				if imagePath != "" {
					// 读取图片文件
					imageData, err := os.ReadFile(imagePath)
					if err != nil {
						log.Printf("Failed to read image: %v", err)
						continue
					}

					// 将图片转换为 Base64 编码
					encodedImage := base64.StdEncoding.EncodeToString(imageData)
					vehicle.Images = append(vehicle.Images, encodedImage)
				}
			}
		}

		vehicles = append(vehicles, vehicle)
	}

	// 返回成功响应，并包含解析后的车辆数据
	c.JSON(http.StatusOK, gin.H{"message": "File uploaded successfully", "data": vehicles})
}
