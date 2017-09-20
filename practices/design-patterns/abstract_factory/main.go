package main

import "fmt"

// The abstract factory pattern provides a way to encapsulate a group of individual factories
// that have a common theme without specifying their concrete classes.[1]
// In normal usage, the client software creates a concrete implementation of the abstract factory
// and then uses the generic interface of the factory to create the concrete objects that are part of the theme.

type db struct {
	name     string
	connPool *pool
}

func (d *db) AddPool(p *pool) {
	d.connPool = p
}

func (d *db) String() string {
	return fmt.Sprintf("\nNAME:%s\nPOOL: %v", d.name, d.connPool)
}

type pool struct {
	connections []*connection
}

func (p *pool) AddConnection(c *connection) {
	p.connections = append(p.connections, c)
}

func (p *pool) String() string {
	var connectionValues string
	for i, v := range p.connections {
		connectionValues += fmt.Sprintf("%dth \n\tconnection: %s\n", i+1, v)
	}
	return connectionValues
}

type connection struct {
	connType string
}

func (c *connection) String() string {
	return c.connType
}

type dbFactory interface {
	NewDB(string) *db
	NewPool() *pool
	NewConnection(string) *connection
}

type myFactory struct{}

func (f myFactory) NewDB(dbName string) *db {
	return &db{
		name: dbName,
	}
}

func (f myFactory) NewPool() *pool {
	return &pool{}
}

func (f myFactory) NewConnection(connType string) *connection {
	return &connection{connType}
}

func main() {
	var factory dbFactory
	factory = myFactory{}

	// use factory to create db
	mysqlDB := factory.NewDB("mysql")
	// use factory to create db pool
	mysqlPool := factory.NewPool()
	mysqlDB.AddPool(mysqlPool)

	// use factory to create db connection
	mysqlConn := factory.NewConnection("cluster-connection")

	mysqlPool.AddConnection(mysqlConn)

	fmt.Println("mysqlDB value should be :", mysqlDB)
}
