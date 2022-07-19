package handlers

import (
	"devexcel-excel-api/internal/services"
	"devexcel-excel-api/internal/types"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

func GenerateExcelHandler(ctx *gin.Context) {
	excel := &types.Excel{}

	err := ctx.Bind(&excel)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	excel.Filename = strings.ReplaceAll(excel.Filename, " ", "_")

	target, err := services.BuildExcel(*excel)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	filename := fmt.Sprintf("%s.xlsx", excel.Filename)

	ctx.Header("Content-Description", "File Transfer")
	ctx.Header("Content-Transfer-Encoding", "binary")
	ctx.Header("Content-Disposition", "attachment; filename="+filename)
	ctx.Header("Content-Type", "application/octet-stream")
	ctx.File(target)

	os.Remove(target)
}
