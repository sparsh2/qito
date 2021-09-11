import random
import os

chars = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ!#$%&()*+-<=>?@^_"

password = ''.join(random.choice(chars) for i in range(12))

target = "target/psm"

go_build_command = (
    "go build -ldflags \"-X 'github.com/sparsh2/pmgr/common.SecretKey="
    + password
    + "'\""
    + " -o "
    + target
)

os.system(go_build_command)
