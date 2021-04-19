// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package followschema

import (
	"context"
	"strconv"
	"sync/atomic"

	"github.com/99designs/gqlgen/codegen/testserver/followschema/otherpkg"
	"github.com/99designs/gqlgen/graphql"
	"github.com/vektah/gqlparser/v2/ast"
)

// region    ************************** generated!.gotpl **************************

type WrappedMapResolver interface {
	Get(ctx context.Context, obj WrappedMap, key string) (string, error)
}
type WrappedSliceResolver interface {
	Get(ctx context.Context, obj WrappedSlice, idx int) (string, error)
}

// endregion ************************** generated!.gotpl **************************

// region    ***************************** args.gotpl *****************************

func (ec *executionContext) field_WrappedMap_get_args(ctx context.Context, rawArgs map[string]interface{}) (map[string]interface{}, error) {
	var err error
	args := map[string]interface{}{}
	var arg0 string
	if tmp, ok := rawArgs["key"]; ok {
		ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("key"))
		arg0, err = ec.unmarshalNString2string(ctx, tmp)
		if err != nil {
			return nil, err
		}
	}
	args["key"] = arg0
	return args, nil
}

func (ec *executionContext) field_WrappedSlice_get_args(ctx context.Context, rawArgs map[string]interface{}) (map[string]interface{}, error) {
	var err error
	args := map[string]interface{}{}
	var arg0 int
	if tmp, ok := rawArgs["idx"]; ok {
		ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("idx"))
		arg0, err = ec.unmarshalNInt2int(ctx, tmp)
		if err != nil {
			return nil, err
		}
	}
	args["idx"] = arg0
	return args, nil
}

// endregion ***************************** args.gotpl *****************************

// region    ************************** directives.gotpl **************************

// endregion ************************** directives.gotpl **************************

// region    **************************** field.gotpl *****************************

func (ec *executionContext) _WrappedMap_get(ctx context.Context, field graphql.CollectedField, obj WrappedMap) (ret graphql.Marshaler) {
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	fc := &graphql.FieldContext{
		Object:     "WrappedMap",
		Field:      field,
		Args:       nil,
		IsMethod:   true,
		IsResolver: true,
	}

	ctx = graphql.WithFieldContext(ctx, fc)
	rawArgs := field.ArgumentMap(ec.Variables)
	args, err := ec.field_WrappedMap_get_args(ctx, rawArgs)
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	fc.Args = args
	resTmp := ec._fieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return ec.resolvers.WrappedMap().Get(rctx, obj, args["key"].(string))
	})

	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(string)
	fc.Result = res
	return ec.marshalNString2string(ctx, field.Selections, res)
}

func (ec *executionContext) _WrappedSlice_get(ctx context.Context, field graphql.CollectedField, obj WrappedSlice) (ret graphql.Marshaler) {
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	fc := &graphql.FieldContext{
		Object:     "WrappedSlice",
		Field:      field,
		Args:       nil,
		IsMethod:   true,
		IsResolver: true,
	}

	ctx = graphql.WithFieldContext(ctx, fc)
	rawArgs := field.ArgumentMap(ec.Variables)
	args, err := ec.field_WrappedSlice_get_args(ctx, rawArgs)
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	fc.Args = args
	resTmp := ec._fieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return ec.resolvers.WrappedSlice().Get(rctx, obj, args["idx"].(int))
	})

	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(string)
	fc.Result = res
	return ec.marshalNString2string(ctx, field.Selections, res)
}

func (ec *executionContext) _WrappedStruct_name(ctx context.Context, field graphql.CollectedField, obj *WrappedStruct) (ret graphql.Marshaler) {
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	fc := &graphql.FieldContext{
		Object:     "WrappedStruct",
		Field:      field,
		Args:       nil,
		IsMethod:   false,
		IsResolver: false,
	}

	ctx = graphql.WithFieldContext(ctx, fc)
	resTmp := ec._fieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Name, nil
	})

	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(otherpkg.Scalar)
	fc.Result = res
	return ec.marshalNWrappedScalar2githubᚗcomᚋ99designsᚋgqlgenᚋcodegenᚋtestserverᚋfollowschemaᚋotherpkgᚐScalar(ctx, field.Selections, res)
}

