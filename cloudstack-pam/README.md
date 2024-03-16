# CLOUD.STACK VERSION

### STEP 1
install ubuntu 22.03LTS

### STEP 2
```
apt-get update -y
apt install libpam-dev golang -y
```

### STEP 3
build
```
go build -buildmode=c-shared -o pam_cloudstack.so
cp pam_cloudstack.so /usr/lib/x86_64-linux-gnu/security/
```

### STEP 4
change /etc/pam.d

### STEP 5
create user test w/o password; add user in group sudo

### STEP 6
```
ssh test@<IP>
>>password toor

in VM
su
>>username: test
>>password toor

sudo
>>username: test
>>password toor

```

###
base on https://github.com/uber/pam-ussh
