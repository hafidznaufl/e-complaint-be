package web

type CommentResponse struct {
	ID           string `json:"id"`
	Complaint_ID string `json:"complaintId"`
	Fullname     string `json:"fullname"`
	Role         string `json:"role"`
	Message      string `json:"message"`
}
