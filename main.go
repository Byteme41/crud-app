package main

import (
	"Uni-Project/database"
	"log"

	"github.com/spf13/cobra"
)

func main() {
    db, err := database.Connect()
    if err != nil {
        log.Fatalf("database connection err: %v", err)
    }
    defer db.Close()

    var id int
    var name string

    var add = &cobra.Command{
        Use: "add an entry",
        Short: "add an entry to database",
        Run: func(cmd *cobra.Command, args []string) {
            database.Add(db, id, name)
        },
    }

    var get = &cobra.Command{
        Use: "get an entry",
        Short: "get an entry from the database",
        Run: func(cmd *cobra.Command, args []string) {
            database.Get(db, id)
        },
    }

    var all = &cobra.Command{
        Use: "all entries",
        Short: "all entries from the database",
        Run: func(cmd *cobra.Command, args []string) {
            database.GetAll(db)
        },
    }

    var update = &cobra.Command{
        Use: "update an entry",
        Short: "update an existing entry from the database",
        Run: func(cmd *cobra.Command, args []string) {
            database.Update(db, id, name)
        },
    }

    var remove = &cobra.Command{
        Use: "remove an entry",
        Short: "remove an entry from the database",
        Run: func(cmd *cobra.Command, args []string) {
            database.Delete(db, id)
        },
    }

    add.Flags().IntVarP(&id, "id", "i", 1, "id of the student")
    add.Flags().StringVarP(&name, "name", "n", "john doe" , "name of the student")

    get.Flags().IntVarP(&id, "id", "i", 1 , "get a specific id from the db")

    update.Flags().IntVarP(&id, "id", "i", 1 , "id")
    update.Flags().StringVarP(&name, "name", "n", "john doe" , "name")

    remove.Flags().IntVarP(&id, "id", "i", 1 , "get a specific id from the db")

    var rootCmd = &cobra.Command{Use: "main"}
    rootCmd.AddCommand(add)
    rootCmd.AddCommand(get)
    rootCmd.AddCommand(all)
    rootCmd.AddCommand(update)
    rootCmd.AddCommand(remove)
    rootCmd.Execute()
}
