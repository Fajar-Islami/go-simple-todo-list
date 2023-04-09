package helper

import (
	"path/filepath"
	"runtime"
)

const (
	FAILEDGETDATA     = "Failed to GET data"
	SUCCEEDGETDATA    = "Succeed to GET data"
	FAILEDPOSTDATA    = "Failed to POST data"
	SUCCEEDPOSTDATA   = "Succeed to POST data"
	FAILEDUPDATEDATA  = "Failed to UPDATE data"
	SUCCEEDUPDATEDATA = "Succeed to UPDATE data"
	FAILEDDELETEDATA  = "Failed to DELETE data"
	SUCCEEDDELETEDATA = "Succeed to DELETE data"
)
const (
	InsertDataFailed  = "Insert Data Failed"
	InsertDataSucceed = "Insert Data Succeed"
	UpdateDataFailed  = "Update Data Failed"
	UpdateDataSucceed = "Update Data Succeed"
	GetDataFailed     = "Get Data Failed"
	GetDataSucceed    = "Get Data Succeed"
	DeleteDataFailed  = "Delete Data Failed"
	DeleteDataSucceed = "Delete Data Succeed"
)

var (
	// Get current file full path from runtime
	_, b, _, _ = runtime.Caller(0)

	// Root folder of this project
	ProjectRootPath = filepath.Join(filepath.Dir(b), "../../")
)
