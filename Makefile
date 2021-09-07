build:
	cd frontend && npm run build
	@cp frontend/public $$CTPath -r
	@cd backend && ./build.bash

install:
	cd frontend && npm install && npm run build
	@cp frontend/public $$CTPath -r
	@cd backend && go build -o "$(target)" .
	make clean

run:
	cd backend && go build -o main .
	cd backend/sampletests && ../main

serve_frontend:
	cd frontend && npm run dev

clean:
	@cd backend; [ -f main ] && rm "main" || :
	@cd backend/sampletests; [ -f main-c ] && rm main-c || :
	@cd backend/sampletests; [ -f main-cpp ] && rm main-cpp || :
	@cd backend/sampletests; [ -f main-go ] && rm main-go || :

	@cd backend; [ -d bin ] && rm bin -rf || :