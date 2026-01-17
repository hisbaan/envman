package config

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/hisbaan/envman/utils"
	"github.com/spf13/viper"
)

type Project struct {
	Path string   `mapstructure:"path"`
	Dirs []string `mapstructure:"dirs"`
}
type Config struct {
	Projects []Project `mapstructure:"projects"`
}

var v *viper.Viper
var once sync.Once

func getConfigDir() (string, error) {
	if home, err := os.UserHomeDir(); err != nil {
		return "", err
	} else {
		return filepath.Join(home, ".config", "envman"), nil
	}
}

func ensureConfigDir(dir string) error {
	return os.MkdirAll(dir, 0755)
}

func InitConfig() error {
	dir, err := getConfigDir()
	if err != nil {
		return fmt.Errorf("unable to get config directory %w", err)
	}

	if err := ensureConfigDir(dir); err != nil {
		return fmt.Errorf("unable to create config directory %w", err)
	}

	v = viper.New()

	v.SetConfigName("config")
	v.SetConfigType("toml")
	v.AddConfigPath(dir)

	if err := v.ReadInConfig(); err != nil {
		var perr viper.ConfigFileNotFoundError
		if errors.As(err, &perr) {
			v.Set("projects", []Project{})
			v.SafeWriteConfig()
		} else {
			return fmt.Errorf("fatal error in config file %w", err)
		}
	}
	return nil
}

func GetViper() *viper.Viper {
	once.Do(func() {
		if v == nil {
			InitConfig()
		}
	})
	return v
}

func GetProjects() []Project {
	var projects []Project
	GetViper().UnmarshalKey("projects", &projects)
	return projects
}

var ErrProjectNotFound = errors.New("project not found in config")

func GetProject(path string) (Project, error) {
	for _, project := range GetProjects() {
		if project.Path == path {
			return project, nil
		}
	}
	return Project{}, ErrProjectNotFound
}

func GetCurrentProject() (Project, error) {
	if cwd, err := os.Getwd(); err != nil {
		return Project{}, err
	} else {
		return GetParentProject(cwd)
	}
}

func GetParentProject(dir string) (Project, error) {
	projects := GetProjects()
	for _, project := range projects {
		if dir == project.Path || strings.HasPrefix(dir, project.Path+"/") {
			return project, nil
		}
	}
	return Project{}, ErrProjectNotFound
}

func AddProject(path string) error {
	info, err := os.Stat(path)
	if err != nil {
		return fmt.Errorf("cannot access project path: %w", err)
	}
	if !info.IsDir() {
		return fmt.Errorf("project path is not a directory")
	}

	projects := GetProjects()
	projects = append(projects, Project{Path: path, Dirs: []string{path}})
	GetViper().Set("projects", projects)
	return GetViper().WriteConfig()
}

func RemoveProject(path string) error {
	projects := GetProjects()
	projects = utils.Filter(projects, func(project Project) bool { return project.Path != path })
	GetViper().Set("projects", projects)
	return GetViper().WriteConfig()
}

func AddDir(projectPath string, dir string) error {
	projects := GetProjects()
	for i := range projects {
		if projects[i].Path == projectPath {
			projects[i].Dirs = append(projects[i].Dirs, dir)
		}
	}
	GetViper().Set("projects", projects)
	return GetViper().WriteConfig()
}

func RemoveDir(projectPath, dir string) error {
	projects := GetProjects()
	for i := range projects {
		if projects[i].Path == projectPath {
			projects[i].Dirs = utils.Filter(projects[i].Dirs, func(d string) bool { return d != dir })
		}
	}
	GetViper().Set("projects", projects)
	return GetViper().WriteConfig()
}
