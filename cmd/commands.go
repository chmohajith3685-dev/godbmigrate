package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

const migrationsDir = "migrations"

var newCmd = &cobra.Command{
	Use:   "new [name]",
	Short: "Create a new migration",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]
		timestamp := time.Now().Format("20060102150405")

		if err := os.MkdirAll(migrationsDir, 0755); err != nil {
			fmt.Printf("Error creating migrations directory: %v\n", err)
			return
		}

		upFile := filepath.Join(migrationsDir, fmt.Sprintf("%s_%s.up.sql", timestamp, name))
		downFile := filepath.Join(migrationsDir, fmt.Sprintf("%s_%s.down.sql", timestamp, name))

		if err := createFile(upFile); err != nil {
			fmt.Printf("Error creating up file: %v\n", err)
			return
		}
		if err := createFile(downFile); err != nil {
			fmt.Printf("Error creating down file: %v\n", err)
			return
		}

		fmt.Printf("Created: %s\n", upFile)
		fmt.Printf("Created: %s\n", downFile)
	},
}

func createFile(path string) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	return f.Close()
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all migrations",
	Run: func(cmd *cobra.Command, args []string) {
		files, err := os.ReadDir(migrationsDir)
		if err != nil {
			if os.IsNotExist(err) {
				fmt.Println("No migrations folder found.")
				return
			}
			fmt.Printf("Error reading migrations: %v\n", err)
			return
		}

		var fileNames []string
		for _, f := range files {
			if !f.IsDir() {
				fileNames = append(fileNames, f.Name())
			}
		}

		sort.Strings(fileNames)

		if len(fileNames) == 0 {
			fmt.Println("No migrations found.")
			return
		}

		fmt.Println("Migrations:")
		for _, name := range fileNames {
			fmt.Printf("- %s\n", name)
		}
	},
}

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Show current migration status",
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := initDB(); err != nil {
			return err
		}
		defer store.Close()

		version, err := store.GetLatestVersion()
		if err != nil {
			return err
		}

		if version == 0 {
			fmt.Println("No migrations have been applied yet.")
		} else {
			fmt.Printf("Current migration version: %d\n", version)
		}

		return nil
	},
}

type migrationFile struct {
	version int64
	name    string
	path    string
}

var upCmd = &cobra.Command{
	Use:   "up",
	Short: "Apply all pending migrations",
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := initDB(); err != nil {
			return err
		}
		defer store.Close()

		currentVersion, err := store.GetLatestVersion()
		if err != nil {
			return err
		}

		files, err := os.ReadDir(migrationsDir)
		if err != nil {
			return fmt.Errorf("could not read migrations directory: %w", err)
		}

		var pending []migrationFile
		for _, f := range files {
			if f.IsDir() || !strings.HasSuffix(f.Name(), ".up.sql") {
				continue
			}

			parts := strings.Split(f.Name(), "_")
			if len(parts) < 2 {
				continue
			}

			version, err := strconv.ParseInt(parts[0], 10, 64)
			if err != nil {
				continue
			}

			if version > currentVersion {
				pending = append(pending, migrationFile{
					version: version,
					name:    f.Name(),
					path:    filepath.Join(migrationsDir, f.Name()),
				})
			}
		}

		sort.Slice(pending, func(i, j int) bool {
			return pending[i].version < pending[j].version
		})

		if len(pending) == 0 {
			fmt.Println("No pending migrations to apply.")
			return nil
		}

		for _, m := range pending {
			fmt.Printf("Applying migration: %s... ", m.name)

			content, err := os.ReadFile(m.path)
			if err != nil {
				return fmt.Errorf("could not read migration file %s: %w", m.name, err)
			}

			if err := store.ApplyMigration(m.version, string(content)); err != nil {
				fmt.Println("FAILED")
				return err
			}

			fmt.Println("OK")
		}

		fmt.Println("All pending migrations applied successfully.")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(newCmd)
	rootCmd.AddCommand(listCmd)
	rootCmd.AddCommand(statusCmd)
	rootCmd.AddCommand(upCmd)
}
