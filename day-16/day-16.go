package main

import (
	"errors"
	"fmt"
	"math"
	"regexp"
	"strconv"
)

type Packet interface {
	Version() int
	TypeID() int
	Value() (int64, error)
	VersionSum() int
}

type LiteralPacket struct {
	version int
	typeID  int
	value   int64
}

func (l LiteralPacket) Version() int {
	return l.version
}

func (l LiteralPacket) TypeID() int {
	return l.typeID
}

func (l LiteralPacket) Value() (int64, error) {
	return l.value, nil
}

func (l LiteralPacket) VersionSum() int {
	return l.version
}

type OperatorPacket struct {
	version int
	typeID  int
	packets []Packet
}

func (o OperatorPacket) Version() int {
	return o.version
}

func (o OperatorPacket) TypeID() int {
	return o.typeID
}

func (o OperatorPacket) Value() (int64, error) {
	switch o.typeID {
	case 0:
		sum := int64(0)
		for _, packet := range o.packets {
			value, err := packet.Value()
			if err != nil {
				return 0, err
			}
			sum += value
		}
		return sum, nil
	case 1:
		product := int64(1)
		for _, packet := range o.packets {
			value, err := packet.Value()
			if err != nil {
				return 0, err
			}
			product *= value
		}
		return product, nil
	case 2:
		minimum := int64(math.MaxInt64)
		for _, packet := range o.packets {
			value, err := packet.Value()
			if err != nil {
				return 0, err
			}
			if value < minimum {
				minimum = value
			}
		}
		return minimum, nil
	case 3:
		maximum := int64(math.MinInt64)
		for _, packet := range o.packets {
			value, err := packet.Value()
			if err != nil {
				return 0, err
			}
			if value > maximum {
				maximum = value
			}
		}
		return maximum, nil
	case 5:
		if len(o.packets) != 2 {
			return 0, errors.New(fmt.Sprintf("expected 2 packets but was %d", len(o.packets)))
		}
		firstValue, err := o.packets[0].Value()
		if err != nil {
			return 0, err
		}
		secondValue, err := o.packets[1].Value()
		if err != nil {
			return 0, err
		}
		if firstValue > secondValue {
			return 1, nil
		} else {
			return 0, nil
		}
	case 6:
		if len(o.packets) != 2 {
			return 0, errors.New(fmt.Sprintf("expected 2 packets but was %d", len(o.packets)))
		}
		firstValue, err := o.packets[0].Value()
		if err != nil {
			return 0, err
		}
		secondValue, err := o.packets[1].Value()
		if err != nil {
			return 0, err
		}
		if firstValue < secondValue {
			return 1, nil
		} else {
			return 0, nil
		}
	case 7:
		if len(o.packets) != 2 {
			return 0, errors.New(fmt.Sprintf("expected 2 packets but was %d", len(o.packets)))
		}
		firstValue, err := o.packets[0].Value()
		if err != nil {
			return 0, err
		}
		secondValue, err := o.packets[1].Value()
		if err != nil {
			return 0, err
		}
		if firstValue == secondValue {
			return 1, nil
		} else {
			return 0, nil
		}
	default:
		return 0, errors.New(fmt.Sprintf("unexpected type ID: %d", o.typeID))
	}
}

func (o OperatorPacket) VersionSum() int {
	sum := o.version
	for _, packet := range o.packets {
		sum += packet.VersionSum()
	}
	return sum
}

func DecodePacket(transmission string) (Packet, error) {
	bits, err := toBinary(transmission)
	if err != nil {
		return nil, err
	}
	var packets []Packet
	packets, _, err = parsePackets(bits, 0)
	if err != nil {
		return LiteralPacket{}, err
	}
	if len(packets) > 1 {
		return LiteralPacket{}, errors.New(fmt.Sprintf("expected one packet but found %d", len(packets)))
	}
	return packets[0], nil
}

func parsePackets(bits string, pos int) ([]Packet, int, error) {
	var packets []Packet
	for pos < len(bits) && !isPadding(bits, pos) {
		packet, nextPos, err := parsePacket(bits, pos)
		if err != nil {
			return nil, 0, err
		}
		packets = append(packets, packet)
		pos = nextPos
	}
	return packets, pos, nil
}

