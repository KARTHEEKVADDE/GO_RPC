

// GO PROTOBUF COMPILER
choco install protoc

// Install Dependencies
go get -u github.com/golang/protobuf/protoc-gen-go

go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28

go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2

//Export Path
export PATH="$PATH:$(go env GOPATH)/bin"

// Inside PROTO file
package chat;
option go_package = "/chat";

// Run the Command
protoc -I=. --go_out=chat chat.proto
protoc -I=. --go_out=. --go-grpc_out=. chat.proto

protoc -I=$SRC_DIR --go_out=$DST_DIR $SRC_DIR/addressbook.proto


C:\Users\Guser2\go\src\github.com\kartheekvadde\GO_GRPC\chat

protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative chat/chat.proto
