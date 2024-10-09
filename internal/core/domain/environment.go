package domain

type Environment struct {
	DBDomain              string
	DBUser                string
	DBPass                string
	DBName                string
	DBSsl                 string
	SecretKeyJWT          string
	ScheduleJobs          string
	StockCleanupFrequency string
}
