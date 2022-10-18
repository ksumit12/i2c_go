# i2c_go
Golang code for i2c communication between CM4 and arduino
### the main.go will read any i2c signal on gpio1 and gpio2 and if the read value is 30 then it will write it 22 
this was specific to my use case
# i2c bus for cm4 is bus 1 
device address for me was 0x04 
if in any case code doesn't that is probabily bus or device address may be incorrect 
