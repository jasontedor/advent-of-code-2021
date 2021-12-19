package main

import (
	"testing"
)

func TestD2FE28(t *testing.T) {
	runVersionTest("D2FE28", 6, t)
}

func Test38006F45291200(t *testing.T) {
	runVersionTest("38006F45291200", 9, t)
}

func TestEE00D40C823060(t *testing.T) {
	runVersionTest("EE00D40C823060", 14, t)
}

func Test8A004A801A8002F478(t *testing.T) {
	runVersionTest("8A004A801A8002F478", 16, t)
}

func Test620080001611562C8802118E34(t *testing.T) {
	runVersionTest("620080001611562C8802118E34", 12, t)
}

func TestC0015000016115A2E0802F182340(t *testing.T) {
	runVersionTest("C0015000016115A2E0802F182340", 23, t)
}

func TestA0016C880162017C3686B18A3D4780(t *testing.T) {
	runVersionTest("A0016C880162017C3686B18A3D4780", 31, t)
}

func runVersionTest(transmission string, expected int, t *testing.T) {
	packet, err := DecodePacket(transmission)
	if err != nil {
		t.Fatal(err)
	}
	actual := packet.VersionSum()
	if actual != expected {
		t.Errorf("expected: %d, actual: %d", expected, actual)
	}
}
