// Code generated by Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package databases

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Set of storage type classes for database instance
type MysqlStorageType string

const (
	MysqlStorageTypeStandard = MysqlStorageType("standard")
	MysqlStorageTypeIo1      = MysqlStorageType("io1")
	MysqlStorageTypeGp2      = MysqlStorageType("gp2")
)

func (MysqlStorageType) ElementType() reflect.Type {
	return reflect.TypeOf((*MysqlStorageType)(nil)).Elem()
}

func (e MysqlStorageType) ToMysqlStorageTypeOutput() MysqlStorageTypeOutput {
	return pulumi.ToOutput(e).(MysqlStorageTypeOutput)
}

func (e MysqlStorageType) ToMysqlStorageTypeOutputWithContext(ctx context.Context) MysqlStorageTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(MysqlStorageTypeOutput)
}

func (e MysqlStorageType) ToMysqlStorageTypePtrOutput() MysqlStorageTypePtrOutput {
	return e.ToMysqlStorageTypePtrOutputWithContext(context.Background())
}

func (e MysqlStorageType) ToMysqlStorageTypePtrOutputWithContext(ctx context.Context) MysqlStorageTypePtrOutput {
	return MysqlStorageType(e).ToMysqlStorageTypeOutputWithContext(ctx).ToMysqlStorageTypePtrOutputWithContext(ctx)
}

func (e MysqlStorageType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e MysqlStorageType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e MysqlStorageType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e MysqlStorageType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type MysqlStorageTypeOutput struct{ *pulumi.OutputState }

func (MysqlStorageTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MysqlStorageType)(nil)).Elem()
}

func (o MysqlStorageTypeOutput) ToMysqlStorageTypeOutput() MysqlStorageTypeOutput {
	return o
}

func (o MysqlStorageTypeOutput) ToMysqlStorageTypeOutputWithContext(ctx context.Context) MysqlStorageTypeOutput {
	return o
}

func (o MysqlStorageTypeOutput) ToMysqlStorageTypePtrOutput() MysqlStorageTypePtrOutput {
	return o.ToMysqlStorageTypePtrOutputWithContext(context.Background())
}

func (o MysqlStorageTypeOutput) ToMysqlStorageTypePtrOutputWithContext(ctx context.Context) MysqlStorageTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v MysqlStorageType) *MysqlStorageType {
		return &v
	}).(MysqlStorageTypePtrOutput)
}

func (o MysqlStorageTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o MysqlStorageTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e MysqlStorageType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o MysqlStorageTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o MysqlStorageTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e MysqlStorageType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type MysqlStorageTypePtrOutput struct{ *pulumi.OutputState }

func (MysqlStorageTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MysqlStorageType)(nil)).Elem()
}

func (o MysqlStorageTypePtrOutput) ToMysqlStorageTypePtrOutput() MysqlStorageTypePtrOutput {
	return o
}

func (o MysqlStorageTypePtrOutput) ToMysqlStorageTypePtrOutputWithContext(ctx context.Context) MysqlStorageTypePtrOutput {
	return o
}

func (o MysqlStorageTypePtrOutput) Elem() MysqlStorageTypeOutput {
	return o.ApplyT(func(v *MysqlStorageType) MysqlStorageType {
		if v != nil {
			return *v
		}
		var ret MysqlStorageType
		return ret
	}).(MysqlStorageTypeOutput)
}

func (o MysqlStorageTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o MysqlStorageTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *MysqlStorageType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// MysqlStorageTypeInput is an input type that accepts MysqlStorageTypeArgs and MysqlStorageTypeOutput values.
// You can construct a concrete instance of `MysqlStorageTypeInput` via:
//
//	MysqlStorageTypeArgs{...}
type MysqlStorageTypeInput interface {
	pulumi.Input

	ToMysqlStorageTypeOutput() MysqlStorageTypeOutput
	ToMysqlStorageTypeOutputWithContext(context.Context) MysqlStorageTypeOutput
}

var mysqlStorageTypePtrType = reflect.TypeOf((**MysqlStorageType)(nil)).Elem()

type MysqlStorageTypePtrInput interface {
	pulumi.Input

	ToMysqlStorageTypePtrOutput() MysqlStorageTypePtrOutput
	ToMysqlStorageTypePtrOutputWithContext(context.Context) MysqlStorageTypePtrOutput
}

type mysqlStorageTypePtr string

func MysqlStorageTypePtr(v string) MysqlStorageTypePtrInput {
	return (*mysqlStorageTypePtr)(&v)
}

func (*mysqlStorageTypePtr) ElementType() reflect.Type {
	return mysqlStorageTypePtrType
}

func (in *mysqlStorageTypePtr) ToMysqlStorageTypePtrOutput() MysqlStorageTypePtrOutput {
	return pulumi.ToOutput(in).(MysqlStorageTypePtrOutput)
}

func (in *mysqlStorageTypePtr) ToMysqlStorageTypePtrOutputWithContext(ctx context.Context) MysqlStorageTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(MysqlStorageTypePtrOutput)
}

// Set of allowed versions for the database instance
type MysqlVersion string

const (
	MysqlVersion_V8_0 = MysqlVersion("8.0")
	MysqlVersion_V5_7 = MysqlVersion("5.7")
)

