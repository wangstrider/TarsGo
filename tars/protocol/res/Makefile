all:
	tars2go -add-servant=false -tarsPath github.com/wangstrider/TarsGo/tars *.tars
	#override file
	sed -i 's|"endpointf"|"github.com/wangstrider/TarsGo/tars/protocol/res/endpointf"|g' queryf/QueryF_IF.go
