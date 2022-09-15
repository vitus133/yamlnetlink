// Package main is generated from a YAML netlink specification for family "nlctrl".
//
// Description: Generic netlink control protocol. Interface to query information about generic netlink families registered in the kernel - their names, ids, accepted messages and attributes.
//
// Code generated by yamlnetlink-go. DO NOT EDIT.
package main

import (
	"errors"

	"github.com/mdlayher/genetlink"
	"github.com/mdlayher/netlink"
	"golang.org/x/sys/unix"
)

// A Conn is a connection to netlink family "nlctrl".
type Conn struct {
	c *genetlink.Conn
	f genetlink.Family
}

// Dial opens a Conn for netlink family "nlctrl". Any options are passed directly
// to the underlying netlink package.
func Dial(cfg *netlink.Config) (*Conn, error) {
	c, err := genetlink.Dial(cfg)
	if err != nil {
		return nil, err
	}

	f, err := c.GetFamily("nlctrl")
	if err != nil {
		return nil, err
	}

	return &Conn{c: c, f: f}, nil
}

// Close closes the Conn's underlying netlink connection.
func (c *Conn) Close() error { return c.c.Close() }

// DoGetfamilyRequest is used with the DoGetfamily method.
type DoGetfamilyRequest struct {
	// Numerical identifier of the family.
	FamilyId uint16
	// String identifier of the family. Guaranteed to be unique.
	FamilyName string
}

// DoGetfamilyReply is used with the DoGetfamily method.
type DoGetfamilyReply struct {
	// Numerical identifier of the family.
	FamilyId uint16
	// String identifier of the family. Guaranteed to be unique.
	FamilyName string
	Version    uint32
	Hdrsize    uint32
	Maxattr    uint32
	// TODO: field "Ops", type "array-nest"
	// TODO: field "McastGroups", type "array-nest"
}

// DoGetfamily wraps the "getfamily" operation:
// Get information about genetlink family.
func (c *Conn) DoGetfamily(req DoGetfamilyRequest) (*DoGetfamilyReply, error) {
	ae := netlink.NewAttributeEncoder()
	if req.FamilyId != 0 {
		ae.Uint16(unix.CTRL_ATTR_FAMILY_ID, req.FamilyId)
	}
	if req.FamilyName != "" {
		ae.String(unix.CTRL_ATTR_FAMILY_NAME, req.FamilyName)
	}

	b, err := ae.Encode()
	if err != nil {
		return nil, err
	}

	msg := genetlink.Message{
		Header: genetlink.Header{
			Command: unix.CTRL_CMD_GETFAMILY,
			Version: c.f.Version,
		},
		Data: b,
	}

	msgs, err := c.c.Execute(msg, c.f.ID, netlink.Request)
	if err != nil {
		return nil, err
	}

	replies := make([]*DoGetfamilyReply, 0, len(msgs))
	for _, m := range msgs {
		ad, err := netlink.NewAttributeDecoder(m.Data)
		if err != nil {
			return nil, err
		}

		var reply DoGetfamilyReply
		for ad.Next() {
			switch ad.Type() {
			case unix.CTRL_ATTR_FAMILY_ID:
				reply.FamilyId = ad.Uint16()
			case unix.CTRL_ATTR_FAMILY_NAME:
				reply.FamilyName = ad.String()
			case unix.CTRL_ATTR_VERSION:
				reply.Version = ad.Uint32()
			case unix.CTRL_ATTR_HDRSIZE:
				reply.Hdrsize = ad.Uint32()
			case unix.CTRL_ATTR_MAXATTR:
				reply.Maxattr = ad.Uint32()
			case unix.CTRL_ATTR_OPS:
				// TODO: field "reply.Ops", type "array-nest"
			case unix.CTRL_ATTR_MCAST_GROUPS:
				// TODO: field "reply.McastGroups", type "array-nest"
			}
		}

		if err := ad.Err(); err != nil {
			return nil, err
		}

		replies = append(replies, &reply)
	}

	if len(replies) != 1 {
		return nil, errors.New("nlctrl: expected exactly one DoGetfamilyReply")
	}

	return replies[0], nil
}

// DumpGetfamilyReply is used with the DumpGetfamily method.
type DumpGetfamilyReply struct {
	// Numerical identifier of the family.
	FamilyId uint16
	// String identifier of the family. Guaranteed to be unique.
	FamilyName string
	Version    uint32
	Hdrsize    uint32
	Maxattr    uint32
	// TODO: field "Ops", type "array-nest"
	// TODO: field "McastGroups", type "array-nest"
}

