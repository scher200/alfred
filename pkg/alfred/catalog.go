package alfred

import (
	"strings"
	"time"
)

func isCatalog(task string) bool {
	return strings.HasPrefix(task, "@")
}

func updateCatalog(dir string, context *Context) {
	// catalogs at this point are just git repositories.
	// assume a git repository, and update
	context.lock.Lock()
	_, ok := execute("git pull", dir)
	context.lock.Unlock()

	// switch the directory
	context.rootDir = dir

	if ok {
		outOK("@catalog", "updated!", context)
		return
	}
	context.lock.Lock()
	outWarn("@catalog", "Unable to update the catalog. It could be out of date.", context)
	<-time.After(3 * time.Second)
	context.lock.Unlock()
}
