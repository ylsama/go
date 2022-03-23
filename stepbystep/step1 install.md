# go install skip 
## go environment
```bash
cd $HOME
mkdir workspace
cd workspace
mkdir go
cd go
mkdir bin
echo "export GO_PATH=${HOME}/workspace/go" >> $HOME/.bashrc
echo "export GOBIN=${HOME}/workspace/go/bin" >> $HOME/.bashrc
go env -w GOBIN=/home/ylongkaka/workspace/go/bin
cd $HOME/workspace/go
```
