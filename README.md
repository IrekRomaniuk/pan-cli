### PaloAlto firewall PAN-OS cli utils

pan-cli load file.csv - load address objects from file

pan-cli create "test_address" ip "1.2.3.4/32"
pan-cli create "test_service" port "123"

panc-cli find keyword - find keyword, i.e. defunct in the results of show system resources | match defunct