package models

type Group struct {
    ID          int    `json:"id"`
    Name        string `json:"name"`
    Description string `json:"description,omitempty"`
    Lists       []List `json:"lists,omitempty"`
}