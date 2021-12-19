package main

import (
	"testing"
)

func TestC200B40A82(t *testing.T) {
	runExpressionTest("C200B40A82", 3, t)
}

func Test04005AC33890(t *testing.T) {
	runExpressionTest("04005AC33890", 54, t)
}

func Test880086C3E88112(t *testing.T) {
	runExpressionTest("880086C3E88112", 7, t)
}

func TestCE00C43D881120(t *testing.T) {
	runExpressionTest("CE00C43D881120", 9, t)
}

func TestD8005AC2A8F0(t *testing.T) {
	runExpressionTest("D8005AC2A8F0", 1, t)
}

func TestF600BC2D8F(t *testing.T) {
	runExpressionTest("F600BC2D8F", 0, t)
}

func Test9C005AC2F8F0(t *testing.T) {
	runExpressionTest("9C005AC2F8F0", 0, t)
}

func Test9C0141080250320F1802104A08(t *testing.T) {
	runExpressionTest("9C0141080250320F1802104A08", 1, t)
}

func runExpressionTest(transmission string, expected int64, t *testing.T) {
	packet, err := DecodePacket(transmission)
	if err != nil {
		t.Fatal(err)
	}
	actual, err := packet.Value()
	if err != nil {
		t.Fatal(err)
	}
	if actual != expected {
		t.Errorf("expected: %d, actual: %d", expected, actual)
	}
}
