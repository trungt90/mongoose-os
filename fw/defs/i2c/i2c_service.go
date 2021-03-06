// Code generated by clubbygen.
// GENERATED FILE DO NOT EDIT
// +build !clubby_strict

package i2c

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"

	"cesanta.com/common/go/mgrpc"
	"cesanta.com/common/go/mgrpc/frame"
	"cesanta.com/common/go/ourjson"
	"cesanta.com/common/go/ourtrace"
	"github.com/cesanta/errors"
	"golang.org/x/net/trace"
)

var _ = bytes.MinRead
var _ = fmt.Errorf
var emptyMessage = ourjson.RawMessage{}
var _ = ourtrace.New
var _ = trace.New

const ServiceID = "http://mongoose-iot.com/fwI2C"

type ReadArgs struct {
	Addr *int64 `json:"addr,omitempty"`
	Len  *int64 `json:"len,omitempty"`
}

type ReadResult struct {
	Data_hex *string `json:"data_hex,omitempty"`
}

type ReadRegBArgs struct {
	Addr *int64 `json:"addr,omitempty"`
	Reg  *int64 `json:"reg,omitempty"`
}

type ReadRegBResult struct {
	Value *int64 `json:"value,omitempty"`
}

type ReadRegWArgs struct {
	Addr *int64 `json:"addr,omitempty"`
	Reg  *int64 `json:"reg,omitempty"`
}

type ReadRegWResult struct {
	Value *int64 `json:"value,omitempty"`
}

type WriteArgs struct {
	Addr     *int64  `json:"addr,omitempty"`
	Data_hex *string `json:"data_hex,omitempty"`
}

type WriteRegBArgs struct {
	Addr  *int64 `json:"addr,omitempty"`
	Reg   *int64 `json:"reg,omitempty"`
	Value *int64 `json:"value,omitempty"`
}

type Service interface {
	Read(ctx context.Context, args *ReadArgs) (*ReadResult, error)
	ReadRegB(ctx context.Context, args *ReadRegBArgs) (*ReadRegBResult, error)
	ReadRegW(ctx context.Context, args *ReadRegWArgs) (*ReadRegWResult, error)
	Scan(ctx context.Context) ([]int64, error)
	Write(ctx context.Context, args *WriteArgs) error
	WriteRegB(ctx context.Context, args *WriteRegBArgs) error
}

type Instance interface {
	Call(context.Context, string, *frame.Command) (*frame.Response, error)
}

func NewClient(i Instance, addr string) Service {
	return &_Client{i: i, addr: addr}
}

type _Client struct {
	i    Instance
	addr string
}

func (c *_Client) Read(ctx context.Context, args *ReadArgs) (res *ReadResult, err error) {
	cmd := &frame.Command{
		Cmd: "I2C.Read",
	}

	cmd.Args = ourjson.DelayMarshaling(args)
	resp, err := c.i.Call(ctx, c.addr, cmd)
	if err != nil {
		return nil, errors.Trace(err)
	}
	if resp.Status != 0 {
		return nil, errors.Trace(&mgrpc.ErrorResponse{Status: resp.Status, Msg: resp.StatusMsg})
	}

	var r *ReadResult
	err = resp.Response.UnmarshalInto(&r)
	if err != nil {
		return nil, errors.Annotatef(err, "unmarshaling response")
	}
	return r, nil
}

func (c *_Client) ReadRegB(ctx context.Context, args *ReadRegBArgs) (res *ReadRegBResult, err error) {
	cmd := &frame.Command{
		Cmd: "I2C.ReadRegB",
	}

	cmd.Args = ourjson.DelayMarshaling(args)
	resp, err := c.i.Call(ctx, c.addr, cmd)
	if err != nil {
		return nil, errors.Trace(err)
	}
	if resp.Status != 0 {
		return nil, errors.Trace(&mgrpc.ErrorResponse{Status: resp.Status, Msg: resp.StatusMsg})
	}

	var r *ReadRegBResult
	err = resp.Response.UnmarshalInto(&r)
	if err != nil {
		return nil, errors.Annotatef(err, "unmarshaling response")
	}
	return r, nil
}

