package handlers

import (
    "encoding/json"
    "net/http"

    "github.com/network-api/fabric"
)

func InvokeChaincode() http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        var request struct {
            ChannelID     string   `json:"channelID"`
            ChaincodeName string   `json:"chaincodeName"`
            FunctionName  string   `json:"functionName"`
            Args          []string `json:"args"`
            Username      string   `json:"username"`
        }
        err := json.NewDecoder(r.Body).Decode(&request)
        if err != nil {
            http.Error(w, err.Error(), http.StatusBadRequest)
            return
        }

        response, err := fabric.InvokeChaincode(request.ChannelID, request.ChaincodeName, request.FunctionName, request.Args, request.Username)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }

        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(map[string]interface{}{
            "response": string(response),
        })
    }
}

func QueryChaincode() http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        var request struct {
            ChannelID     string   `json:"channelID"`
            ChaincodeName string   `json:"chaincodeName"`
            FunctionName  string   `json:"functionName"`
            Args          []string `json:"args"`
            Username      string   `json:"username"`
        }
        err := json.NewDecoder(r.Body).Decode(&request)
        if err != nil {
            http.Error(w, err.Error(), http.StatusBadRequest)
            return
        }

        response, err := fabric.QueryChaincode(request.ChannelID, request.ChaincodeName, request.FunctionName, request.Args, request.Username)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }

        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(map[string]interface{}{
            "response": string(response),
        })
    }
}