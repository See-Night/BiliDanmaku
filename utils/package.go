package utils

type Pkg struct {
	packageLength uint32
	headerLength  uint16
	version       uint16
	operation     uint32
	sequence      uint32
	body          []byte
}

func (pkg *Pkg) Init(operation uint32, msg WsMsg) {
	pkg.operation = operation
	pkg.body = msg.toBytes()
}

func (pkg *Pkg) ToBytes() []byte {
	pkg.headerLength = 16
	pkg.version = 1
	pkg.sequence = 1

	pkg.packageLength = uint32(len(pkg.body) + 16)

	var msg []byte
	msg = append(msg, Uint32ToBytes(pkg.packageLength)...)
	msg = append(msg, Uint16ToBytes(pkg.headerLength)...)
	msg = append(msg, Uint16ToBytes(pkg.version)...)
	msg = append(msg, Uint32ToBytes(pkg.operation)...)
	msg = append(msg, Uint32ToBytes(pkg.sequence)...)
	msg = append(msg, pkg.body...)

	return msg
}
