// Code generated by Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package kubernetes

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The subnet type
type ClusterSubnetsType string

const (
	ClusterSubnetsTypePrivate = ClusterSubnetsType("private")
	ClusterSubnetsTypePublic  = ClusterSubnetsType("public")
)

func (ClusterSubnetsType) ElementType() reflect.Type {
	return reflect.TypeOf((*ClusterSubnetsType)(nil)).Elem()
}

func (e ClusterSubnetsType) ToClusterSubnetsTypeOutput() ClusterSubnetsTypeOutput {
	return pulumi.ToOutput(e).(ClusterSubnetsTypeOutput)
}

func (e ClusterSubnetsType) ToClusterSubnetsTypeOutputWithContext(ctx context.Context) ClusterSubnetsTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ClusterSubnetsTypeOutput)
}

func (e ClusterSubnetsType) ToClusterSubnetsTypePtrOutput() ClusterSubnetsTypePtrOutput {
	return e.ToClusterSubnetsTypePtrOutputWithContext(context.Background())
}

func (e ClusterSubnetsType) ToClusterSubnetsTypePtrOutputWithContext(ctx context.Context) ClusterSubnetsTypePtrOutput {
	return ClusterSubnetsType(e).ToClusterSubnetsTypeOutputWithContext(ctx).ToClusterSubnetsTypePtrOutputWithContext(ctx)
}

func (e ClusterSubnetsType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ClusterSubnetsType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ClusterSubnetsType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ClusterSubnetsType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ClusterSubnetsTypeOutput struct{ *pulumi.OutputState }

func (ClusterSubnetsTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ClusterSubnetsType)(nil)).Elem()
}

func (o ClusterSubnetsTypeOutput) ToClusterSubnetsTypeOutput() ClusterSubnetsTypeOutput {
	return o
}

func (o ClusterSubnetsTypeOutput) ToClusterSubnetsTypeOutputWithContext(ctx context.Context) ClusterSubnetsTypeOutput {
	return o
}

func (o ClusterSubnetsTypeOutput) ToClusterSubnetsTypePtrOutput() ClusterSubnetsTypePtrOutput {
	return o.ToClusterSubnetsTypePtrOutputWithContext(context.Background())
}

func (o ClusterSubnetsTypeOutput) ToClusterSubnetsTypePtrOutputWithContext(ctx context.Context) ClusterSubnetsTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ClusterSubnetsType) *ClusterSubnetsType {
		return &v
	}).(ClusterSubnetsTypePtrOutput)
}

func (o ClusterSubnetsTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ClusterSubnetsTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ClusterSubnetsType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ClusterSubnetsTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ClusterSubnetsTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ClusterSubnetsType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ClusterSubnetsTypePtrOutput struct{ *pulumi.OutputState }

func (ClusterSubnetsTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ClusterSubnetsType)(nil)).Elem()
}

func (o ClusterSubnetsTypePtrOutput) ToClusterSubnetsTypePtrOutput() ClusterSubnetsTypePtrOutput {
	return o
}

func (o ClusterSubnetsTypePtrOutput) ToClusterSubnetsTypePtrOutputWithContext(ctx context.Context) ClusterSubnetsTypePtrOutput {
	return o
}

func (o ClusterSubnetsTypePtrOutput) Elem() ClusterSubnetsTypeOutput {
	return o.ApplyT(func(v *ClusterSubnetsType) ClusterSubnetsType {
		if v != nil {
			return *v
		}
		var ret ClusterSubnetsType
		return ret
	}).(ClusterSubnetsTypeOutput)
}

func (o ClusterSubnetsTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ClusterSubnetsTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ClusterSubnetsType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// ClusterSubnetsTypeInput is an input type that accepts ClusterSubnetsTypeArgs and ClusterSubnetsTypeOutput values.
// You can construct a concrete instance of `ClusterSubnetsTypeInput` via:
//
//	ClusterSubnetsTypeArgs{...}
type ClusterSubnetsTypeInput interface {
	pulumi.Input

	ToClusterSubnetsTypeOutput() ClusterSubnetsTypeOutput
	ToClusterSubnetsTypeOutputWithContext(context.Context) ClusterSubnetsTypeOutput
}

var clusterSubnetsTypePtrType = reflect.TypeOf((**ClusterSubnetsType)(nil)).Elem()

type ClusterSubnetsTypePtrInput interface {
	pulumi.Input

	ToClusterSubnetsTypePtrOutput() ClusterSubnetsTypePtrOutput
	ToClusterSubnetsTypePtrOutputWithContext(context.Context) ClusterSubnetsTypePtrOutput
}

type clusterSubnetsTypePtr string

func ClusterSubnetsTypePtr(v string) ClusterSubnetsTypePtrInput {
	return (*clusterSubnetsTypePtr)(&v)
}

func (*clusterSubnetsTypePtr) ElementType() reflect.Type {
	return clusterSubnetsTypePtrType
}

func (in *clusterSubnetsTypePtr) ToClusterSubnetsTypePtrOutput() ClusterSubnetsTypePtrOutput {
	return pulumi.ToOutput(in).(ClusterSubnetsTypePtrOutput)
}

func (in *clusterSubnetsTypePtr) ToClusterSubnetsTypePtrOutputWithContext(ctx context.Context) ClusterSubnetsTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ClusterSubnetsTypePtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ClusterSubnetsTypeInput)(nil)).Elem(), ClusterSubnetsType("private"))
	pulumi.RegisterInputType(reflect.TypeOf((*ClusterSubnetsTypePtrInput)(nil)).Elem(), ClusterSubnetsType("private"))
	pulumi.RegisterOutputType(ClusterSubnetsTypeOutput{})
	pulumi.RegisterOutputType(ClusterSubnetsTypePtrOutput{})
}