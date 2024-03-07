package genapi

import "github.com/gogf/gf/v2/database/gdb"

func sortFieldKeyForDao(fieldMap map[string]*gdb.TableField) []string {
	names := make(map[int]string)
	for _, field := range fieldMap {
		names[field.Index] = field.Name
	}
	var (
		i      = 0
		j      = 0
		result = make([]string, len(names))
	)
	for {
		if len(names) == 0 {
			break
		}
		if val, ok := names[i]; ok {
			result[j] = val
			j++
			delete(names, i)
		}
		i++
	}
	return result
}


type generateStructDefinitionInput struct {
	CGenDaoInternalInput
	TableName  string                     // Table name.
	StructName string                     // Struct name.
	FieldMap   map[string]*gdb.TableField // Table field map.
	IsDo       bool                       // Is generating DTO struct.
}