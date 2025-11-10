package exotel

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type ConnectPayload struct {
	From           string `json:"From"`
	To             string `json:"To"`
	CallerID       string `json:"CallerId"`
	StatusCallback string `json:"StatusCallbackUrl"`
	Record         string `json:"Record"`
}

func MakeCall(agent, lead, exophone, webhook string) error {
	apiKey := os.Getenv("EXO_API_KEY")
	apiToken := os.Getenv("EXO_API_TOKEN")
	sid := os.Getenv("EXO_SID")

	payload := ConnectPayload{
		From:           agent,
		To:             lead,
		CallerID:       exophone,
		StatusCallback: webhook,
		Record:         "true",
	}
	data, _ := json.Marshal(payload)

	url := fmt.Sprintf("https://api.exotel.com/v3/accounts/%s/calls/connect", sid)
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(data))
	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth(apiKey, apiToken)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	fmt.Println("📞 Exotel response:", resp.Status)
	return nil
}
