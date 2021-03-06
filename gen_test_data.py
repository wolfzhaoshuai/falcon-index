#!/env/bin/python

import sys
import json


def gen(count):
    print "[",
    for x in range(0, count):
        d = {
            "endpoint": "laiwei-test%s" % (x,),
            "metric": "cpu.idle",
            "value": 1,
            "step": 30,
            "counterType": "GAUGE",
            "tags": "home=bj,srv=falcon",
            "timestamp": 1234567
        }
        print json.dumps(d),

        if x < count-1:
            print ","

    print "]"


def gen_zys(count):
    print "[",
    for x in range(0, count):
        d = {
            "endpoint": "wolf-zys",
            "metric": "cpu.idle",
            "value": 1,
            "step": 30,
            "counterType": "GAUGE",
            "tags": "home=bj,srv=falcon,id=%s" % str(x),
            "timestamp": 1234567
        }
        print json.dumps(d),

        if x < count-1:
            print ","

    print "]"

if __name__ == "__main__":
    gen_zys(int(sys.argv[1]))