func (ec *executionContext) _WrappedStruct_desc(ctx context.Context, field graphql.CollectedField, obj *WrappedStruct) (ret graphql.Marshaler) {
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	fc := &graphql.FieldContext{
		Object:     "WrappedStruct",
		Field:      field,
		Args:       nil,
		IsMethod:   false,
		IsResolver: false,
	}

	ctx = graphql.WithFieldContext(ctx, fc)
	resTmp := ec._fieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.Desc, nil
	})

	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(*otherpkg.Scalar)
	fc.Result = res
	return ec.marshalOWrappedScalar2ᚖgithubᚗcomᚋ99designsᚋgqlgenᚋcodegenᚋtestserverᚋfollowschemaᚋotherpkgᚐScalar(ctx, field.Selections, res)
}

// endregion **************************** field.gotpl *****************************

// region    **************************** input.gotpl *****************************

// endregion **************************** input.gotpl *****************************

// region    ************************** interface.gotpl ***************************

// endregion ************************** interface.gotpl ***************************

// region    **************************** object.gotpl ****************************

var wrappedMapImplementors = []string{"WrappedMap"}

func (ec *executionContext) _WrappedMap(ctx context.Context, sel ast.SelectionSet, obj WrappedMap) graphql.Marshaler {
	fields := graphql.CollectFields(ec.OperationContext, sel, wrappedMapImplementors)

	out := graphql.NewFieldSet(fields)
	var invalids uint32
	for i, field := range fields {
		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("WrappedMap")
		case "get":
			field := field
			out.Concurrently(i, func() (res graphql.Marshaler) {
				defer func() {
					if r := recover(); r != nil {
						ec.Error(ctx, ec.Recover(ctx, r))
					}
				}()
				res = ec._WrappedMap_get(ctx, field, obj)
				if res == graphql.Null {
					atomic.AddUint32(&invalids, 1)
				}
				return res
			})
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}
	out.Dispatch()
	if invalids > 0 {
		return graphql.Null
	}
	return out
}

var wrappedSliceImplementors = []string{"WrappedSlice"}

func (ec *executionContext) _WrappedSlice(ctx context.Context, sel ast.SelectionSet, obj WrappedSlice) graphql.Marshaler {
	fields := graphql.CollectFields(ec.OperationContext, sel, wrappedSliceImplementors)

	out := graphql.NewFieldSet(fields)
	var invalids uint32
	for i, field := range fields {
		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("WrappedSlice")
		case "get":
			field := field
			out.Concurrently(i, func() (res graphql.Marshaler) {
				defer func() {
					if r := recover(); r != nil {
						ec.Error(ctx, ec.Recover(ctx, r))
					}
				}()
				res = ec._WrappedSlice_get(ctx, field, obj)
				if res == graphql.Null {
					atomic.AddUint32(&invalids, 1)
				}
				return res
			})
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}
	out.Dispatch()
	if invalids > 0 {
		return graphql.Null
	}
	return out
}

var wrappedStructImplementors = []string{"WrappedStruct"}

func (ec *executionContext) _WrappedStruct(ctx context.Context, sel ast.SelectionSet, obj *WrappedStruct) graphql.Marshaler {
	fields := graphql.CollectFields(ec.OperationContext, sel, wrappedStructImplementors)

	out := graphql.NewFieldSet(fields)
	var invalids uint32
	for i, field := range fields {
		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("WrappedStruct")
		case "name":
			out.Values[i] = ec._WrappedStruct_name(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				invalids++
			}
		case "desc":
			out.Values[i] = ec._WrappedStruct_desc(ctx, field, obj)
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}
	out.Dispatch()
	if invalids > 0 {
		return graphql.Null
	}
	return out
}

