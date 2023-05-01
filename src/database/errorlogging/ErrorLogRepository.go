package errorlogging

import (
	"EspressoTracking/database"
	"context"
)

type ErrorLogRepository struct {
	databaseName string
	collection   string
	ctx          database.DatabaseContext
}

// new instance
func NewErrorLogRepository() *ErrorLogRepository {
	return &ErrorLogRepository{
		databaseName: "ErrorLogging",
		collection:   "Logs",
		ctx:          database.DatabaseContext{},
	}

}

func (elr *ErrorLogRepository) CreateErrorLogRecord(errorLogRecord ErrorLogRecord) {
	dbContext, err := elr.ctx.NewConnection()
	if err != nil {
		panic(err)
	}
	defer elr.ctx.Disconnect()

	collection := dbContext.Database(elr.databaseName).Collection(elr.collection)
	_, err = collection.InsertOne(context.TODO(), errorLogRecord)
	if err != nil {
		panic(err)
	}
}