func (MysqlVersion) ElementType() reflect.Type {
	return reflect.TypeOf((*MysqlVersion)(nil)).Elem()
}

func (e MysqlVersion) ToMysqlVersionOutput() MysqlVersionOutput {
	return pulumi.ToOutput(e).(MysqlVersionOutput)
}

func (e MysqlVersion) ToMysqlVersionOutputWithContext(ctx context.Context) MysqlVersionOutput {
	return pulumi.ToOutputWithContext(ctx, e).(MysqlVersionOutput)
}

func (e MysqlVersion) ToMysqlVersionPtrOutput() MysqlVersionPtrOutput {
	return e.ToMysqlVersionPtrOutputWithContext(context.Background())
}

func (e MysqlVersion) ToMysqlVersionPtrOutputWithContext(ctx context.Context) MysqlVersionPtrOutput {
	return MysqlVersion(e).ToMysqlVersionOutputWithContext(ctx).ToMysqlVersionPtrOutputWithContext(ctx)
}

func (e MysqlVersion) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e MysqlVersion) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e MysqlVersion) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e MysqlVersion) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type MysqlVersionOutput struct{ *pulumi.OutputState }

func (MysqlVersionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MysqlVersion)(nil)).Elem()
}

func (o MysqlVersionOutput) ToMysqlVersionOutput() MysqlVersionOutput {
	return o
}

func (o MysqlVersionOutput) ToMysqlVersionOutputWithContext(ctx context.Context) MysqlVersionOutput {
	return o
}

func (o MysqlVersionOutput) ToMysqlVersionPtrOutput() MysqlVersionPtrOutput {
	return o.ToMysqlVersionPtrOutputWithContext(context.Background())
}

func (o MysqlVersionOutput) ToMysqlVersionPtrOutputWithContext(ctx context.Context) MysqlVersionPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v MysqlVersion) *MysqlVersion {
		return &v
	}).(MysqlVersionPtrOutput)
}

func (o MysqlVersionOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o MysqlVersionOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e MysqlVersion) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o MysqlVersionOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o MysqlVersionOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e MysqlVersion) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type MysqlVersionPtrOutput struct{ *pulumi.OutputState }

func (MysqlVersionPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MysqlVersion)(nil)).Elem()
}

func (o MysqlVersionPtrOutput) ToMysqlVersionPtrOutput() MysqlVersionPtrOutput {
	return o
}

func (o MysqlVersionPtrOutput) ToMysqlVersionPtrOutputWithContext(ctx context.Context) MysqlVersionPtrOutput {
	return o
}

func (o MysqlVersionPtrOutput) Elem() MysqlVersionOutput {
	return o.ApplyT(func(v *MysqlVersion) MysqlVersion {
		if v != nil {
			return *v
		}
		var ret MysqlVersion
		return ret
	}).(MysqlVersionOutput)
}

func (o MysqlVersionPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o MysqlVersionPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *MysqlVersion) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// MysqlVersionInput is an input type that accepts MysqlVersionArgs and MysqlVersionOutput values.
// You can construct a concrete instance of `MysqlVersionInput` via:
//
//	MysqlVersionArgs{...}
type MysqlVersionInput interface {
	pulumi.Input

	ToMysqlVersionOutput() MysqlVersionOutput
	ToMysqlVersionOutputWithContext(context.Context) MysqlVersionOutput
}

var mysqlVersionPtrType = reflect.TypeOf((**MysqlVersion)(nil)).Elem()

type MysqlVersionPtrInput interface {
	pulumi.Input

	ToMysqlVersionPtrOutput() MysqlVersionPtrOutput
	ToMysqlVersionPtrOutputWithContext(context.Context) MysqlVersionPtrOutput
}

type mysqlVersionPtr string

func MysqlVersionPtr(v string) MysqlVersionPtrInput {
	return (*mysqlVersionPtr)(&v)
}

func (*mysqlVersionPtr) ElementType() reflect.Type {
	return mysqlVersionPtrType
}

func (in *mysqlVersionPtr) ToMysqlVersionPtrOutput() MysqlVersionPtrOutput {
	return pulumi.ToOutput(in).(MysqlVersionPtrOutput)
}

func (in *mysqlVersionPtr) ToMysqlVersionPtrOutputWithContext(ctx context.Context) MysqlVersionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(MysqlVersionPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*MysqlStorageTypeInput)(nil)).Elem(), MysqlStorageType("standard"))
	pulumi.RegisterInputType(reflect.TypeOf((*MysqlStorageTypePtrInput)(nil)).Elem(), MysqlStorageType("standard"))
	pulumi.RegisterInputType(reflect.TypeOf((*MysqlVersionInput)(nil)).Elem(), MysqlVersion("8.0"))
	pulumi.RegisterInputType(reflect.TypeOf((*MysqlVersionPtrInput)(nil)).Elem(), MysqlVersion("8.0"))
	pulumi.RegisterOutputType(MysqlStorageTypeOutput{})
	pulumi.RegisterOutputType(MysqlStorageTypePtrOutput{})
	pulumi.RegisterOutputType(MysqlVersionOutput{})
	pulumi.RegisterOutputType(MysqlVersionPtrOutput{})
}