

build:
	cd proto && protoc --go_out=../protogen/golang --go_opt=paths=source_relative ./google/api/*.proto \
 	&& protoc --go_out=../protogen/golang --go_opt=paths=source_relative ./users/*.proto


gw_build:
	cd proto \
	&& protoc --go_out=../protogen/golang --go_opt=paths=source_relative --go-grpc_out=../protogen/golang --go-grpc_opt=paths=source_relative --grpc-gateway_out=../protogen/golang --grpc-gateway_opt paths=source_relative --grpc-gateway_opt generate_unbound_methods=true .\google\api\*.proto \
	&& protoc --go_out=../protogen/golang --go_opt=paths=source_relative --go-grpc_out=../protogen/golang --go-grpc_opt=paths=source_relative --grpc-gateway_out=../protogen/golang --grpc-gateway_opt paths=source_relative --grpc-gateway_opt generate_unbound_methods=true ./users/*.proto

clean:
	rmdir /s /q protogen && mkdir protogen && cd protogen && mkdir golang


build_v2:
	cd proto && protoc --go_out=../protogen/golang --go_opt=paths=source_relative --go-grpc_out=../protogen/golang --go-grpc_opt=paths=source_relative --grpc-gateway_out=../protogen/golang --grpc-gateway_opt paths=source_relative --grpc-gateway_opt generate_unbound_methods=true ./users/*.proto

#cd proto && protoc --go_out=../protogen/golang --go_opt=paths=source_relative --go-grpc_out=../protogen/golang --go-grpc_opt=paths=source_relative ./users/*.proto

install:
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2 \
	&& go get google.golang.org/grpc \
	&& go get github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway \
    && go get github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2 \
    && go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway \
    && go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2