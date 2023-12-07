.PHONY: idl_gen

userhub:
	make idl_gen app=userhub

demo:
	 make idl_gen app=demo

idl_gen:
	cd idl && protoc --go_out=../idl_gen --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative --go-grpc_out=../idl_gen ${app}/**



