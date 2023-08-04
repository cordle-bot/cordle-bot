//go:build integration

package database

import (
	"cordle/internal/config"
	"cordle/internal/users"
	"errors"
	"fmt"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

var d *Db

func TestMain(m *testing.M) {
	d = NewDb(config.Config.Database)
	m.Run()
}

func TestPing(t *testing.T) {
	err := d.Ping()
	assert.NoError(t, err)
}

func TestDb(t *testing.T) {
	assert.NotNil(t, d)
}

func TestAddUser(t *testing.T) {
	u := users.User{
		Id:     "7567",
		Wins:   20,
		Losses: 53,
		Draws:  151,
		Elo:    341,
	}

	err := d.AddUser(u.Id)
	assert.NoError(t, err)

	e, err := d.CheckUser(u.Id)
	assert.NoError(t, err)
	if !e {
		log.Fatalln(errors.New("failed to add user"))
	}
	err = d.DeleteUser(u.Id)
	assert.NoError(t, err)
}

func TestAddUsers(t *testing.T) {
	u1 := users.User{
		Id:     "7567",
		Wins:   20,
		Losses: 53,
		Draws:  151,
		Elo:    341,
	}
	u2 := users.User{
		Id:     "1577",
		Wins:   20,
		Losses: 13,
		Draws:  51,
		Elo:    541,
	}

	u := make([]users.User, 2)
	u[0] = u1
	u[1] = u2
	d.AddUsers(&u)

	e, err := d.CheckUser(u1.Id)
	assert.NoError(t, err)
	if !e {
		log.Fatalln(errors.New("failed to add user"))
	}
	d.DeleteUser(u1.Id)

	e, err = d.CheckUser(u2.Id)
	assert.NoError(t, err)
	if !e {
		log.Fatalln(errors.New("failed to add user"))
	}
	d.DeleteUser(u2.Id)
}

func TestUpdateUser(t *testing.T) {
	d.AddUser("123")
	e, err := d.CheckUser("123")
	assert.NoError(t, err)
	assert.True(t, e)

	u, err := d.ReadUser("123")
	assert.NoError(t, err)
	draws := u.Draws
	u.Draws += 1
	d.UpdateUser(u)
	u, err = d.ReadUser("123")
	assert.NoError(t, err)
	if u.Draws != draws+1 {
		log.Fatalln(errors.New("error updating draw count"))
	}

	err = d.DeleteUser("123")
	assert.NoError(t, err)
}

func TestUpdateUsers(t *testing.T) {

}

func TestReadUser(t *testing.T) {
	d.AddUser("123")

	u, err := d.ReadUser("123")
	assert.NoError(t, err)
	if u.Id != "123" {
		log.Fatalln(errors.New("read nil user error"))
	}
	assert.NotNil(t, u)
}

func TestReadAllUsers(t *testing.T) {
	d.AddUser("123")
	d.AddUser("456")

	u, err := d.ReadAllUsers()
	assert.NoError(t, err)
	assert.NotNil(t, u)

	if len(u) == 0 {
		log.Fatalln(fmt.Errorf(
			"incorrect array length for all users %d", len(u)),
		)
	}
	err = d.DeleteUsers([]string{"123", "456"})
	assert.NoError(t, err)

	u, err = d.ReadAllUsers()
	assert.NoError(t, err)
	assert.Zero(t, len(u))
}

func TestReadTop(t *testing.T) {
	tt, err := d.ReadTop()
	assert.NoError(t, err)
	assert.NotNil(t, tt)

	for i := 0; i < len(tt)-2; i++ {
		if tt[i].Elo < tt[i+1].Elo {
			continue
		}
		log.Fatalln(fmt.Errorf(
			"top ten wrong order %d", len(tt)),
		)
	}
}

func TestCheckUser(t *testing.T) {
	d = NewDb(config.Config.Database)
	defer d.Close()

	e, err := d.CheckUser("69")
	assert.NoError(t, err)
	assert.False(t, e)

	// Temporary until new API, create a user for this test
	err = d.AddUser("123")
	assert.NoError(t, err)
	e, err = d.CheckUser("123")
	assert.NoError(t, err)
	assert.True(t, e)

	// Delete the temporary user
	err = d.DeleteUser("123")
	assert.NoError(t, err)
}

func TestDeleteUser(t *testing.T) {
	u := users.User{
		Id:     "61567",
		Wins:   22,
		Losses: 51,
		Draws:  101,
		Elo:    371,
	}

	d = NewDb(config.Config.Database)
	defer d.Close()

	err := d.AddUser(u.Id)
	assert.NoError(t, err)

	err = d.DeleteUser(u.Id)
	assert.NoError(t, err)

	e, err := d.CheckUser(u.Id)
	assert.NoError(t, err)
	if e {
		log.Fatalln(errors.New("failed to delete user"))
	}
}

func TestDeleteUsers(t *testing.T) {

}

func TestCreateTable(t *testing.T) {
	d = NewDb(config.Config.Database)
	defer d.Close()

	c := `
		create table test_table(
			id int not null primary key,
			name varchar(10)
		);
	`

	d.CreateTable(c)
	e := d.CheckTable("test_table")
	assert.True(t, e)

	d.DeleteTable("test_table")
}

func TestUpdateTable(t *testing.T) {

}

func TestDeleteTable(t *testing.T) {
	d = NewDb(config.Config.Database)
	defer d.Close()

	c := `
		create table test_table(
			id int not null primary key,
			name varchar(10)
		);
	`

	d.CreateTable(c)
	e := d.CheckTable("test_table")
	assert.True(t, e)

	d.DeleteTable("test_table")
	e = d.CheckTable("test_table")
	assert.False(t, e)
}

func TestCheckTable(t *testing.T) {
	d = NewDb(config.Config.Database)
	defer d.Close()

	e := d.CheckTable("check_table")
	assert.False(t, e)

	c := `
		create table check_table(
			id int not null primary key
		);
	`

	d.CreateTable(c)
	e = d.CheckTable("check_table")
	assert.True(t, e)

	d.DeleteTable("check_table")
	e = d.CheckTable("check_table")
	assert.False(t, e)
}
