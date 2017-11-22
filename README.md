# mgo.model
A golang mgo.v2 model helper

## Usage

#### 1. define  `Simple` model 

```
packages models

var mongo *model.Model

func NewWithConfig(cfg *model.Config, logger gogo.Logger) {
	mongo = model.NewModel(cfg, logger)
}

func Model() *model.Model {
	return mongo
}

type _Simple struct{}

var (
	Simple *_Simple

	simpleCollection = "simples"
	simpleIndexes    = []mgo.Index{}
)

type SimpleModel struct {
	ID bson.ObjectId `bson:"_id"`

	Data  string `bson:"data1"`
	Data2 string `bson:"data2"`
}

func NewSimpleModel() *SimpleModel {
	return &SimpleModel{
		ID: bson.NewObjectId(),
	}
}

func (_ *_Simple) Find(id string) (sm *SimpleModel, err error) {
	if !bson.IsObjectIdHex(id) {
		return nil, model.ErrInvalidId
	}

	Simple.Query(func(c *mgo.Collection) {
		err = c.FindId(bson.ObjectIdHex(id)).One(&sm)
	})

	return
}

func (_ *_Simple) Insert(sm *SimpleModel) (err error) {
	Simple.Query(func(c *mgo.Collection) {
		err = c.Insert(sm)
	})
	return
}

func (_ *_Simple) Query(query func(c *mgo.Collection)) {
	Model().Query(simpleCollection, simpleIndexes, query)
}

```

#### Use `Simple` to insert and find record

```
func main() {
  // init mgoConfig, logger
  
	models.NewWithConfig(mgoConfig, logger)

	sm := models.NewSimpleModel()
	sm.Data = "data1"
	sm.Data2 = "data2"
	err := models.Simple.Insert(sm)
	fmt.Println(err) // should be nil

	foundModel, _ := models.Simple.Find(sm.ID.Hex())

	fmt.Println(foundModel.Data)  // should be data1
	fmt.Println(foundModel.Data2) // should be data2
}

```

details to check example folder.
