package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CrearBdNoticiasCrud_20240203_162222 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CrearBdNoticiasCrud_20240203_162222{}
	m.Created = "20240203_162222"

	migration.Register("CrearBdNoticiasCrud_20240203_162222", m)
}

// Run the migrations
func (m *CrearBdNoticiasCrud_20240203_162222) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update

}

// Reverse the migrations
func (m *CrearBdNoticiasCrud_20240203_162222) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update

}
