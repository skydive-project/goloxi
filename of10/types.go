package of10

import (
	"github.com/skydive-project/goloxi"
)

// TODO: set real types
type uint128 = goloxi.Uint128
type Checksum128 [16]byte
type Bitmap128 uint128
type Bitmap512 struct {
	a, b, c, d uint128
}
type Unimplemented struct{}
type BSNVport uint16
type ControllerURI uint16

func (h *Header) MessageType() uint8 {
	return h.Type
}

func (self *Checksum128) Decode(decoder *goloxi.Decoder) error {
	return nil
}

func (self *Checksum128) Serialize(encoder *goloxi.Encoder) error {
	return nil
}
func (self *Bitmap128) Decode(decoder *goloxi.Decoder) error {
	return nil
}

func (self *Bitmap128) Serialize(encoder *goloxi.Encoder) error {
	return nil
}
func (self *Bitmap512) Decode(decoder *goloxi.Decoder) error {
	return nil
}

func (self *Bitmap512) Serialize(encoder *goloxi.Encoder) error {
	return nil
}
func (self *BSNVport) Decode(decoder *goloxi.Decoder) error {
	return nil
}

func (self *BSNVport) Serialize(encoder *goloxi.Encoder) error {
	return nil
}
func (self *ControllerURI) Decode(decoder *goloxi.Decoder) error {
	return nil
}

func (self *ControllerURI) Serialize(encoder *goloxi.Encoder) error {
	return nil
}

type FmCmd uint16

func (self *FmCmd) Serialize(encoder *goloxi.Encoder) error {
	encoder.PutUint16(uint16(*self))
	return nil
}

func (self *FmCmd) Decode(decoder *goloxi.Decoder) error {
	*self = FmCmd(decoder.ReadUint16())
	return nil
}

type MatchBmap uint32

func (self *MatchBmap) Serialize(encoder *goloxi.Encoder) error {
	encoder.PutUint32(uint32(*self))
	return nil
}

func (self *MatchBmap) Decode(decoder *goloxi.Decoder) error {
	*self = MatchBmap(decoder.ReadUint32())
	return nil
}

type WcBmap uint32

func (self *WcBmap) Serialize(encoder *goloxi.Encoder) error {
	encoder.PutUint32(uint32(*self))
	return nil
}

func (self *WcBmap) Decode(decoder *goloxi.Decoder) error {
	*self = WcBmap(decoder.ReadUint32())
	return nil
}

type Match = MatchV1
type PortNo uint16

func (self *PortNo) Serialize(encoder *goloxi.Encoder) error {
	encoder.PutUint16(uint16(*self))
	return nil
}

func (self *PortNo) Decode(decoder *goloxi.Decoder) error {
	*self = PortNo(decoder.ReadUint16())
	return nil
}

func DecodeMessage(data []byte) (goloxi.Message, error) {
	header, err := decodeHeader(goloxi.NewDecoder(data))
	if err != nil {
		return nil, err
	}

	return header.(goloxi.Message), nil
}
