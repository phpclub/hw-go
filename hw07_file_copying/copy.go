package main

import (
	"errors"
	"fmt"
	"io"
	"os"

	"github.com/cheggaaa/pb/v3"
)

var (
	ErrUnsupportedFile       = errors.New("unsupported file")
	ErrOffsetExceedsFileSize = errors.New("offset exceeds file size")
	ErrMissingParamFrom      = errors.New("missing input -from")
	ErrMissingParamTo        = errors.New("missing output -to")
)

const tmpl = `{{ red "Process:" }} {{ bar . "[" "*" (cycle . "↖" "↗" "↘" "↙" ) "." "]" | green }} {{speed . | green }} {{percent . | red }}`

func Copy(fromPath string, toPath string, offset, limit int64) error {
	if fromPath == "" {
		return ErrMissingParamFrom
	}
	if toPath == "" {
		return ErrMissingParamTo
	}
	// Обработаем fileFrom
	fileFrom, err := os.OpenFile(fromPath, os.O_RDONLY, 0)
	if err != nil {
		return err
	}
	defer fileFrom.Close()
	fileFromStat, err := fileFrom.Stat()
	if err != nil {
		return ErrUnsupportedFile
	}
	if fileFromStat.Size() == 0 {
		return ErrUnsupportedFile
	}
	if offset > fileFromStat.Size() {
		return ErrOffsetExceedsFileSize
	}
	if limit == 0 {
		limit = fileFromStat.Size()
	} else if limit > fileFromStat.Size() {
		limit = fileFromStat.Size()
	}
	// Обработаем fileTo
	fileTo, err := os.Create(toPath)
	if err != nil {
		return err
	}
	defer fileTo.Close()

	if offset > 0 {
		//Перемотаем позицию файла
		_, err := fileFrom.Seek(offset, 0)
		if err != nil {
			return ErrOffsetExceedsFileSize
		}
	}

	reader := io.LimitReader(fileFrom, limit)
	// start new bar
	bar := pb.ProgressBarTemplate(tmpl).Start64(limit)
	// create proxy reader
	barReader := bar.NewProxyReader(reader)
	//if offset == 0 {
	written, err := io.CopyN(fileTo, barReader, limit)
	if err != nil {
		return err
	}
	fmt.Printf("Copyng: %d bytes\n", written)
	// finish bar
	bar.Finish()
	return nil
}
