package utils

import (
	"fmt"
	"go/ast"
	"os"
	"reflect"
	"strings"
	"unicode"
)

func TypeToString(expr ast.Expr) string {
	switch t := expr.(type) {
	case *ast.Ident:
		// 기본 타입이나 커스텀 타입일 경우
		return t.Name
	case *ast.ArrayType:
		// 배열 타입
		return fmt.Sprintf("[]%s", TypeToString(t.Elt))
	case *ast.MapType:
		// 맵 타입
		return fmt.Sprintf("map[%s]%s", TypeToString(t.Key), TypeToString(t.Value))
	case *ast.StructType:
		// 구조체 타입
		return "struct"
	case *ast.FuncType:
		// 함수 타입
		return "func"
	// 기타 타입 처리 추가 가능
	default:
		return fmt.Sprintf("unknown type: %T", t)
	}
}

func ToSnakeCase(s string) string {
	var result strings.Builder

	for i, char := range s {
		if unicode.IsUpper(char) {
			// 첫 번째 글자이거나 이전 글자가 대문자가 아닐 경우 언더스코어 추가
			if i > 0 && !unicode.IsUpper(rune(s[i-1])) {
				result.WriteRune('_')
			}
			result.WriteRune(unicode.ToLower(char))
		} else {
			result.WriteRune(char)
		}
	}

	return result.String()
}

func MkDir(d string) error {
	_, err := os.Stat(d)
	if os.IsNotExist(err) {
		err = os.MkdirAll(d, os.ModePerm)
		if err != nil {
			return err
		}
	}
	return nil
}

func IsValidType(typeStr string) bool {
	// 기본 타입 목록
	validTypes := map[string]reflect.Kind{
		"string":  reflect.String,
		"int":     reflect.Int,
		"int8":    reflect.Int8,
		"int16":   reflect.Int16,
		"int32":   reflect.Int32,
		"int64":   reflect.Int64,
		"uint":    reflect.Uint,
		"uint8":   reflect.Uint8,
		"uint16":  reflect.Uint16,
		"uint32":  reflect.Uint32,
		"uint64":  reflect.Uint64,
		"float32": reflect.Float32,
		"float64": reflect.Float64,
		"bool":    reflect.Bool,
		"rune":    reflect.Int32, // rune은 int32의 별칭
		"byte":    reflect.Uint8, // byte는 uint8의 별칭
	}

	// 해당 타입이 유효한지 확인
	_, exists := validTypes[typeStr]
	return exists
}
