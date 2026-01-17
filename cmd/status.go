package cmd

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/fatih/color"
	"github.com/hisbaan/envman/config"
	"github.com/hisbaan/envman/utils"
	"github.com/jwalton/go-supportscolor"
	"github.com/spf13/cobra"
)

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Get the status of your current project",
	RunE: func(cmd *cobra.Command, args []string) error {
		project, err := config.GetCurrentProject()
		if err != nil {
			return err
		}
		envs, err := utils.GetCommonEnvs(project.Dirs)
		if err != nil {
			return err
		}
		localPaths := []string{}
		for _, dir := range project.Dirs {
			path, err := filepath.Rel(project.Path, dir)
			if err != nil {
				return err
			}
			localPaths = append(localPaths, path)
		}

		activeEnv, err := utils.GetActiveEnv(project.Dirs)
		if err != nil {
			return err
		}

		fmt.Printf("project: %s\n", project.Path)
		fmt.Printf("dirs:    %s\n", strings.Join(localPaths, ", "))
		fmt.Printf("envs:    %s\n", strings.Join(
			utils.Map(
				envs,
				func(env string) string {
					if env == activeEnv {
						if supportscolor.Stdout().SupportsColor {
							return fmt.Sprintf("%s", color.GreenString(env))
						}
						return fmt.Sprintf("*%s", env)
					}
					return env
				},
			),
			", "))

		return nil
	},
}

func init() {
	rootCmd.AddCommand(statusCmd)
}
