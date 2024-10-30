package reader

import (
	"encoding/json"
	"github.com/ciscomsk/otus_golang_basic_2023_07_hw/hw02_fix_app/types"
)
import "io"
import "os"

func ReadJSON(filePath string) ([]types.Employee, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}

	bytes, err := io.ReadAll(f)
	if err != nil {
		return nil, err
	}

	var data []types.Employee

	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}

	return data, nil
}
