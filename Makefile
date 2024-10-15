local_setup_up:
	 docker-compose -f docker-compose.yml up -d

local_setup_down:
	 docker-compose -f docker-compose.yml down -v

test:
	go test -cover ./...

run:
	go run main.go

run_lambda_receiver:
	 cd cmd/productPromotionLambda && go run main.go