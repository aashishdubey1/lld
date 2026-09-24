package srp

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

type Video struct {
	UserId int
	Title  string
	Path   string
	Status string
}

type VideoRepo struct {
	db *sql.DB
}

func (repo *VideoRepo) Create(video Video) error {
	fmt.Println("upllading video", video.Title)
	return nil
}

type Transcoder struct {
	ffmpegpath string
	outputPath string
}

func (t *Transcoder) Transcod(filepath string) error {
	fmt.Println("transcoding file", filepath)
	fmt.Println("saving the output to ", t.outputPath)
	return nil
}

type VidoValidator struct {
	maxSize int
}

func (v *VidoValidator) Validate(path string) error {

	info, err := os.Stat(path)
	if err != nil {
		return err
	}
	if info.Size() > int64(v.maxSize) {
		return errors.New("file is to large")
	}
	ext := filepath.Ext(path)

	if ext != ".mp4" && ext != ".mov" {
		return errors.New("unsupported format")
	}
	return nil
}

type VideoService struct {
	validator  *VidoValidator
	Transcoder *Transcoder
	repo       *VideoRepo
}

// check filepath
// check size limit
// upload meta to db
// give to the transcoding service
// send response
func (s *VideoService) Upload(path string) {
	err := s.validator.Validate(path)
	if err != nil {
		log.Fatalf("invalid video")
	}
	video := Video{
		UserId: 123,
		Title:  "test",
		Path:   path,
		Status: "pending",
	}
	err = s.repo.Create(video)
	if err != nil {
		log.Fatalf("error uploading video: %v", err)
	}

	err = s.Transcoder.Transcod(path)
	if err != nil {
		log.Fatalf("error transcoding video: %v", err)
	}
}
