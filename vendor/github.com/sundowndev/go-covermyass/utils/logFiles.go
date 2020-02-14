package utils

var fileProcessor = &FileProcessor{}

/*
 * Unmock	Function that delete any existing symbolic link
 */
func Unmock(files []string) {
	for _, path := range files {
		fileProcessor.Register(path)
	}

	fileProcessor.Proceed(func(path string) {
		// rm -rf /var/log/auth.log
		// echo "" > /var/log/auth.log
		LoggerService.Success("Delete symbolic link for " + path + " to /dev/null")
	})
}

/*
 * Mock	Function that transform log files into symbolic links
 */
func Mock(files []string) {
	for _, path := range files {
		fileProcessor.Register(path)
	}

	fileProcessor.Proceed(func(path string) {
		// ln ...
		LoggerService.Success("Create symbolic link for " + path + " to /dev/null")
	})
}
