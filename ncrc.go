package nchash

import "hash/crc32"

type NginxCrc struct {
	crc uint32
}

func NewNiginxCrc() *NginxCrc {
	return &NginxCrc{crc: 0xffffffff}
}

func (n *NginxCrc) Size() int { return 4 }

func (n *NginxCrc) BlockSize() int { return 1 }

func (n *NginxCrc) Reset() { n.crc = 0xffffffff }

func (n *NginxCrc) Write(p []byte) (length int, err error) {
	for _, v := range p {
		n.crc = crc32.IEEETable[byte(n.crc)^v] ^ (n.crc >> 8)
	}
	return len(p), nil
}

func (n *NginxCrc) Sum32() uint32 {
	return ^n.crc
}

func (n *NginxCrc) Sum(in []byte) []byte {
	s := n.Sum32()
	return append(in, byte(s>>24), byte(s>>16), byte(s>>8), byte(s))
}

func Update(crc uint32, p []byte) uint32 {
	for _, v := range p {
		crc = crc32.IEEETable[byte(crc)^v] ^ (crc >> 8)
	}
	return ^crc
}
