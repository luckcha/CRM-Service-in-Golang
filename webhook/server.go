package webhook

import (
	"fmt"
	"log"
	"net/http"

	"crm-exotel/db"
)

func StartWebhookServer() {
	http.HandleFunc("/webhook", handleWebhook)
	fmt.Println("🌐 Webhook running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleWebhook(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Invalid form", 400)
		return
	}

	entry := db.CallLog{
		CallSID:      r.FormValue("CallSid"),
		FromNumber:   r.FormValue("From"),
		ToNumber:     r.FormValue("To"),
		Status:       r.FormValue("Status"),
		RecordingURL: r.FormValue("RecordingUrl"),
	}

	db.InsertCallLog(entry)
	fmt.Fprint(w, "ok")
}
