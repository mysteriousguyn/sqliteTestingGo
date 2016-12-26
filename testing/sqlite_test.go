package testing_test

import (
	"os"
	"sqliteTesting/sqlite"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Sqlite Database Testing", func() {

	Context("Here we create the new database connection", func() {
		var value bool

		It("if connection is created", func() {
			_, value = sqlite.CreateConnection()
			Expect(value).To(Equal(true))
		})
	})

	Context("Here we create the new Database table", func() {
		var query string
		var err error
		BeforeEach(func() {
			os.Remove("/.nilav.db")
			query = "create table if not exists employee (id INTEGER PRIMARY KEY   AUTOINCREMENT, name TEXT NOT NULL, address TEXT NOT NULL)"
		})

		JustBeforeEach(func() {
			db, _ := sqlite.CreateConnection()
			_, err = db.Exec(query)
		})

		It("create table success or not", func() {
			Expect(err).NotTo(HaveOccurred())
		})
	})

	Context("Testing to insert into table", func() {
		var query string
		var err error

		BeforeEach(func() {
			os.Remove("/nilav.db")
			query = "INSERT INTO employee(name, address) values(?,?)"
		})

		JustBeforeEach(func() {
			db, _ := sqlite.CreateConnection()
			stmt, _ := db.Prepare(query)
			_, err = stmt.Exec("nilav", nil)
		})

		It("Insertion successful or not", func() {
			Expect(err).NotTo(HaveOccurred())
		})
	})

})
