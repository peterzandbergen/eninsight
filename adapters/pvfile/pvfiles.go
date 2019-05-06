package pvfile

import (
	"encoding/csv"
	"errors"
	"io"
	"strconv"
	"time"
)

type Scanner interface {
	Next() (*PvLine, error)
}

var (
	// ErrEndOfFile               = io.EOF
	ErrInvalidFormat           = errors.New("invalid format")
	ErrIncorrectNumberOfFields = errors.New("incorrect number of fields")
	ErrNotImplementedYet       = errors.New("not implemented yet")
)

const (
	NumFields = 21
	Separator = ","
)

const (
	fldInverter = iota
	fldVpv
	_
	_
	fldIpv
	_
	_
	fldVac
	_
	_
	fldIac
	_
	_
	fldPac
	_
	_
	fldFac
	fldTemperature
	fldTime
	fldTodaysEnergy
	fldTotalEnergy
)

// PvLine contains the converted data from one line from the PV Omnik Portal export file.
type PvLine struct {
	Inverter     string //
	Vpv          [3]float32
	Ipv          [3]float32
	Vac          [3]float32
	Iac          [3]float32
	Pac          [3]float32
	Fac          float32
	Temperature  float32
	Time         time.Time
	TodaysEnergy float32
	TotalEnergy  float32
}

// pvScanner reads and parses a PV exported file to PvLine structs.
// It implements the Scanner interface.
// TODO: Use encode.csv to read the file.
type pvScanner struct {
	s *csv.Reader
}

// Next returns the next PvLine or an error.
// After an error the scanner is no longer valid and should no longer be used.
func (p *pvScanner) Next() (*PvLine, error) {
	var pvl *PvLine

	var err error
	// Read till record found or End of file
	for pvl, err = p.next(); err != nil && err != io.EOF; pvl, err = p.next() {
		// TODO: Add logging.
	}
	if err == io.EOF {
		return nil, err
	}
	return pvl, nil
}

func (p *pvScanner) next() (*PvLine, error) {
	// Scan record.
	var r []string
	var err error
	if r, err = p.s.Read(); err != nil {
		return nil, err
	}

	return parseFields(r)
}

func parseFields(f []string) (*PvLine, error) {
	pvl := new(PvLine)

	// Parse the identifier and skip the single quote.
	pvl.Inverter = f[0][1:]
	if err := parseNFloat32(len(pvl.Vpv), pvl.Vpv[:], f[fldVpv:]); err != nil {
		return nil, err
	}
	if err := parseNFloat32(len(pvl.Ipv), pvl.Ipv[:], f[fldIpv:]); err != nil {
		return nil, err
	}
	if err := parseNFloat32(len(pvl.Vac), pvl.Vac[:], f[fldVac:]); err != nil {
		return nil, err
	}
	if err := parseNFloat32(len(pvl.Iac), pvl.Iac[:], f[fldIac:]); err != nil {
		return nil, err
	}
	if err := parseNFloat32(len(pvl.Pac), pvl.Pac[:], f[fldPac:]); err != nil {
		return nil, err
	}
	// Fac
	if x, err := strconv.ParseFloat(f[fldFac], 32); err != nil {
		return nil, err
	} else {
		pvl.Fac = float32(x)
	}
	// Temperature
	if x, err := strconv.ParseFloat(f[fldTemperature], 32); err != nil {
		return nil, err
	} else {
		pvl.Temperature = float32(x)
	}
	// TodaysEnergy
	if x, err := strconv.ParseFloat(f[fldTodaysEnergy], 32); err != nil {
		return nil, err
	} else {
		pvl.TodaysEnergy = float32(x)
	}
	// TotalEnergy
	if x, err := strconv.ParseFloat(f[fldTotalEnergy], 32); err != nil {
		return nil, err
	} else {
		pvl.TotalEnergy = float32(x)
	}
	// Time: May.01 00:04,2019
	if t, err := time.Parse("Jan.2 15:04,2006", f[fldTime]); err != nil {
		return nil, err
	} else {
		pvl.Time = t
	}
	return pvl, nil
}

func parseNFloat32(n int, f []float32, s []string) error {
	for i := 0; i < n; i++ {
		if x, err := strconv.ParseFloat(s[i], 32); err != nil {
			return err
		} else {
			f[i] = float32(x)
		}
	}
	return nil
}

func NewScanner(r io.Reader) Scanner {
	s := &pvScanner{
		s: csv.NewReader(r),
	}
	s.s.FieldsPerRecord = NumFields
	return s
}