func (c *_Client) ReadRegW(ctx context.Context, args *ReadRegWArgs) (res *ReadRegWResult, err error) {
	cmd := &frame.Command{
		Cmd: "I2C.ReadRegW",
	}

	cmd.Args = ourjson.DelayMarshaling(args)
	resp, err := c.i.Call(ctx, c.addr, cmd)
	if err != nil {
		return nil, errors.Trace(err)
	}
	if resp.Status != 0 {
		return nil, errors.Trace(&mgrpc.ErrorResponse{Status: resp.Status, Msg: resp.StatusMsg})
	}

	var r *ReadRegWResult
	err = resp.Response.UnmarshalInto(&r)
	if err != nil {
		return nil, errors.Annotatef(err, "unmarshaling response")
	}
	return r, nil
}

func (c *_Client) Scan(ctx context.Context) (res []int64, err error) {
	cmd := &frame.Command{
		Cmd: "I2C.Scan",
	}
	resp, err := c.i.Call(ctx, c.addr, cmd)
	if err != nil {
		return nil, errors.Trace(err)
	}
	if resp.Status != 0 {
		return nil, errors.Trace(&mgrpc.ErrorResponse{Status: resp.Status, Msg: resp.StatusMsg})
	}

	var r []int64
	err = resp.Response.UnmarshalInto(&r)
	if err != nil {
		return nil, errors.Annotatef(err, "unmarshaling response")
	}
	return r, nil
}

func (c *_Client) Write(ctx context.Context, args *WriteArgs) (err error) {
	cmd := &frame.Command{
		Cmd: "I2C.Write",
	}

	cmd.Args = ourjson.DelayMarshaling(args)
	resp, err := c.i.Call(ctx, c.addr, cmd)
	if err != nil {
		return errors.Trace(err)
	}
	if resp.Status != 0 {
		return errors.Trace(&mgrpc.ErrorResponse{Status: resp.Status, Msg: resp.StatusMsg})
	}
	return nil
}

func (c *_Client) WriteRegB(ctx context.Context, args *WriteRegBArgs) (err error) {
	cmd := &frame.Command{
		Cmd: "I2C.WriteRegB",
	}

	cmd.Args = ourjson.DelayMarshaling(args)
	resp, err := c.i.Call(ctx, c.addr, cmd)
	if err != nil {
		return errors.Trace(err)
	}
	if resp.Status != 0 {
		return errors.Trace(&mgrpc.ErrorResponse{Status: resp.Status, Msg: resp.StatusMsg})
	}
	return nil
}

//func RegisterService(i *clubby.Instance, impl Service) error {
//s := &_Server{impl}
//i.RegisterCommandHandler("I2C.Read", s.Read)
//i.RegisterCommandHandler("I2C.ReadRegB", s.ReadRegB)
//i.RegisterCommandHandler("I2C.ReadRegW", s.ReadRegW)
//i.RegisterCommandHandler("I2C.Scan", s.Scan)
//i.RegisterCommandHandler("I2C.Write", s.Write)
//i.RegisterCommandHandler("I2C.WriteRegB", s.WriteRegB)
//i.RegisterService(ServiceID, _ServiceDefinition)
//return nil
//}

type _Server struct {
	impl Service
}

func (s *_Server) Read(ctx context.Context, src string, cmd *frame.Command) (interface{}, error) {
	var args ReadArgs
	if len(cmd.Args) > 0 {
		if err := cmd.Args.UnmarshalInto(&args); err != nil {
			return nil, errors.Annotatef(err, "unmarshaling args")
		}
	}
	return s.impl.Read(ctx, &args)
}

func (s *_Server) ReadRegB(ctx context.Context, src string, cmd *frame.Command) (interface{}, error) {
	var args ReadRegBArgs
	if len(cmd.Args) > 0 {
		if err := cmd.Args.UnmarshalInto(&args); err != nil {
			return nil, errors.Annotatef(err, "unmarshaling args")
		}
	}
	return s.impl.ReadRegB(ctx, &args)
}

