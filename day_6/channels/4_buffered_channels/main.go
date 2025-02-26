package buffered_channels

func addEmailsToQueue(emails []string) chan string {
	ch := make(chan string, len(emails))
	for _, m := range emails {
		ch <- m
	}

	return ch
}
