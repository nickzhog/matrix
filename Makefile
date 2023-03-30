bench:
	@go test -count=1 -run=^$$ -bench=. -benchmem ./... | grep -vE "^(\?|ok|PASS) " ; true