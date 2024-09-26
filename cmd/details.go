package cmd

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/gofrs/uuid/v5"
)

func printDetails(id uuid.UUID) error {
	details := struct {
		UUID      string `json:"uuid"`
		Version   byte   `json:"version"`
	}{
		id.String(),
		id.Version(),
	}

	output, err := json.MarshalIndent(details, "", "    ")
	if err != nil {
		return err
	}

	fmt.Println(string(output))

	return nil
}

func printV7Details(id uuid.UUID) error {
	timestamp, err := uuid.TimestampFromV7(id)
	if err != nil {
		return fmt.Errorf("Analyze Error: %v", err)
	}

	timestampAsTime, err := timestamp.Time()
	if err != nil {
		return fmt.Errorf("Analyze Error: %v", err)
	}

	details := struct {
		UUID      string `json:"uuid"`
		Version   byte   `json:"version"`
		Timestamp string `json:"timestamp"`
	}{
		id.String(),
		id.Version(),
		timestampAsTime.Format(time.RFC3339),
	}

	output, err := json.MarshalIndent(details, "", "    ")
	if err != nil {
		return err
	}

	fmt.Println(string(output))

	return nil
}