func isPadding(bits string, pos int) bool {
	for p := pos; p < len(bits); p++ {
		if bits[p] == '1' {
			return false
		}
	}
	return true
}

func parsePacket(bits string, pos int) (Packet, int, error) {
	version, typeIDPos, err := parseVersion(bits, pos)
	if err != nil {
		return nil, 0, err
	}
	typeID, packetPos, err := parseTypeID(bits, typeIDPos)
	if err != nil {
		return nil, 0, err
	}
	var packet Packet
	var nextPos int
	if typeID == 4 {
		packet, nextPos, err = parseLiteral(version, typeID, bits, packetPos)
		if err != nil {
			return LiteralPacket{}, 0, err
		}
	} else {
		packet, nextPos, err = parseOperator(version, typeID, bits, packetPos)
		if err != nil {
			return OperatorPacket{}, 0, err
		}
	}
	return packet, nextPos, nil
}

func parseBitsAsInt(bits string, pos int, length int) (int, int, error) {
	if length < 1 || length > 32 {
		return 0, 0, errors.New(fmt.Sprintf("length out of range [1-32], was %d", length))
	}
	if pos+length >= len(bits) {
		return 0, 0, errors.New(fmt.Sprintf("slice out of range, was [%d:%d]", pos, pos+length))
	}
	value, err := strconv.ParseInt(bits[pos:pos+length], 2, 64)
	if err != nil {
		return 0, 0, err
	}
	return int(value), pos + length, nil
}

func parseVersion(bits string, pos int) (int, int, error) {
	return parseBitsAsInt(bits, pos, 3)
}

func parseTypeID(bits string, pos int) (int, int, error) {
	return parseBitsAsInt(bits, pos, 3)
}

func parseLiteral(version int, typeID int, bits string, pos int) (LiteralPacket, int, error) {
	a := ""
	currentPos := pos
	for bits[currentPos] == '1' {
		a += bits[currentPos+1 : currentPos+4+1]
		currentPos += 5
	}
	a += bits[currentPos+1 : currentPos+4+1]
	value, err := strconv.ParseInt(a, 2, 64)
	if err != nil {
		return LiteralPacket{}, 0, err
	}
	return LiteralPacket{version, typeID, value}, currentPos + 5, err
}

func parseOperator(version int, typeID int, bits string, pos int) (OperatorPacket, int, error) {
	if bits[pos] == '0' {
		return parseLengthOperator(version, typeID, bits, pos+1)
	} else {
		return parseCountOperator(version, typeID, bits, pos+1)
	}
}

func parseLengthOperator(version int, typeID int, bits string, pos int) (OperatorPacket, int, error) {
	length, err := strconv.ParseInt(bits[pos:pos+15], 2, 32)
	if err != nil {
		return OperatorPacket{}, 0, err
	}
	var packets []Packet
	var offset int
	packets, offset, err = parsePackets(bits[pos+15:pos+15+int(length)], 0)
	if err != nil {
		return OperatorPacket{}, 0, err
	}
	return OperatorPacket{version, typeID, packets}, pos + 15 + offset, nil
}

func parseCountOperator(version int, typeID int, bits string, pos int) (OperatorPacket, int, error) {
	count, err := strconv.ParseInt(bits[pos:pos+11], 2, 32)
	if err != nil {
		return OperatorPacket{}, 0, err
	}
	currentPos := pos + 11
	var packets []Packet
	for i := 0; i < int(count); i++ {
		var packet Packet
		packet, currentPos, err = parsePacket(bits, currentPos)
		if err != nil {
			return OperatorPacket{}, 0, nil
		}
		packets = append(packets, packet)
	}
	return OperatorPacket{version, typeID, packets}, currentPos, nil
}

func toBinary(transmission string) (string, error) {
	match, err := regexp.MatchString("^[0-9A-F]+$", transmission)
	if err != nil {
		return "", err
	}
	if !match {
		err = errors.New("expected input to match \\[^[0-9A-F]+$\\]")
		return "", err
	}
	bits := ""
	for _, c := range transmission {
		value, err := strconv.ParseInt(string(c), 16, 32)
		if err != nil {
			return "", err
		}
		bits += fmt.Sprintf("%04s", strconv.FormatInt(value, 2))
	}
	return bits, nil
}
