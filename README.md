### PaloAlto firewall PAN-OS cli utils based on go-panos: github.com/scottdware/go-panos
```
pan-cli load --config pan-cli.yaml -p password -f file.csv -> load address objects from file

pan-cli create -u admin -p password -o "test_address" -v "1.2.3.4/32"
```
#### TO DO
```
pan-cli create -o "test_service" -v "123"  -> object type detected thru regex
```
panc-cli find keyword - find keyword, i.e. defunct in the results of 'show system resources | match defunct'