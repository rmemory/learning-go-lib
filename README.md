# learning-go-lib

To manually run tests ...

go test ./... -v

(1) Install mockery so that clients (like learning-go-api) can mock functionality provided by learning-go-lib

go install github.com/vektra/mockery/v2@v2.39.1

which mockery
/usr/local/go/bin/mockery

mockery --all --inpackage --case snake

Now there is new struct called MockInterface. This struct is the receiver of the mock implementation of the NewUuid function.

(2)Install act so that we can run GitHub Actions locally. 

go install github.com/nektos/act@latest

Run act -h to ensure the installation is working.

(3) To locally simulate a "git push" run, "act push"