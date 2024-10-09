package jobs

import (
	"ecommerce-go/internal/core/domain"
	"ecommerce-go/internal/core/ports"
	"fmt"
	"strconv"
	"time"
)

func ScheduleAllJobs(ecommerceService ports.EcommerceService, envVariables domain.Environment) {
	runScheduleJobs, _ := strconv.ParseBool(envVariables.ScheduleJobs)
	if runScheduleJobs {
		scheduleJob(stockCleanUp, ecommerceService, envVariables, envVariables.StockCleanupFrequency)
	}
}

func stockCleanUp(service ports.EcommerceService) {
	rowsAffected := service.DeleteProductWithoutStock()

	fmt.Println(rowsAffected)
}

func scheduleJob(job func(ports.EcommerceService),
	ecommerceService ports.EcommerceService, envVariables domain.Environment,
	frequency string) {
	// Calculate the duration until the next execution after the first execution
	durationUntilNextExecution, err := time.ParseDuration(frequency)
	if err != nil {
		errorMessage := fmt.Errorf("invalid job frequency: %w", err)
		panic(errorMessage)
	}

	// Create a ticker that ticks at given frequency
	ticker := time.NewTicker(durationUntilNextExecution)
	go func() {
		for {
			job(ecommerceService)
			<-ticker.C
		}
	}()
}
