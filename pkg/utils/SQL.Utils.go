package utils

import (
	"database/sql"
	"fmt"
	"strings"
)

type USPOptions struct {
	name string
	// optional
	inputParams  []sql.NamedArg
	outputParams []sql.NamedArg
}

func NewUSPOptions(name string) USPOptions {
	s := USPOptions{}
	s.name = name
	return s
}

func (s USPOptions) WithInputParams(input ...sql.NamedArg) USPOptions {
	s.inputParams = input
	return s
}

/*
Example:

	_, err := db.ExecContext(ctx, "ProcName", sql.Named("Arg1", sql.Out{Dest: &outArg}))
*/
func (s USPOptions) WithOutputParams(output ...sql.NamedArg) USPOptions {
	s.outputParams = output
	return s
}

func (s *USPOptions) GetProcName() string {
	return s.name
}

func (s *USPOptions) GetProcedureQuery() string {
	var args []string

	for _, item := range s.inputParams {
		args = append(args, "@"+item.Name)
	}

	for _, item := range s.outputParams {
		args = append(args, "@"+item.Name+" OUTPUT")
	}

	return fmt.Sprintf("exec %v %v", s.GetProcName(), strings.Join(args, ","))

}

func (s *USPOptions) GetParams() []any {
	var result []any
	if len(s.inputParams) > 0 {
		for _, item := range s.inputParams {
			result = append(result, item)
		}
	}
	if len(s.outputParams) > 0 {
		for _, item := range s.outputParams {
			result = append(result, item)
		}
	}
	return result
}
