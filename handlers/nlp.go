package handlers

import (
    "bytes"
    "encoding/json"
    "net/http"
)

type SentimentRequest struct {
    Text string `json:"text"`
}

type SentimentResponse struct {
    Negative float64 `json:"neg"`
    Neutral  float64 `json:"neu"`
    Positive float64 `json:"pos"`
    Compound float64 `json:"compound"`
}

func AnalyzeSentimentHandler(w http.ResponseWriter, r *http.Request) {
    var req SentimentRequest
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    pythonReqBody, err := json.Marshal(req)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    pythonResp, err := http.Post("http://localhost:5000/analyze", "application/json", bytes.NewBuffer(pythonReqBody))
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    defer pythonResp.Body.Close()

    var sentimentResponse SentimentResponse
    if err := json.NewDecoder(pythonResp.Body).Decode(&sentimentResponse); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(sentimentResponse)
}
