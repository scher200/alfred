package alfred

import (
	"fmt"
	"strings"
)

// MagicTaskURL will parse "magic" tasks, as denoted by "__"
func MagicTaskURL(task string) string {
	url, _ := TaskParser(task, "alfred:list")
	if strings.HasPrefix(url, "http") {
		return url + ":"
	}
	return ""
}

// TaskParser returns a url and a task(default if necessary) (location, task)
func TaskParser(task, defaultTask string) (string, string) {
	// does it start with http?
	if strings.HasPrefix(task, "http") && strings.Contains(task, "://") {
		// we have to get the http: colon out of the way :(
		bits := strings.SplitN(task, ":", 3)
		url := strings.Join(bits[0:2], ":")
		if len(bits) >= 3 {
			// alright, so we have tasks and args ...
			return url, bits[2]
		}
		return url, defaultTask
	}

	// is this a catalog?
	if isCatalog(task) {
		bits := strings.Split(task, ":")
		if len(bits) >= 2 {
			return bits[0], bits[1]
		}
		return bits[0], defaultTask
	}

	// does it start with a /? AND not a local directory? Aka a remote task?
	if strings.HasPrefix(task, "/") {
		bits := strings.Split(task, ":")
		url := fmt.Sprintf("https://raw.githubusercontent.com/kcmerrill/alfred/master/remote-modules%s.yml", bits[0])

		if len(bits) >= 2 {
			return url, bits[1]
		}
		return url, defaultTask
	}

	// lets check if this is a github file(and not a local ./ task)
	if !strings.HasPrefix(task, "./") && strings.Contains(task, "/") {
		bits := strings.Split(task, ":")
		url := fmt.Sprintf("https://raw.githubusercontent.com/%s/master/alfred.yml", bits[0])
		if len(bits) >= 2 {
			return url, bits[1]
		}
		return url, defaultTask
	}

	if task == "" {
		return "./", defaultTask
	}

	if strings.HasPrefix(task, "!") {
		return "./", "!exec.command"
	}

	// alright, so it's not a url, it's not a github repo, it must be just a regular local task
	bits := strings.Split(task, ":")
	if len(bits) >= 2 {
		return bits[0], bits[1]
	}
	return "./", bits[0]
}
