#!/bin/bash

#Use OpenSSL cli to decode base64 text and decrypt using aes-ecb with Hex encoded Key of "YELLOW SUBMARINE"
openssl enc -base64 -d -aes-128-ecb -K 59454c4c4f57205355424d4152494e45 -in 7.txt


