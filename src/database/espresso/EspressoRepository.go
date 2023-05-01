package espresso

import (
	"EspressoTracking/database"
	"context"
	"go.mongodb.org/mongo-driver/bson"
)

type EspressoRepository struct {
	databaseName string
	collection   string
	ctx          database.DatabaseContext
}

func NewEspressoRepository() *EspressoRepository {
	return &EspressoRepository{
		databaseName: "Espresso",
		collection:   "Brands",
		ctx:          database.DatabaseContext{},
	}
}

// Brand Queries
// TODO: Refactor these into a separate file
func (er *EspressoRepository) CreateBrand(brand Brand) error {
	dbContext, err := er.ctx.NewConnection()
	if err != nil {
		return err
	}
	defer er.ctx.Disconnect()

	collection := dbContext.Database(er.databaseName).Collection(er.collection)
	_, err = collection.InsertOne(context.TODO(), brand)
	if err != nil {
		return err
	}

	return nil
}

func (er *EspressoRepository) GetAllBrands() []Brand {
	dbContext, err := er.ctx.NewConnection()
	if err != nil {
		panic(err)
	}
	defer er.ctx.Disconnect()

	collection := dbContext.Database(er.databaseName).Collection(er.collection)
	cursor, err := collection.Find(context.TODO(), bson.D{})
	if err != nil {
		panic(err)
	}
	var brands []Brand
	for cursor.Next(context.TODO()) {
		var brand Brand
		err = cursor.Decode(&brand)
		if err != nil {
			panic(err)
		}
		brands = append(brands, brand)
	}
	return brands
}

func (er *EspressoRepository) GetBrandByName(brandName string) (*Brand, error) {
	dbContext, err := er.ctx.NewConnection()
	if err != nil {
		return nil, err
	}
	defer er.ctx.Disconnect()

	collection := dbContext.Database(er.databaseName).Collection(er.collection)
	var brand Brand
	err = collection.FindOne(context.TODO(), bson.D{
		{"name", brandName},
	}).Decode(&brand)
	if err != nil {
		return nil, err
	}
	return &brand, nil
}

func (er *EspressoRepository) GetById(id string) (*Brand, error) {
	dbContext, err := er.ctx.NewConnection()
	if err != nil {
		return nil, err
	}
	defer er.ctx.Disconnect()

	collection := dbContext.Database(er.databaseName).Collection(er.collection)
	var brand Brand
	err = collection.FindOne(context.TODO(), bson.D{{"id", id}}).Decode(&brand)

	if err != nil {
		return nil, err
	}
	return &brand, nil
}

func (er *EspressoRepository) UpdateBrand(brands *Brand) (*Brand, error) {
	dbContext, err := er.ctx.NewConnection()
	if err != nil {
		return nil, err
	}
	defer er.ctx.Disconnect()

	collection := dbContext.Database(er.databaseName).Collection(er.collection)
	_, err = collection.UpdateOne(context.TODO(), bson.D{{"id", brands.Id}}, bson.D{
		{"$set", bson.D{
			{"name", brands.Name},
			{"coffee", brands.Coffee},
		}},
	})
	if err != nil {
		return nil, err
	}
	return brands, nil

}

func (er *EspressoRepository) DeleteBrand(id string) error {
	dbContext, err := er.ctx.NewConnection()
	if err != nil {
		return err
	}
	defer er.ctx.Disconnect()

	collection := dbContext.Database(er.databaseName).Collection(er.collection)
	_, err = collection.DeleteOne(context.TODO(), bson.D{{"id", id}})
	if err != nil {
		return err
	}

	return nil
}

// Coffee Queries
// TODO: Refactor these into a separate file
func (er *EspressoRepository) AddCoffee(brandId string, coffee Coffee) (*Brand, error) {
	dbContext, err := er.ctx.NewConnection()
	if err != nil {
		return nil, err
	}
	defer er.ctx.Disconnect()

	collection := dbContext.Database(er.databaseName).Collection(er.collection)
	_, err = collection.UpdateOne(context.TODO(), bson.D{{"id", brandId}}, bson.D{
		{"$push", bson.D{
			{"coffee", coffee},
		}},
	})
	if err != nil {
		return nil, err
	}
	return er.GetById(brandId)
}

func (er *EspressoRepository) DeleteCoffee(brandId string, coffeeId string) (*Brand, error) {
	dbContext, err := er.ctx.NewConnection()
	if err != nil {
		return nil, err
	}
	defer er.ctx.Disconnect()

	collection := dbContext.Database(er.databaseName).Collection(er.collection)
	_, err = collection.UpdateOne(context.TODO(), bson.D{{"id", brandId}}, bson.D{
		{"$pull", bson.D{
			{"coffee", bson.D{
				{"id", coffeeId},
			}},
		}},
	})
	if err != nil {
		return nil, err
	}
	return er.GetById(brandId)
}

func (er *EspressoRepository) UpdateCoffee(brandId string, coffee Coffee) (*Brand, error) {
	dbContext, err := er.ctx.NewConnection()
	if err != nil {
		return nil, err
	}
	defer er.ctx.Disconnect()

	collection := dbContext.Database(er.databaseName).Collection(er.collection)
	_, err = collection.UpdateOne(context.TODO(), bson.D{{"id", brandId}, {"coffee.id", coffee.Id}}, bson.D{
		{"$set", bson.D{
			{"coffee.$.name", coffee.Name},
			{"coffee.$.price", coffee.Price},
			{"coffee.$.Description", coffee.Description},
			{"coffee.$.Tags", coffee.Tags},
		}},
	})
	if err != nil {
		return nil, err
	}
	return er.GetById(brandId)
}
