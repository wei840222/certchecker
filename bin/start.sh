#!/bin/bash
cd /opt/helen/certchecker/ ; nohup go run main.go > log/cert-`date +%Y%m%d%H%M%S`.log & 
