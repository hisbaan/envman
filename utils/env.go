package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func GetEnvFiles(dir string) ([]string, error) {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}
	files := Map(entries, func(val os.DirEntry) string { return val.Name() })
	return Filter(files, func(val string) bool { return strings.HasPrefix(val, ".env.") }), nil
}

func GetEnvs(dir string) ([]string, error) {
	files, err := GetEnvFiles(dir)
	if err != nil {
		return nil, err
	}
	return Map(files, func(val string) string { return val[5:] }), nil
}

func GetCommonEnvs(dirs []string) ([]string, error) {
	dirEnvs := make([][]string, len(dirs))
	for i, dir := range dirs {
		val, err := GetEnvs(dir)
		if err != nil {
			return nil, err
		}
		dirEnvs[i] = val
	}

	counts := make(map[string]int)
	for _, envs := range dirEnvs {
		for _, env := range envs {
			if val, ok := counts[env]; ok {
				counts[env] = val + 1
			} else {
				counts[env] = 1
			}
		}
	}

	envs := []string{}
	for env, count := range counts {
		if count == len(dirEnvs) {
			envs = append(envs, env)
		}
	}

	return envs, nil
}

func GetActiveEnv(dirs []string) (string, error) {
	if len(dirs) == 0 {
		return "", nil
	}

	names := []string{}
	for _, dir := range dirs {
		result, err := os.Readlink(dir + "/.env")
		if err != nil {
			if os.IsNotExist(err) {
				return "", nil
			}
			return "", fmt.Errorf("error reading symlink %s: %w", dir+"/.env", err)
		}
		prefix := filepath.Join(dir + ".env.")
		names = append(names, result[len(prefix):])
	}

	for _, name := range names {
		if name != names[0] {
			return "", fmt.Errorf("invalid state, mismatching symlinks")
		}
	}

	return names[0], nil
}
