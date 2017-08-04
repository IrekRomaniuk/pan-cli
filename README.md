### PaloAlto firewall (PAN-OS) cli utils based on http://github.com/scottdware/go-panos
#### More at my blog at http://medium.com/@IrekRomaniuk

```
pan-cli load --config pan-cli.yaml -p password -f file.csv -> load address objects from file

pan-cli create -u admin -p password -o "test_address" -v "1.2.3.4/32"

pan-cli commit -d "192.168.3.101" -p 'password'             // Firewall
pan-cli commit -d "192.168.3.100" -p 'password' -g 'DC2'    // Panorama device group 'DC2'
```

#### config file (flags take precedence)
```
login: admin
device: 192.168.3.101
```
##### TO DO:
```
pan-cli create -o "test_service" -v "123"  -> object type detected thru regex

panc-cli find keyword - find keyword, i.e. defunct in the results of 'show system resources | match defunct'
```