import random
import os
import sys

chars = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ!#%&()*+-<=>?@^_"

password = ''.join(random.choice(chars) for i in range(12))

target = "target/qito"

filepath = "~/.qito.psm"
tmp = input("enter secret file path[~/.qito.psm]: ")
if tmp != "":
	filepath = tmp

if os.system("> " + filepath):
	print("error creating file")
	sys.exit()

go_build_command = (
    "go build -ldflags \"-X 'github.com/sparsh2/qito/common.SecretKey="
    + password
    + "' -X 'github.com/sparsh2/qito/common.SecretFilepath="
		+ filepath
		+ "'\""
    + " -o "
    + target
)

if os.system(go_build_command):
	print("error building")
	sys.exit()
