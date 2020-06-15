//Package nodef comment
// This file war generated by tars2go 1.1
// Generated from NodeF.tars
package nodef

import (
	"fmt"
	"github.com/wangstrider/TarsGo/tars/protocol/codec"
)

//ServerInfo strcut implement
type ServerInfo struct {
	Application string `json:"application"`
	ServerName  string `json:"serverName"`
	Pid         int32  `json:"pid"`
	Adapter     string `json:"adapter"`
}

func (st *ServerInfo) resetDefault() {
}

//ReadFrom reads  from _is and put into struct.
func (st *ServerInfo) ReadFrom(_is *codec.Reader) error {
	var err error
	var length int32
	var have bool
	var ty byte
	st.resetDefault()

	err = _is.Read_string(&st.Application, 0, true)
	if err != nil {
		return err
	}

	err = _is.Read_string(&st.ServerName, 1, true)
	if err != nil {
		return err
	}

	err = _is.Read_int32(&st.Pid, 2, true)
	if err != nil {
		return err
	}

	err = _is.Read_string(&st.Adapter, 3, false)
	if err != nil {
		return err
	}

	_ = length
	_ = have
	_ = ty
	return nil
}

//ReadBlock reads struct from the given tag , require or optional.
func (st *ServerInfo) ReadBlock(_is *codec.Reader, tag byte, require bool) error {
	var err error
	var have bool
	st.resetDefault()

	err, have = _is.SkipTo(codec.STRUCT_BEGIN, tag, require)
	if err != nil {
		return err
	}
	if !have {
		if require {
			return fmt.Errorf("require ServerInfo, but not exist. tag %d", tag)
		}
		return nil

	}

	st.ReadFrom(_is)

	err = _is.SkipToStructEnd()
	if err != nil {
		return err
	}
	_ = have
	return nil
}

//WriteTo encode struct to buffer
func (st *ServerInfo) WriteTo(_os *codec.Buffer) error {
	var err error

	err = _os.Write_string(st.Application, 0)
	if err != nil {
		return err
	}

	err = _os.Write_string(st.ServerName, 1)
	if err != nil {
		return err
	}

	err = _os.Write_int32(st.Pid, 2)
	if err != nil {
		return err
	}

	err = _os.Write_string(st.Adapter, 3)
	if err != nil {
		return err
	}

	return nil
}

//WriteBlock encode struct
func (st *ServerInfo) WriteBlock(_os *codec.Buffer, tag byte) error {
	var err error
	err = _os.WriteHead(codec.STRUCT_BEGIN, tag)
	if err != nil {
		return err
	}

	st.WriteTo(_os)

	err = _os.WriteHead(codec.STRUCT_END, 0)
	if err != nil {
		return err
	}
	return nil
}
