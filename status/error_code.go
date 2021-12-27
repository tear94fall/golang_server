package status

const (
	MysqlEnvVarNotSet      = 1000
	MysqlConnFail          = 1001
	MysqlConfFileNotExist  = 1002
	MysqlConfFileIsInvalid = 1003
)

var ErrorMsg = map[int]string{
	MysqlEnvVarNotSet:      "Mysql Env Var Not Set",
	MysqlConnFail:          "Mysql connection fail",
	MysqlConfFileNotExist:  "Mysql conf file is not exist",
	MysqlConfFileIsInvalid: "Mysql conf file is Invalid",
}

func GetErrMsg(errCode int) string {
	return ErrorMsg[errCode]
}
