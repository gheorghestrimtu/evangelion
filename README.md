# evangelion

#Project setup

1. clone project
2. open in goland
3. go to goland preferences
4. search 'modules'
5. select 'Enable modules configuration'
6. go to terminal
7. cd to project directory
8. run `go mod download`
9. run `go mod tidy`

#Running tests

1. config/config.yml contains blockchain config
2. tests live in contracts/contract_test.go
3. to run tests go to contracts/contract_suite_go and click on the green arrow left to `func TestContracts()`
,there is also the possibility to run the tests from the terminal
4. If you want to run just one test, put a `F` in front of an `It`. So `It` becomes
`FIt`. `F` stands for Focus
