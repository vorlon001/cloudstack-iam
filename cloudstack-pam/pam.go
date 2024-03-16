// +build darwin linux

/*
Copyright (c) 2017 Uber Technologies, Inc.

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/

package main

import "C"

// code in here can't be tested because it relies on cgo. :(

import (
//	"os"
//	"fmt"
	"unsafe"
)

/*
#cgo LDFLAGS: -lpam -fPIC
#include <security/pam_appl.h>
#include <stdlib.h>

char *string_from_argv(int, char**);
char *get_user(pam_handle_t *pamh);
int get_uid(char *user);
char *get_password(pam_handle_t *pamh);
*/
import "C"

func init() {
	if !disablePtrace() {
		pamLog("unable to disable ptrace")
	}
}

func sliceFromArgv(argc C.int, argv **C.char) []string {
	r := make([]string, 0, argc)
	for i := 0; i < int(argc); i++ {
		s := C.string_from_argv(C.int(i), argv)
		defer C.free(unsafe.Pointer(s))
		r = append(r, C.GoString(s))
	}
	return r
}

func GoFunction(str *C.char) string {
	return C.GoString(str)
}

//export pam_sm_authenticate
func pam_sm_authenticate(pamh *C.pam_handle_t, flags, argc C.int, argv **C.char) C.int {

	cUsername := C.get_user(pamh)
	if cUsername == nil {
		return C.PAM_USER_UNKNOWN
	}

	user := GoFunction(cUsername)

	pamLog("cUsername: %v\n", user)

	defer C.free(unsafe.Pointer(cUsername))

	uid := int(C.get_uid(cUsername))
	if uid < 0 {
		return C.PAM_USER_UNKNOWN
	}
	pamLog("cUsername uid: %v:%v \n", user, uid )

        cPassword := C.get_password(pamh)
        if cPassword == nil {
                return C.PAM_USER_UNKNOWN
        }

        password := GoFunction(cPassword)

        pamLog("cPassword: %v\n", password)

        defer C.free(unsafe.Pointer(cPassword))


        if user == "testguy3" && password == "toor" { //&& password == "password" {
                return C.PAM_SUCCESS
        }

        if user == "testguy4" && password == "toor" { //&& password == "password" {
                return C.PAM_SUCCESS
        }


        if user == "testguy5" && password == "toor" { //&& password == "password" {
                return C.PAM_SUCCESS
        }

	return C.PAM_USER_UNKNOWN;

}

//export pam_sm_setcred
//https://fossies.org/dox/openpam-20190224/pam__sm__setcred_8c.html
func pam_sm_setcred(pamh *C.pam_handle_t, flags, argc C.int, argv **C.char) C.int {
	pamLog("pam_sm_setcred\n")
	pamLog("pamh: %v\n", pamh)
	pamLog("flags: %v\n", flags)
	pamLog("argc: %v\n", argc)
	pamLog("argv: %v\n", *argv)
	return C.PAM_SUCCESS
}

//export pam_sm_acct_mgmt
//https://fossies.org/dox/openpam-20190224/pam__sm__acct__mgmt_8c.html
func pam_sm_acct_mgmt(pamh *C.pam_handle_t, flags, argc C.int, argv **C.char) C.int {
	pamLog("pam_sm_acct_mgmt\n")
        pamLog("pamh: %v\n", pamh)
        pamLog("flags: %v\n", flags)
        pamLog("argc: %v\n", argc)
        pamLog("argv: %v\n", *argv)
	return C.PAM_SUCCESS
}

//export pam_sm_open_session
//https://fossies.org/dox/openpam-20190224/pam__sm__open__session_8c.html
func pam_sm_open_session(pamh *C.pam_handle_t, flags, argc C.int, argv **C.char) C.int {
	pamLog("pam_sm_open_session\n")
        pamLog("pamh: %v\n", pamh)
        pamLog("flags: %v\n", flags)
        pamLog("argc: %v\n", argc)
        pamLog("argv: %v\n", *argv)
	return C.PAM_SUCCESS
}

//export pam_sm_close_session
//https://fossies.org/dox/openpam-20190224/pam__sm__close__session_8c.html
func pam_sm_close_session(pamh *C.pam_handle_t, flags, argc C.int, argv **C.char) C.int {
	pamLog("pam_sm_close_session\n")
        pamLog("pamh: %v\n", pamh)
        pamLog("flags: %v\n", flags)
        pamLog("argc: %v\n", argc)
        pamLog("argv: %v\n", *argv)
	return C.PAM_SUCCESS
}

//export pam_sm_chauthtok
//https://fossies.org/dox/openpam-20190224/pam__get__authtok_8c.html
func pam_sm_chauthtok(pamh *C.pam_handle_t, flags, argc C.int, argv **C.char) C.int {
	pamLog("pam_sm_chauthtok\n")
        pamLog("pamh: %v\n", pamh)
        pamLog("flags: %v\n", flags)
        pamLog("argc: %v\n", argc)
        pamLog("argv: %v\n", *argv)
	return C.PAM_SUCCESS
}
