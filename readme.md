
# go-acl

An access control layer object that provides a basic interface for managing groups and actions.


## sales pitch

My library offers:

- a struct that tracks actions and groups
- provides simple ACL methods `Grant`, `Revoke`, `Can`, `Not`, and `Clear`

It does not:

- connect to any data sources (directly or abstractly)


## usage

Using my library is very simple:

    import "github.com/cdelorme/go-acl"

The package name is `acl` and you can get a new Acl with:

    anAcl := acl.Acl{}

You can Grant one or more permissions with:

    anAcl.Grant("Admin", "Read", "Write", "Execute", "Water the flowers")

You can revoke one or more permissions with:

    anAcl.Revoke("Admin", "Water the flowers")

You can clear your groups (such as when reloading from a data source) with:

    anAcl.Clear()

Actions and Groups are added on first reference.

You have direct access to Groups (a basic `map[string][]string`) and Actions (a `[]string` array), if you need it.

**Permissions are case-sensative; use of constants is recommended.**

_I chose strings to represent permissions since they are the most flexible form of generic-type storage._

