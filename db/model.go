package db

import "log"

type CallLog struct {
	CallSID      string
	FromNumber   string
	ToNumber     string
	Status       string
	RecordingURL string
}

func InsertCallLog(logEntry CallLog) {
	_, err := DB.Exec(`
		INSERT INTO call_logs (exotel_call_sid, from_number, to_number, status, recording_url)
		VALUES (@p1,@p2,@p3,@p4,@p5)
	`, logEntry.CallSID, logEntry.FromNumber, logEntry.ToNumber, logEntry.Status, logEntry.RecordingURL)
	if err != nil {
		log.Println("❌ DB insert error:", err)
	} else {
		log.Println("✅ Call log saved:", logEntry.CallSID)
	}
}
