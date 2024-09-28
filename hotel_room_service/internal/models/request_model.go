package models

type RoomServiceRequest struct {
    ID             string `json:"id"`
    GuestName      string `json:"guestName"`
    RoomNumber     int    `json:"roomNumber"`
    RequestDetails string `json:"requestDetails"`
    Priority       int    `json:"priority"` // Lower numbers indicate higher priority
    Status         string `json:"status"`   // 'received', 'in progress', 'awaiting confirmation', 'completed', 'canceled'
}