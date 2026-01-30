package main

// #cgo CFLAGS: -I/opt/halon/include
// #cgo LDFLAGS: -Wl,--unresolved-symbols=ignore-all
// #include <HalonMTA.h>
// #include <stdlib.h>
// #include <stdint.h>
import "C"
import (
	"errors"
	"fmt"
	"unsafe"

	"github.com/linkedin/goavro/v2"
)

func main() {}

func GetArgumentAsString(args *C.HalonHSLArguments, pos uint64, required bool) (string, error) {
	var x = C.HalonMTA_hsl_argument_get(args, C.ulong(pos))
	if x == nil {
		if required {
			return "", fmt.Errorf("missing argument at position %d", pos)
		} else {
			return "", nil
		}
	}
	var y *C.char
	var l C.size_t
	if C.HalonMTA_hsl_value_get(x, C.HALONMTA_HSL_TYPE_STRING, unsafe.Pointer(&y), &l) {
		return string(C.GoBytes(unsafe.Pointer(y), C.int(l))), nil
	} else {
		return "", fmt.Errorf("invalid argument at position %d", pos)
	}
}

func GetArgumentAsJSON(args *C.HalonHSLArguments, pos uint64, required bool) (string, error) {
	var x = C.HalonMTA_hsl_argument_get(args, C.ulong(pos))
	if x == nil {
		if required {
			return "", fmt.Errorf("missing argument at position %d", pos)
		} else {
			return "", nil
		}
	}
	var y *C.char
	z := C.HalonMTA_hsl_value_to_json(x, &y, nil)
	defer C.free(unsafe.Pointer(y))
	if z {
		return C.GoString(y), nil
	} else {
		return "", fmt.Errorf("invalid argument at position %d", pos)
	}
}

func SetException(hhc *C.HalonHSLContext, msg string) {
	x := C.CString(msg)
	y := unsafe.Pointer(x)
	defer C.free(y)
	exception := C.HalonMTA_hsl_throw(hhc)
	C.HalonMTA_hsl_value_set(exception, C.HALONMTA_HSL_TYPE_EXCEPTION, y, 0)
}

func SetReturnValueToString(ret *C.HalonHSLValue, val string) {
	x := C.CString(val)
	y := unsafe.Pointer(x)
	defer C.free(y)
	C.HalonMTA_hsl_value_set(ret, C.HALONMTA_HSL_TYPE_STRING, y, C.size_t(len(val)))
}

func SetReturnValueFromJson(ret *C.HalonHSLValue, json string) error {
	y := C.CString(json)
	defer C.free(unsafe.Pointer(y))
	var z *C.char
	if !(C.HalonMTA_hsl_value_from_json(ret, y, &z, nil)) {
		if z != nil {
			err := errors.New(C.GoString(z))
			C.free(unsafe.Pointer(z))
			return err
		}
		return errors.New("failed to parse return value")
	}
	return nil
}

//export Halon_version
func Halon_version() C.int {
	return C.HALONMTA_PLUGIN_VERSION
}

//export avro_encode
func avro_encode(hhc *C.HalonHSLContext, args *C.HalonHSLArguments, ret *C.HalonHSLValue) {
	schema, err := GetArgumentAsString(args, 0, true)
	if err != nil {
		SetException(hhc, err.Error())
		return
	}

	json, err := GetArgumentAsJSON(args, 1, true)
	if err != nil {
		SetException(hhc, err.Error())
		return
	}

	codec, err := goavro.NewCodec(schema)
	if err != nil {
		SetException(hhc, err.Error())
		return
	}

	native, _, err := codec.NativeFromTextual([]byte(json))
	if err != nil {
		SetException(hhc, err.Error())
		return
	}

	binary, err := codec.BinaryFromNative(nil, native)
	if err != nil {
		SetException(hhc, err.Error())
		return
	}

	SetReturnValueToString(ret, string(binary))
}

//export avro_decode
func avro_decode(hhc *C.HalonHSLContext, args *C.HalonHSLArguments, ret *C.HalonHSLValue) {
	schema, err := GetArgumentAsString(args, 0, true)
	if err != nil {
		SetException(hhc, err.Error())
		return
	}

	avro, err := GetArgumentAsString(args, 1, true)
	if err != nil {
		SetException(hhc, err.Error())
		return
	}

	codec, err := goavro.NewCodec(schema)
	if err != nil {
		SetException(hhc, err.Error())
		return
	}

	native, _, err := codec.NativeFromBinary([]byte(avro))
	if err != nil {
		SetException(hhc, err.Error())
		return
	}

	textual, err := codec.TextualFromNative(nil, native)
	if err != nil {
		SetException(hhc, err.Error())
		return
	}

	err = SetReturnValueFromJson(ret, string(textual))
	if err != nil {
		SetException(hhc, err.Error())
		return
	}
}

//export Halon_hsl_register
func Halon_hsl_register(hhrc *C.HalonHSLRegisterContext) C.bool {
	avro_encode_cs := C.CString("avro_encode")
	C.HalonMTA_hsl_module_register_function(hhrc, avro_encode_cs, nil)
	avro_decode_cs := C.CString("avro_decode")
	C.HalonMTA_hsl_module_register_function(hhrc, avro_decode_cs, nil)
	return true
}
