package mock

import (
	_ "embed"
	"encoding/json"
	"sort"
	"sync"
)

type File struct {
	//预览真正关键的其实只有
	ID   string `json:"id"`
	Name string `json:"name"`
	Size int64  `json:"size"`
	URL  string `json:"url"`
	// 而这些字段，对于一次性只读预览通常没有实际业务价值，但 WPS 为了统一“预览、编辑、协作、版本管理”整套协议，没有把只读预览单独设计成更轻的契约，因此仍要求提供
	Version    int64  `json:"version"`
	CreateTime int64  `json:"create_time"`
	ModifyTime int64  `json:"modify_time"`
	CreatorID  string `json:"creator_id"`
	ModifierID string `json:"modifier_id"`
}

//go:embed file.json
var filesJSON []byte

var (
	files     map[string]File
	filesErr  error
	filesOnce sync.Once
)

func loadFiles() error {
	filesOnce.Do(func() {
		filesErr = json.Unmarshal(filesJSON, &files)
	})
	return filesErr
}

func GetFiles() ([]File, error) {
	if err := loadFiles(); err != nil {
		return nil, err
	}

	result := make([]File, 0, len(files))
	for _, file := range files {
		result = append(result, file)
	}
	sort.Slice(result, func(i, j int) bool {
		return result[i].ID < result[j].ID
	})
	return result, nil
}

func GetFile(fileID string) (File, bool, error) {
	if err := loadFiles(); err != nil {
		return File{}, false, err
	}

	file, exists := files[fileID]
	return file, exists, nil
}
