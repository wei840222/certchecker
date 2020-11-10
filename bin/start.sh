#!/bin/bash
cd /data/helen/certchecker/ ; nohup go run main.go > log/cert-`date +%Y%m%d%H%M%S`.log & 