// endregion **************************** object.gotpl ****************************

// region    ***************************** type.gotpl *****************************

func (ec *executionContext) marshalNWrappedMap2githubᚗcomᚋ99designsᚋgqlgenᚋcodegenᚋtestserverᚋfollowschemaᚐWrappedMap(ctx context.Context, sel ast.SelectionSet, v WrappedMap) graphql.Marshaler {
	if v == nil {
		if !graphql.HasFieldError(ctx, graphql.GetFieldContext(ctx)) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	return ec._WrappedMap(ctx, sel, v)
}

func (ec *executionContext) unmarshalNWrappedScalar2githubᚗcomᚋ99designsᚋgqlgenᚋcodegenᚋtestserverᚋfollowschemaᚋotherpkgᚐScalar(ctx context.Context, v interface{}) (otherpkg.Scalar, error) {
	tmp, err := graphql.UnmarshalString(v)
	res := otherpkg.Scalar(tmp)
	return res, graphql.ErrorOnPath(ctx, err)
}

func (ec *executionContext) marshalNWrappedScalar2githubᚗcomᚋ99designsᚋgqlgenᚋcodegenᚋtestserverᚋfollowschemaᚋotherpkgᚐScalar(ctx context.Context, sel ast.SelectionSet, v otherpkg.Scalar) graphql.Marshaler {
	res := graphql.MarshalString(string(v))
	if res == graphql.Null {
		if !graphql.HasFieldError(ctx, graphql.GetFieldContext(ctx)) {
			ec.Errorf(ctx, "must not be null")
		}
	}
	return res
}

func (ec *executionContext) marshalNWrappedSlice2githubᚗcomᚋ99designsᚋgqlgenᚋcodegenᚋtestserverᚋfollowschemaᚐWrappedSlice(ctx context.Context, sel ast.SelectionSet, v WrappedSlice) graphql.Marshaler {
	if v == nil {
		if !graphql.HasFieldError(ctx, graphql.GetFieldContext(ctx)) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	return ec._WrappedSlice(ctx, sel, v)
}

func (ec *executionContext) marshalNWrappedStruct2githubᚗcomᚋ99designsᚋgqlgenᚋcodegenᚋtestserverᚋfollowschemaᚐWrappedStruct(ctx context.Context, sel ast.SelectionSet, v WrappedStruct) graphql.Marshaler {
	return ec._WrappedStruct(ctx, sel, &v)
}

func (ec *executionContext) marshalNWrappedStruct2ᚖgithubᚗcomᚋ99designsᚋgqlgenᚋcodegenᚋtestserverᚋfollowschemaᚐWrappedStruct(ctx context.Context, sel ast.SelectionSet, v *WrappedStruct) graphql.Marshaler {
	if v == nil {
		if !graphql.HasFieldError(ctx, graphql.GetFieldContext(ctx)) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	return ec._WrappedStruct(ctx, sel, v)
}

func (ec *executionContext) unmarshalOWrappedScalar2ᚖgithubᚗcomᚋ99designsᚋgqlgenᚋcodegenᚋtestserverᚋfollowschemaᚋotherpkgᚐScalar(ctx context.Context, v interface{}) (*otherpkg.Scalar, error) {
	if v == nil {
		return nil, nil
	}
	tmp, err := graphql.UnmarshalString(v)
	res := otherpkg.Scalar(tmp)
	return &res, graphql.ErrorOnPath(ctx, err)
}

func (ec *executionContext) marshalOWrappedScalar2ᚖgithubᚗcomᚋ99designsᚋgqlgenᚋcodegenᚋtestserverᚋfollowschemaᚋotherpkgᚐScalar(ctx context.Context, sel ast.SelectionSet, v *otherpkg.Scalar) graphql.Marshaler {
	if v == nil {
		return graphql.Null
	}
	return graphql.MarshalString(string(*v))
}

// endregion ***************************** type.gotpl *****************************
