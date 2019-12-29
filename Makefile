BUILD=go build 
FILENAME=gorename
DIR=build

all:
	$(BUILD) -o $(DIR)/$(FILENAME)_osx_x64
	#env GOOS=darwin GOARCH=amd64 $(BUILD) -o $(DIR)/$(FILENAME)_osx_64	
	env GOOS=windows GOARCH=amd64 $(BUILD) -o $(DIR)/$(FILENAME)_win_64
	env GOOS=windows GOARCH=386 $(BUILD) -o $(DIR)/$(FILENAME)_win_32
	env GOOS=linux GOARCH=amd64 $(BUILD) -o $(DIR)/$(FILENAME)_linux_64
	env GOOS=linux GOARCH=arm $(BUILD) -o $(DIR)/$(FILENAME)_linux_arm