func (s *_Server) ReadRegW(ctx context.Context, src string, cmd *frame.Command) (interface{}, error) {
	var args ReadRegWArgs
	if len(cmd.Args) > 0 {
		if err := cmd.Args.UnmarshalInto(&args); err != nil {
			return nil, errors.Annotatef(err, "unmarshaling args")
		}
	}
	return s.impl.ReadRegW(ctx, &args)
}

func (s *_Server) Scan(ctx context.Context, src string, cmd *frame.Command) (interface{}, error) {
	return s.impl.Scan(ctx)
}

func (s *_Server) Write(ctx context.Context, src string, cmd *frame.Command) (interface{}, error) {
	var args WriteArgs
	if len(cmd.Args) > 0 {
		if err := cmd.Args.UnmarshalInto(&args); err != nil {
			return nil, errors.Annotatef(err, "unmarshaling args")
		}
	}
	return nil, s.impl.Write(ctx, &args)
}

func (s *_Server) WriteRegB(ctx context.Context, src string, cmd *frame.Command) (interface{}, error) {
	var args WriteRegBArgs
	if len(cmd.Args) > 0 {
		if err := cmd.Args.UnmarshalInto(&args); err != nil {
			return nil, errors.Annotatef(err, "unmarshaling args")
		}
	}
	return nil, s.impl.WriteRegB(ctx, &args)
}

var _ServiceDefinition = json.RawMessage([]byte(`{
  "methods": {
    "Read": {
      "args": {
        "addr": {
          "doc": "Address of the device, 7 or 10 bits (not including the r/w bit).",
          "type": "integer"
        },
        "len": {
          "doc": "Number of bytes to read.",
          "type": "integer"
        }
      },
      "doc": "Read specified number of bytes from the specified device.",
      "result": {
        "properties": {
          "data_hex": {
            "doc": "Hex-encoded data.",
            "type": "string"
          }
        },
        "type": "object"
      }
    },
    "ReadRegB": {
      "args": {
        "addr": {
          "doc": "Address of the device, 7 or 10 bits (not including the r/w bit).",
          "type": "integer"
        },
        "reg": {
          "doc": "Register number.",
          "type": "integer"
        }
      },
      "doc": "Read value of a byte-sized (8-bit) register.",
      "result": {
        "properties": {
          "value": {
            "doc": "Register value read from the device.",
            "type": "integer"
          }
        },
        "type": "object"
      }
    },
    "ReadRegW": {
      "args": {
        "addr": {
          "doc": "Address of the device, 7 or 10 bits (not including the r/w bit).",
          "type": "integer"
        },
        "reg": {
          "doc": "Register number.",
          "type": "integer"
        }
      },
      "doc": "Read value of a word-sized (16-bit) register.",
      "result": {
        "properties": {
          "value": {
            "doc": "Register value read from the device.",
            "type": "integer"
          }
        },
        "type": "object"
      }
    },
    "Scan": {
      "doc": "Scan the I2C bus, returning addresses of devices that responded.",
      "result": {
        "items": {
          "doc": "List of device addresses for which an ACK was received.",
          "type": "integer"
        },
        "type": "array"
      }
    },
    "Write": {
      "args": {
        "addr": {
          "doc": "Address of the device, 7 or 10 bits (not including the r/w bit).",
          "type": "integer"
        },
        "data_hex": {
          "doc": "Hex-encoded data to write.",
          "type": "string"
        }
      },
      "doc": "Write the specified data to the device with the specified address."
    },
    "WriteRegB": {
      "args": {
        "addr": {
          "doc": "Address of the device, 7 or 10 bits (not including the r/w bit).",
          "type": "integer"
        },
        "reg": {
          "doc": "Register number.",
          "type": "integer"
        },
        "value": {
          "doc": "Register value to write.",
          "type": "integer"
        }
      },
      "doc": "Write value of a word-sized (16-bit) register."
    }
  },
  "name": "I2C",
  "namespace": "http://mongoose-iot.com/fw"
}`))
