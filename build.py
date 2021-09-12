import random
import os
import sys

chars = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ!#%&()*+-<=>?@^_"

password = ''.join(random.choice(chars) for i in range(12))

target = "target/psm"

filepath = "~/.pmgr.key"
tmp = input("enter secret file path[~/.pmgr.key]: ")
if tmp != "":
	filepath = tmp

if os.system("> " + filepath):
	print("error creating file")
	sys.exit()

go_build_command = (
    "go build -ldflags \"-X 'github.com/sparsh2/pmgr/common.SecretKey="
    + password
    + "' -X 'github.com/sparsh2/pmgr/common.SecretFilepath="
		+ filepath
		+ "'\""
    + " -o "
    + target
)

if os.system(go_build_command):
	print("error building")
	sys.exit()
