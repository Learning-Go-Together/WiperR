package closing_channels_go

func countReports(numSentCh chan int) int {
	reports := 0
	for true {
		sent, ok := <-numSentCh
		if !ok {
			break
		}
		reports += sent
	}

	return reports
}

// don't touch below this line

func sendReports(numBatches int, ch chan int) {
	for i := 0; i < numBatches; i++ {
		numReports := i*23 + 32%17
		ch <- numReports
	}
	close(ch)
}
