package handlers

type Coordinate struct {
	Lat float64 `json:"lat"`
	Lon float64 `json:"lon"`
}

type DataRequest struct {
	CisternaId    uint       `json:"cisternaId"`
	GpsCoordinate Coordinate `json:"gpsCoordinate"`
	Volume        float64    `json:"volume"`
	StateMotor    string     `json:"stateMotor"`
	HashDevice    string     `json:"hashDevice"`
	SendDate      string     `json:"sendDate" time_format:"2006-01-02"`
	HashBlck      string     `json:"hashBlck"`
}

type DataResponse struct {
	Id        uint `json:"id"`
	GroupBlck int  `json:"groupBlck"`
}

type Group struct {
	Id        uint   `json:"id"`
	HashGroup string `json:"hashGroup"`
	IdStart   uint   `json:"idStart"`
	IdFinish  uint   `json:"idFinish"`
	HashBlck  string `json:"hashBlck"`
}

type PendingGroupsRequest struct {
	Limit int `json:"limit"`
}

type PendingGroupsResponse struct {
	Groups []Group `json:"groups"`
}

type ProcessPendingGroupsRequest struct {
	Limit int `json:"limit"`
}

type ProcessPendingGroupsResponse struct {
	Groups []Group `json:"groups"`
}

type ExecSetGroupRequest struct {
}

type ExecSetGroupResponse struct {
}

type GroupsRequest struct {
	Page       int `json:"page"`
	TotalPages int `json:"totalPages"`
	Records    int `json:"records"`
}

type GroupResponse struct {
	Groups     []Group `json:"groups"`
	TotalPages int     `json:"totalPages"`
}
