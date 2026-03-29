package store

// ScriptRecord represents a pyRevit script execution record (v2 schema).
type ScriptRecord struct {
	SchemaVersion     string                 `json:"schema"`
	TimeStamp         string                 `json:"timestamp"`
	UserName          string                 `json:"username"`
	HostUserName      string                 `json:"host_user"`
	RevitVersion      string                 `json:"revit"`
	RevitBuild        string                 `json:"revitbuild"`
	SessionId         string                 `json:"sessionid"`
	PyRevitVersion    string                 `json:"pyrevit"`
	Clone             string                 `json:"clone"`
	IsDebugMode       bool                   `json:"debug"`
	IsConfigMode      bool                   `json:"config"`
	IsExecFromGUI     bool                   `json:"from_gui"`
	ExecId            string                 `json:"exec_id"`
	ExecTimeStamp     string                 `json:"exec_timestamp"`
	CommandName       string                 `json:"commandname"`
	CommandUniqueName string                 `json:"commanduniquename"`
	BundleName        string                 `json:"commandbundle"`
	ExtensionName     string                 `json:"commandextension"`
	DocumentName      string                 `json:"docname"`
	DocumentPath      string                 `json:"docpath"`
	ResultCode        int                    `json:"resultcode"`
	CommandResults    map[string]interface{} `json:"commandresults"`
	ScriptPath        string                 `json:"scriptpath"`
	EngineType        string                 `json:"engine_type"`
	EngineVersion     string                 `json:"engine_version"`
}

// EventRecord represents a pyRevit application event record (v2 schema).
type EventRecord struct {
	SchemaVersion    string                 `json:"schema"`
	TimeStamp        string                 `json:"timestamp"`
	HandlerId        string                 `json:"handler_id"`
	EventType        string                 `json:"type"`
	EventArgs        map[string]interface{} `json:"args"`
	UserName         string                 `json:"username"`
	HostUserName     string                 `json:"host_user"`
	RevitVersion     string                 `json:"revit"`
	RevitBuild       string                 `json:"revitbuild"`
	Cancellable      bool                   `json:"cancellable"`
	Cancelled        bool                   `json:"cancelled"`
	DocumentId       int                    `json:"docid"`
	DocumentType     string                 `json:"doctype"`
	DocumentTemplate string                 `json:"doctemplate"`
	DocumentName     string                 `json:"docname"`
	DocumentPath     string                 `json:"docpath"`
	ProjectNumber    string                 `json:"projectnum"`
	ProjectName      string                 `json:"projectname"`
}
