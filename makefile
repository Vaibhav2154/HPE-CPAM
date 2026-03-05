build:
	go build -ldflags="-X 'github.com/aneeshsunganahalli/HPE-CPAM/cmd.Version=1.0.0'" -o cpam main.go

install:
	go install