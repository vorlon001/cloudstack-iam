package main

import (
        . "github.com/vorlon001/go-libnss/structs"
)

// Test database objects.
var dbtest_passwd []Passwd
var dbtest_group []Group
var dbtest_shadow []Shadow

func init() {
	// Populates the passwd test db.
	dbtest_passwd = append(dbtest_passwd,
		Passwd{
			Username: "testguy1",
			Password: "x",
			UID:      1500,
			GID:      1500,
			Gecos:    "Test user 1",
			Dir:      "/home/testguy1",
			Shell:    "/bin/bash",
		},
		Passwd{
			Username: "testguy2",
			Password: "x",
			UID:      1501,
			GID:      1501,
			Gecos:    "Test user 2",
			Dir:      "/home/testguy2",
			Shell:    "/bin/bash",
		},
                Passwd{
                        Username: "testguy3",
                        Password: "x",
                        UID:      1502,
                        GID:      1502,
                        Gecos:    "Test user 3",
                        Dir:      "/home/testguy3",
                        Shell:    "/bin/bash",
                },
                Passwd{
                        Username: "testguy4",
                        Password: "x",
                        UID:      1503,
                        GID:      1503,
                        Gecos:    "Test user 4",
                        Dir:      "/home/testguy4",
                        Shell:    "/bin/bash",
                },
                Passwd{
                        Username: "testguy5",
                        Password: "x",
                        UID:      1504,
                        GID:      1504,
                        Gecos:    "Test user 5",
                        Dir:      "/home/testguy5",
                        Shell:    "/bin/bash",
                },
	)

	// Populates the group test db.
	dbtest_group = append(dbtest_group,
		Group{
			Groupname: "testguy1",
			Password:  "x",
			GID:       1500,
			Members:   []string{"testguy1"},
		},
		Group{
			Groupname: "testguy2",
			Password:  "x",
			GID:       1501,
			Members:   []string{"testguy2","sudo"},
		},
		Group{
			Groupname: "testguyz",
			Password:  "x",
			GID:       1499,
			Members:   []string{"testguy1", "testguy2"},
		},
                Group{
                        Groupname: "testguy3",
                        Password:  "x",
                        GID:       1502,
                        Members:   []string{"testguy3","sudo"},
                },
                Group{
                        Groupname: "testguy4",
                        Password:  "x",
                        GID:       1503,
                        Members:   []string{"testguy4","sudo"},
                },
                Group{
                        Groupname: "testguy5",
                        Password:  "x",
                        GID:       1504,
                        Members:   []string{"testguy5"},
                },
                Group{
                        Groupname: "sudo",
                        Password:  "x",
                        GID:       27,
                        Members:   []string{"testguy2", "testguy3","testguy4"},
                },
	)

	// Populates the shadow test db.
	dbtest_shadow = append(dbtest_shadow,
		Shadow{
			Username:        "testguy1",
			Password:        "$6$yZcX.DOY$7bgsJhILMYl3DfMZsYUwoObbVt5Sj9FuujuhVn05Vg9hk.2AXLNy6o1DcPNq0SIyaRZ5YBZer2rYaycuh3qtg1", // Password is "password"
			LastChange:      17920,
			MinChange:       0,
			MaxChange:       99999,
			PasswordWarn:    7,
			InactiveLockout: -1,
			ExpirationDate:  -1,
			Reserved:        -1,
		},
		Shadow{
			Username:        "testguy2",
			Password:        "$6$yZcX.DOY$7bgsJhILMYl3DfMZsYUwoObbVt5Sj9FuujuhVn05Vg9hk.2AXLNy6o1DcPNq0SIyaRZ5YBZer2rYaycuh3qtg1", // Password is "password"
			LastChange:      17920,
			MinChange:       0,
			MaxChange:       99999,
			PasswordWarn:    7,
			InactiveLockout: 0,
			ExpirationDate:  0,
			Reserved:        -1,
		},
                Shadow{
                        Username:        "testguy3",
                        Password:        "*",
                        LastChange:      17920,
                        MinChange:       0,
                        MaxChange:       99999,
                        PasswordWarn:    7,
                        InactiveLockout: 0,
                        ExpirationDate:  0,
                        Reserved:        -1,
                },
	)
}
