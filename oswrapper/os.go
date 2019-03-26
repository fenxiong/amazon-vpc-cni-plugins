// Copyright 2019 Amazon.com, Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

package oswrapper

import (
	"os/user"

	"golang.org/x/sys/unix"
)

// OS wraps all os related methods need to be used in testing.
type OS interface {
	User
	Unix
}

// User wraps the methods of os/user package to be used in testing.
type User interface {
	Current() (*user.User, error)
	Lookup(username string) (*user.User, error)
}

type userImpl struct {
}

// NewUser creates a new User object
func NewUser() User {
	return &userImpl{}
}

func (*userImpl) Current() (*user.User, error) {
	return user.Current()
}

func (*userImpl) Lookup(username string) (*user.User, error) {
	return user.Lookup(username)
}

// Unix wraps the methods of sys/unix package to be used in testing.
type Unix interface {
	IoctlSetInt(fd int, req uint, value int) error
}

type unixImpl struct {
}

func NewUnix() Unix {
	return &unixImpl{}
}

func (*unixImpl) IoctlSetInt(fd int, req uint, value int) error {
	return unix.IoctlSetInt(fd, req, value)
}