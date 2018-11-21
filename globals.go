package goloxi

const (
	VERSION_1_0 = 1
	VERSION_1_1 = 2
	VERSION_1_2 = 3
	VERSION_1_3 = 4
	VERSION_1_4 = 5
	VERSION_1_5 = 6
)

type Serializable interface {
	Serialize(encoder *Encoder) error
}

type Deserializable interface {
	Decode(decoder *Decoder) error
}

type Message interface {
	Serializable
	MessageType() uint8
}

type Uint128 struct {
	Hi uint64
	Lo uint64
}
