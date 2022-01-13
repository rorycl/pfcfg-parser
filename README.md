# pfcfg-parser

version 0.0.1 : 13 January 2022

A quick and simple parser for PFSense XML configuration files to
generate a plain text file of the main configuration items, focusing on
aliases and firewall rules, useful for auditing.

The output can be easily modified by altering the output template. For
reference to the naming of struct fields, refer to pfconfig.go.

Build

    go build

Run

    ./pfcfg-parser config-20220112112427.xml

Output

    PFSense Configuration Report

    Report generated at 2022-01-13 17:30
    System      mycorp-fwl17b
    Version     21.7

    Interfaces
    WAN on igb3 ip xx.xx.xx.xx
    LAN on igb4 ip yy.yy.yy.yy

    Virtual IPs
    carp wan WAN CARP zz.zz.zz.zz

    Aliases
    hostA host aa.bb.cc.25
    web   host aa.bb.cc.20 aa.bb.cc.21

    Filters
    2581947622 2021-02-17 wan hostA to any
    2425581024 2019-03-05 wan any to hostB:25
    2425581120 2019-03-05 wan any to hostB:587
    2425581216 2019-03-05 wan any to hostC:443

