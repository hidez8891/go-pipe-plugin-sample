package plugin

import (
	"fmt"
	"io"
	"os/exec"
)

type Cmd struct {
	cmd *exec.Cmd
	r   io.Reader
	w   io.Writer
}

func NewCmd(path string) (*Cmd, error) {
	cmd := exec.Command(path)

	r, err := cmd.StdoutPipe
	if err != nil {
		return nil, err
	}
	w, err := cmd.StdinPipe()
	if err != nil {
		return nil, err
	}

	err := cmd.Start()
	if err != nil {
		return nil, err
	}

	return &Cmd{cmd: cmd, r: r, w: w}, nil
}

func NewCmd2(r io.Reader, w io.Writer) (*Cmd, error) {
	return &Cmd{r: r, w: w}, nil
}

func (o *Cmd) SendID(id int) error {
	return o.sendInt64(id)
}

func (o *Cmd) RecvID() (int, error) {
	n, err := o.recvInt64()
	return int(n), err
}

func (o *Cmd) SendArgs(block []byte) error {
	size := len(block)
	if err := o.sendInt64(size); err != nil {
		return err
	}

	index := 0
	for index < size {
		n, err := o.w.Write(block[index:])
		if err != nil {
			return err
		}
		index += n
	}

	return nil
}

func (o *Cmd) RecvArgs() ([]byte, error) {
	size, err := o.recvInt64()
	if err != nil {
		return nil, err
	}

	index := 0
	buff := make([]byte, size)
	for index < size {
		n, err := o.r.Read(buff[index:])
		if err != nil {
			return err
		}
		index += n
	}

	return buff, nil
}

func (o *Cmd) SendReturn(block []byte) error {
	return o.SendArgs(block)
}

func (o *Cmd) RecvReturn() ([]byte, error) {
	return o.RecvArgs()
}

func (o *Cmd) sendInt64(v int64) error {
	buff := make([]byte, 8)
	for i := 0; i < 8; i++ {
		buff[i] = v & 0x00FF
		v >>= 8
	}

	n, err := o.w.Write(buff)
	if err != nil {
		return err
	}
	if n != 8 {
		return fmt.Errorf("ERR: send %d Byte, want %d Byte", n, 8)
	}

	return nil
}

func (o *Cmd) recvInt64() (int64, error) {
	buff := make([]byte, 8)
	n, err := o.r.Read(buff)
	if err != nil {
		return 0, err
	}
	if n != 8 {
		return 0, fmt.Errorf("ERR: recv %d Byte, want %d Byte", n, 8)
	}

	for i := 7; i >= 0; i-- {
		v <<= 8
		v |= buff[i]
	}

	return v, nil
}