// DumpGetfamily wraps the "getfamily" operation:
// Get information about genetlink family.
func (c *Conn) DumpGetfamily() ([]*DumpGetfamilyReply, error) {
	// No attribute arguments.
	var b []byte

	msg := genetlink.Message{
		Header: genetlink.Header{
			Command: unix.CTRL_CMD_GETFAMILY,
			Version: c.f.Version,
		},
		Data: b,
	}

	msgs, err := c.c.Execute(msg, c.f.ID, netlink.Request|netlink.Dump)
	if err != nil {
		return nil, err
	}

	replies := make([]*DumpGetfamilyReply, 0, len(msgs))
	for _, m := range msgs {
		ad, err := netlink.NewAttributeDecoder(m.Data)
		if err != nil {
			return nil, err
		}

		var reply DumpGetfamilyReply
		for ad.Next() {
			switch ad.Type() {
			case unix.CTRL_ATTR_FAMILY_ID:
				reply.FamilyId = ad.Uint16()
			case unix.CTRL_ATTR_FAMILY_NAME:
				reply.FamilyName = ad.String()
			case unix.CTRL_ATTR_VERSION:
				reply.Version = ad.Uint32()
			case unix.CTRL_ATTR_HDRSIZE:
				reply.Hdrsize = ad.Uint32()
			case unix.CTRL_ATTR_MAXATTR:
				reply.Maxattr = ad.Uint32()
			case unix.CTRL_ATTR_OPS:
				// TODO: field "reply.Ops", type "array-nest"
			case unix.CTRL_ATTR_MCAST_GROUPS:
				// TODO: field "reply.McastGroups", type "array-nest"
			}
		}

		if err := ad.Err(); err != nil {
			return nil, err
		}

		replies = append(replies, &reply)
	}

	return replies, nil
}

// DumpGetpolicyRequest is used with the DumpGetpolicy method.
type DumpGetpolicyRequest struct {
	// Numerical identifier of the family.
	FamilyId uint16
	// String identifier of the family. Guaranteed to be unique.
	FamilyName string
	Op         uint32
}

// DumpGetpolicyReply is used with the DumpGetpolicy method.
type DumpGetpolicyReply struct {
	// Numerical identifier of the family.
	FamilyId uint16
	// TODO: field "OpPolicy", type "nest-type-value"
	// TODO: field "Policy", type "nest-type-value"
}

// DumpGetpolicy wraps the "getpolicy" operation:
// Get attribute policy for a genetlink family.
func (c *Conn) DumpGetpolicy(req DumpGetpolicyRequest) ([]*DumpGetpolicyReply, error) {
	ae := netlink.NewAttributeEncoder()
	if req.FamilyId != 0 {
		ae.Uint16(unix.CTRL_ATTR_FAMILY_ID, req.FamilyId)
	}
	if req.FamilyName != "" {
		ae.String(unix.CTRL_ATTR_FAMILY_NAME, req.FamilyName)
	}
	if req.Op != 0 {
		ae.Uint32(unix.CTRL_ATTR_OP, req.Op)
	}

	b, err := ae.Encode()
	if err != nil {
		return nil, err
	}

	msg := genetlink.Message{
		Header: genetlink.Header{
			Command: unix.CTRL_CMD_GETPOLICY,
			Version: c.f.Version,
		},
		Data: b,
	}

	msgs, err := c.c.Execute(msg, c.f.ID, netlink.Request|netlink.Dump)
	if err != nil {
		return nil, err
	}

	replies := make([]*DumpGetpolicyReply, 0, len(msgs))
	for _, m := range msgs {
		ad, err := netlink.NewAttributeDecoder(m.Data)
		if err != nil {
			return nil, err
		}

		var reply DumpGetpolicyReply
		for ad.Next() {
			switch ad.Type() {
			case unix.CTRL_ATTR_FAMILY_ID:
				reply.FamilyId = ad.Uint16()
			case unix.CTRL_ATTR_OP_POLICY:
				// TODO: field "reply.OpPolicy", type "nest-type-value"
			case unix.CTRL_ATTR_POLICY:
				// TODO: field "reply.Policy", type "nest-type-value"
			}
		}

		if err := ad.Err(); err != nil {
			return nil, err
		}

		replies = append(replies, &reply)
	}

	return replies, nil
}
