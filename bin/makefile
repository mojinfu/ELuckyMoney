RM= rm -rf

all: clean release

release: ELuckyMoney

clean:
	$(RM) ../log/* 
	$(RM)  ELuckyMoney

ELuckyMoney:
	go build -ldflags "-s -w" -o $@ ../src/main/*.go
