########## build ##########
.Phony: build
build:
	GOOS=linux GOARCH=amd64 go build -o dist/server ./cmd/main.go
	scp -i ~/.ssh/fn dist/server ubuntu@ec2-44-203-175-111.compute-1.amazonaws.com:~/
	ssh -i ~/.ssh/fn ubuntu@ec2-44-203-175-111.compute-1.amazonaws.com "cd /data/www/bevm/; sudo mv ~/server ./; sudo systemctl restart bevm"
