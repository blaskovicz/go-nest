package models

type ErrorResponse struct {
	Error       string `json:"error"` // error code
	Description string `json:"error_description"`
	Type        string `json:"type"`
	Instance    string `json:"instance"`
}

type CreateAccessTokenResponse struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int64  `json:"expires_in"` // seconds
}

type Device struct {
	ID              string `json:"device_id"`
	SoftwareVersion string `json:"software_version"`
	StructureID     string `json:"structure_id"`
	WhereID         string `json:"where_id"`
	WhereName       string `json:"where_name"`
	Name            string `json:"name"`
	NameLong        string `json:"name_long"`
}
type Camera struct {
	Device
	IsOnline    bool `json:"is_online"`    // is connected to nest?
	IsStreaming bool `json:"is_streaming"` // actively streaming video?
}
