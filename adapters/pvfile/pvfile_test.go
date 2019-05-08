
package pvfile

import (
	"bytes"
	"io"
	"testing"
)

const (
	File10Lines = `"'NLNN3020186F5208","11.70","9.60","0","0.00","0.00","0","233.30","0.00","0.00","0.40","0.00","0.00","0","0","0","50.01","32.40","May.01 00:04,2019","0.05","904.90"
"'NLNN3020186F5208","11.70","9.60","0","0.00","0.00","0","233.30","0.00","0.00","0.40","0.00","0.00","0","0","0","50.01","32.40","May.01 00:09,2019","0.05","904.90"
"'NLNN3020186F5208","11.70","9.60","0","0.00","0.00","0","233.30","0.00","0.00","0.40","0.00","0.00","0","0","0","50.01","32.40","May.01 00:16,2019","0.05","904.90"
"'NLNN3020186F5208","11.70","9.60","0","0.00","0.00","0","233.30","0.00","0.00","0.40","0.00","0.00","0","0","0","50.01","32.40","May.01 00:21,2019","0.05","904.90"
"'NLNN3020186F5208","11.70","9.60","0","0.00","0.00","0","233.30","0.00","0.00","0.40","0.00","0.00","0","0","0","50.01","32.40","May.01 00:28,2019","0.05","904.90"
"'NLNN3020186F5208","11.70","9.60","0","0.00","0.00","0","233.30","0.00","0.00","0.40","0.00","0.00","0","0","0","50.01","32.40","May.01 00:33,2019","0.05","904.90"
"'NLNN3020186F5208","11.70","9.60","0","0.00","0.00","0","233.30","0.00","0.00","0.40","0.00","0.00","0","0","0","50.01","32.40","May.01 00:40,2019","0.05","904.90"
"'NLNN3020186F5208","11.70","9.60","0","0.00","0.00","0","233.30","0.00","0.00","0.40","0.00","0.00","0","0","0","50.01","32.40","May.01 00:45,2019","0.05","904.90"
"'NLNN3020186F5208","11.70","9.60","0","0.00","0.00","0","233.30","0.00","0.00","0.40","0.00","0.00","0","0","0","50.01","32.40","May.01 00:51,2019","0.05","904.90"
"'NLNN3020186F5208","11.70","9.60","0","0.00","0.00","0","233.30","0.00","0.00","0.40","0.00","0.00","0","0","0","50.01","32.40","May.01 00:57,2019","0.05","904.90"
`

	File10LinesOneEmpty = `"'NLNN3020186F5208","11.70","9.60","0","0.00","0.00","0","233.30","0.00","0.00","0.40","0.00","0.00","0","0","0","50.01","32.40","May.01 00:04,2019","0.05","904.90"
"'NLNN3020186F5208","11.70","9.60","0","0.00","0.00","0","233.30","0.00","0.00","0.40","0.00","0.00","0","0","0","50.01","32.40","May.01 00:09,2019","0.05","904.90"
"'NLNN3020186F5208","11.70","9.60","0","0.00","0.00","0","233.30","0.00","0.00","0.40","0.00","0.00","0","0","0","50.01","32.40","May.01 00:16,2019","0.05","904.90"
"'NLNN3020186F5208","11.70","9.60","0","0.00","0.00","0","233.30","0.00","0.00","0.40","0.00","0.00","0","0","0","50.01","32.40","May.01 00:21,2019","0.05","904.90"
"'NLNN3020186F5208","11.70","9.60","0","0.00","0.00","0","233.30","0.00","0.00","0.40","0.00","0.00","0","0","0","50.01","32.40","May.01 00:28,2019","0.05","904.90"
"'NLNN3020186F5208","11.70","9.60","0","0.00","0.00","0","233.30","0.00","0.00","0.40","0.00","0.00","0","0","0","50.01","32.40","May.01 00:33,2019","0.05","904.90"
"'NLNN3020186F5208","11.70","9.60","0","0.00","0.00","0","233.30","0.00","0.00","0.40","0.00","0.00","0","0","0","50.01","32.40","May.01 00:40,2019","0.05","904.90"
"'NLNN3020186F5208","11.70","9.60","0","0.00","0.00","0","233.30","0.00","0.00","0.40","0.00","0.00","0","0","0","50.01","32.40","May.01 00:45,2019","0.05","904.90"
"'NLNN3020186F5208","11.70","9.60","0","0.00","0.00","0","233.30","0.00","0.00","0.40","0.00","0.00","0","0","0","50.01","32.40","May.01 00:51,2019","0.05","904.90"
"'NLNN3020186F5208","11.70","9.60","0","0.00","0.00","0","233.30","0.00","0.00","0.40","0.00","0.00","0","0","0","50.01","32.40","May.01 00:57,2019","0.05","904.90"
asdfadfadf
`

	File10LinesBadStartAndEnd = `Weekly Report
Project Overview
Site Name:,Zandbergen Eijsden,Total Energy:,"1711.20kWh",Weekly electricity:,"55.00kWh"
System Size:,"6.720kWp",Time:,"May.01,2019/May.07,2019"

Summary
Inverter,"'NLNN3020186F5208"
Inverter,Vpv1,Vpv2,Vpv3,Ipv1,Ipv2,Ipv3,Vac1,Vac2,Vac3,Iac1,Iac2,Iac3,Pac1,Pac2,Pac3,Fac,Temperature(℃),Time(GMT +1),Today's Energy(kWh),Total Energy(kWh)
"'NLNN3020186F5208","11.70","9.60","0","0.00","0.00","0","233.30","0.00","0.00","0.40","0.00","0.00","0","0","0","50.01","32.40","May.01 00:04,2019","0.05","904.90"


Inverter,"'NLRN2520189J3042"
Inverter,Vpv1,Vpv2,Vpv3,Ipv1,Ipv2,Ipv3,Vac1,Vac2,Vac3,Iac1,Iac2,Iac3,Pac1,Pac2,Pac3,Fac,Temperature(℃),Time(GMT +1),Today's Energy(kWh),Total Energy(kWh)
"'NLRN2520189J3042","332.90","0.00","0","3.70","0.00","0","234.40","0.00","0.00","5.30","0.00","0.00","1207","0","0","50.02","43.20","May.03 12:49,2019","3.85","776.40"

Weather Data:
Gateway,Relative humidity,Radiation,Module temperature,Wind speed,Wind direction,Ambient temperature,PowerstationID,Time(GMT 2)
`

	File10LinesBadStart = `Weekly Report
Project Overview
Site Name:,Zandbergen Eijsden,Total Energy:,"1711.20kWh",Weekly electricity:,"55.00kWh"
System Size:,"6.720kWp",Time:,"May.01,2019/May.07,2019"

Summary
Inverter,"'NLNN3020186F5208"
Inverter,Vpv1,Vpv2,Vpv3,Ipv1,Ipv2,Ipv3,Vac1,Vac2,Vac3,Iac1,Iac2,Iac3,Pac1,Pac2,Pac3,Fac,Temperature(℃),Time(GMT +1),Today's Energy(kWh),Total Energy(kWh)
"'NLNN3020186F5208","11.70","9.60","0","0.00","0.00","0","233.30","0.00","0.00","0.40","0.00","0.00","0","0","0","50.01","32.40","May.01 00:04,2019","0.05","904.90"
`
)

