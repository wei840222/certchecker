#!/bin/bash
a=`ss -antlp | grep 8080 | awk -F ',' '{print $2}'| awk -F '=' '{print $2}'`
kill $a
