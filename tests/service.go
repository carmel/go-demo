package tests

type db struct{}

// DB is fake Database interface.
type DB interface {
	FetchMessage(Lang string) (string, error)
	FetchDefaultMessage() (string, error)
}

type Greeter struct {
	Database DB
	Lang     string
}

// GreeterService is service to greet your friends.
type GreeterService interface {
	Greet() string
	GreetInDefaultMsg() string
}

func (d *db) FetchMessage(Lang string) (string, error) {
	// in real life, this code will call an external db
	// but for this sample we will just return the hardcoded example value
	if Lang == "en" {
		return "hello", nil
	}
	if Lang == "es" {
		return "holla", nil
	}
	return "bzzzz", nil
}

func (d *db) FetchDefaultMessage() (string, error) {
	return "default message", nil
}
func (g Greeter) Greet() string {
	msg, _ := g.Database.FetchMessage(g.Lang) // call Database to get the message based on the Lang
	return "Message is: " + msg
}

func (g Greeter) GreetInDefaultMsg() string {
	msg, _ := g.Database.FetchDefaultMessage() // call Database to get the default message
	return "Message is: " + msg
}
func NewDB() DB {
	return new(db)
}

func NewGreeter(db DB, Lang string) GreeterService {
	return Greeter{db, Lang}
}