func TestFile10Lines(t *testing.T) {

	r := bytes.NewBufferString(File10Lines)

	s := NewScanner(r)

	var i int
	var err error
	// Scan until error is not nil
	for _, err = s.Next(); err == nil; _, err = s.Next() {
		i++
	}
	if err != io.EOF {
		t.Fatalf("expected %s, got %s", io.EOF.Error(), err.Error())
	}
	const expected = 10
	if i != expected {
		t.Errorf("expected %d lines, got %d", expected, i)
	}
}

func TestFile10LinesOneEmpty(t *testing.T) {

	r := bytes.NewBufferString(File10LinesOneEmpty)

	s := NewScanner(r)

	var i int
	var err error
	// Scan until error is not nil
	for _, err = s.Next(); err == nil; _, err = s.Next() {
		i++
	}
	if err != io.EOF {
		t.Fatalf("expected %s, got %s", io.EOF.Error(), err.Error())
	}
	const expected = 10
	if i != expected {
		t.Errorf("expected %d lines, got %d", expected, i)
	}
}

// File10LinesBadStart
func TestFile10LinesBadStart(t *testing.T) {
	r := bytes.NewBufferString(File10LinesBadStart)

	s := NewScanner(r)

	var i int
	var err error
	// Scan until error is not nil
	for _, err = s.Next(); err == nil; _, err = s.Next() {
		i++
	}
	if err != io.EOF {
		t.Fatalf("expected %s, got %s", io.EOF.Error(), err.Error())
	}
	const expected = 1
	if i != expected {
		t.Errorf("expected %d lines, got %d", expected, i)
	}
}

func TestFile10LinesBadStartAndEnd(t *testing.T) {
	r := bytes.NewBufferString(File10LinesBadStartAndEnd)

	s := NewScanner(r)

	var i int
	var err error
	// Scan until error is not nil
	for _, err = s.Next(); err == nil; _, err = s.Next() {
		i++
	}
	if err != io.EOF {
		t.Fatalf("expected %s, got %s", io.EOF.Error(), err.Error())
	}
	const expected = 2
	if i != expected {
		t.Errorf("expected %d lines, got %d", expected, i)
	}
}
