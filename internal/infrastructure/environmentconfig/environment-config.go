package environmentconfig

import (
	"ecommerce-go/internal/core/domain"
	"ecommerce-go/internal/core/ports"
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

type environment struct {
	DBDomain              string
	DBUser                string
	DBPass                string
	DBName                string
	DBSsl                 string
	SecretKeyJWT          string
	ScheduleJobs          string
	StockCleanupFrequency string
}

func (env *environment) LoadEnvConfig() error {
	if err := godotenv.Load(".env"); err != nil {
		return err
	}
	env.DBDomain = os.Getenv("DB_DOMAIN")
	env.DBUser = os.Getenv("DB_USER")
	env.DBPass = os.Getenv("DB_PASS")
	env.DBName = os.Getenv("DB_NAME")
	env.DBSsl = os.Getenv("DB_SSL")
	env.SecretKeyJWT = os.Getenv("SECRET_KEY_JWT")
	env.ScheduleJobs = os.Getenv("SCHEDULE_JOBS")
	env.StockCleanupFrequency = os.Getenv("STOCK_CLEANUP_FREQUENCY")

	return nil

}

func NewConfigService() ports.ConfigService {
	envService := environment{}
	if err := envService.LoadEnvConfig(); err != nil {
		fmt.Println("Error loading env config", err)
	}
	return &envService
}

func (env *environment) GetDomainEnv() domain.Environment {
	envDomain := domain.Environment{
		DBDomain:              env.DBDomain,
		DBUser:                env.DBUser,
		DBPass:                env.DBPass,
		DBName:                env.DBName,
		DBSsl:                 env.DBSsl,
		SecretKeyJWT:          env.SecretKeyJWT,
		ScheduleJobs:          env.ScheduleJobs,
		StockCleanupFrequency: env.StockCleanupFrequency}
	return envDomain
}
