.PHONY: all
USER=username
FILENAME=go.mod
FILE_PATH=./src/step9
FILE_EXIST=$(shell ls ${FILE_PATH} |grep ${FILENAME})

init:
ifeq ("$(shell echo ${USER})","$(shell echo username)")
	echo "Invalid username"
	exit 1
endif

ifeq ("$(shell echo ${FILENAME})" , "$(shell echo ${FILE_EXIST})")
	rm $(FILE_PATH)/$(FILENAME)
else
	echo  $(FILE_EXIST)
endif

	cd ./src/step9 && go mod init github.com/$(USER)/go-study/step9