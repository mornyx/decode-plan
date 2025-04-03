package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/pingcap/tidb/util/plancodec"
)

func main() {
	var buf bytes.Buffer
	_, err := io.Copy(&buf, os.Stdin)
	check(err)
	raw := strings.TrimSpace(buf.String())
	var r string
	if strings.HasPrefix(raw, "tidb_decode_plan('") {
		raw = strings.TrimPrefix(raw, "tidb_decode_plan('")
		raw = strings.TrimSuffix(raw, "')")
		r, err = plancodec.DecodePlan(raw)
		check(err)
	} else if strings.HasPrefix(raw, "tidb_decode_binary_plan('") {
		raw = strings.TrimPrefix(raw, "tidb_decode_binary_plan('")
		raw = strings.TrimSuffix(raw, "')")
		r, err = plancodec.DecodeBinaryPlan(raw)
		check(err)
	}
	fmt.Println(r)
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
