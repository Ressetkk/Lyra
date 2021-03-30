package score

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
)

func Encode(s *Score) (string, error) {
	var buffer bytes.Buffer
	jNotes, err := json.Marshal(s)
	if err != nil {
		return "", err
	}
	gz := gzip.NewWriter(&buffer)

	if _, err := gz.Write(jNotes); err != nil {
		return "", err
	}
	if err := gz.Close(); err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(buffer.Bytes()), nil
}

func Decode(exp string) (*Score, error) {
	dec, err := base64.StdEncoding.DecodeString(exp)
	if err != nil {
		return nil, DecodeError{err}
	}
	r, err := gzip.NewReader(bytes.NewReader(dec))
	if err != nil {
		return nil, DecodeError{err}
	}
	res, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, DecodeError{err}
	}

	var score Score
	if err := json.Unmarshal(res, &score); err != nil {
		return nil, DecodeError{err}
	}
	return &score, nil
}
