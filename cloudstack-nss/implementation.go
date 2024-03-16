package main

import (
	. "github.com/vorlon001/go-libnss"
	. "github.com/vorlon001/go-libnss/structs"
        "log/syslog"
	"fmt"
)


func pamLog(format string, args ...interface{}) {
        l, err := syslog.New(syslog.LOG_AUTH|syslog.LOG_WARNING, "pam-cloudstack")
        if err != nil {
                return
        }
        l.Warning(fmt.Sprintf(format, args...))
}


// Placeholder main() stub is neccessary for compile.
func main() {}

func init(){
	// We set our implementation to "CloudStack", so that go-libnss will use the methods we create
	SetImpl(CloudStack{})
}

// We're creating a struct that implements LIBNSS stub methods.
type CloudStack struct { LIBNSS }

////////////////////////////////////////////////////////////////
// Passwd Methods
////////////////////////////////////////////////////////////////

// PasswdAll() will populate all entries for libnss
func (self CloudStack) PasswdAll() (Status, []Passwd) {
	if len(dbtest_passwd) == 0 {
		return StatusUnavail, []Passwd{}
	}
	pamLog("PasswdAll:%v\n",dbtest_passwd)
	return StatusSuccess, dbtest_passwd
}

// PasswdByName() returns a single entry by name.
func (self CloudStack) PasswdByName(name string) (Status, Passwd) {
	for _, entry := range dbtest_passwd {
		if entry.Username == name {
			pamLog("PasswdByName name:%v, %v\n", name, entry)
			return StatusSuccess, entry
		}
	}
	return StatusNotfound, Passwd{}
}

// PasswdByUid() returns a single entry by uid.
func (self CloudStack) PasswdByUid(uid uint) (Status, Passwd) {
	for _, entry := range dbtest_passwd {
		if entry.UID == uid {
			pamLog("PasswdByUid uid:%v, %v\n",uid, entry)
			return StatusSuccess, entry
		}
	}
	return StatusNotfound, Passwd{}
}


////////////////////////////////////////////////////////////////
// Group Methods
////////////////////////////////////////////////////////////////
// endgrent
func (self CloudStack) GroupAll() (Status, []Group) {
	if len(dbtest_group) == 0 {
		return StatusUnavail, []Group{}
	}
	pamLog("GroupAll %v\n", dbtest_group)
	return StatusSuccess, dbtest_group
}

// getgrent
func (self CloudStack) GroupByName(name string) (Status, Group) {
	for _, entry := range dbtest_group {
		if entry.Groupname == name {
			pamLog("GroupByName name:%v, %v\n", name, entry)
			return StatusSuccess, entry
		}
	}
	return StatusNotfound, Group{}
}

// getgrnam
func (self CloudStack) GroupByGid(gid uint) (Status, Group) {
	for _, entry := range dbtest_group {
		if entry.GID == gid {
			pamLog("GroupByName gid:%v, %v\n", gid, entry)
			return StatusSuccess, entry
		}
	}
	return StatusNotfound, Group{}
}

////////////////////////////////////////////////////////////////
// Shadow Methods
////////////////////////////////////////////////////////////////
// endspent
func (self CloudStack) ShadowAll() (Status, []Shadow) {
	if len(dbtest_shadow) == 0 {
		return StatusUnavail, []Shadow{}
	}
	pamLog("ShadowAll %v\n", dbtest_shadow)
	return StatusSuccess, dbtest_shadow
}

// getspent
func (self CloudStack) ShadowByName(name string) (Status, Shadow) {
	for _, entry := range dbtest_shadow {
		if entry.Username == name {
			pamLog("ShadowByName name:%v, %v\n", name, entry)
			return StatusSuccess, entry
		}
	}
	return StatusNotfound, Shadow{}
}
