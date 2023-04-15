// Code generated by Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package storage

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Versioning object mode that is selected for the bucket.
// Enabled - Bucket keeps track of object versioning.
// Disabled - Bucket does not keep track of object versioning. Once versioning is Enabled we cannot Disabled versioning.
// Suspended - Once versioning is Enabled, this feature can be suspended
type BucketVersioningState string

const (
	BucketVersioningStateEnabled   = BucketVersioningState("Enabled")
	BucketVersioningStateDisabled  = BucketVersioningState("Disabled")
	BucketVersioningStateSuspended = BucketVersioningState("Suspended")
)

func (BucketVersioningState) ElementType() reflect.Type {
	return reflect.TypeOf((*BucketVersioningState)(nil)).Elem()
}

func (e BucketVersioningState) ToBucketVersioningStateOutput() BucketVersioningStateOutput {
	return pulumi.ToOutput(e).(BucketVersioningStateOutput)
}

func (e BucketVersioningState) ToBucketVersioningStateOutputWithContext(ctx context.Context) BucketVersioningStateOutput {
	return pulumi.ToOutputWithContext(ctx, e).(BucketVersioningStateOutput)
}

func (e BucketVersioningState) ToBucketVersioningStatePtrOutput() BucketVersioningStatePtrOutput {
	return e.ToBucketVersioningStatePtrOutputWithContext(context.Background())
}

func (e BucketVersioningState) ToBucketVersioningStatePtrOutputWithContext(ctx context.Context) BucketVersioningStatePtrOutput {
	return BucketVersioningState(e).ToBucketVersioningStateOutputWithContext(ctx).ToBucketVersioningStatePtrOutputWithContext(ctx)
}

func (e BucketVersioningState) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e BucketVersioningState) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e BucketVersioningState) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e BucketVersioningState) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type BucketVersioningStateOutput struct{ *pulumi.OutputState }

func (BucketVersioningStateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BucketVersioningState)(nil)).Elem()
}

func (o BucketVersioningStateOutput) ToBucketVersioningStateOutput() BucketVersioningStateOutput {
	return o
}

func (o BucketVersioningStateOutput) ToBucketVersioningStateOutputWithContext(ctx context.Context) BucketVersioningStateOutput {
	return o
}

func (o BucketVersioningStateOutput) ToBucketVersioningStatePtrOutput() BucketVersioningStatePtrOutput {
	return o.ToBucketVersioningStatePtrOutputWithContext(context.Background())
}

func (o BucketVersioningStateOutput) ToBucketVersioningStatePtrOutputWithContext(ctx context.Context) BucketVersioningStatePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v BucketVersioningState) *BucketVersioningState {
		return &v
	}).(BucketVersioningStatePtrOutput)
}

func (o BucketVersioningStateOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o BucketVersioningStateOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e BucketVersioningState) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o BucketVersioningStateOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o BucketVersioningStateOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e BucketVersioningState) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type BucketVersioningStatePtrOutput struct{ *pulumi.OutputState }

func (BucketVersioningStatePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BucketVersioningState)(nil)).Elem()
}

func (o BucketVersioningStatePtrOutput) ToBucketVersioningStatePtrOutput() BucketVersioningStatePtrOutput {
	return o
}

func (o BucketVersioningStatePtrOutput) ToBucketVersioningStatePtrOutputWithContext(ctx context.Context) BucketVersioningStatePtrOutput {
	return o
}

func (o BucketVersioningStatePtrOutput) Elem() BucketVersioningStateOutput {
	return o.ApplyT(func(v *BucketVersioningState) BucketVersioningState {
		if v != nil {
			return *v
		}
		var ret BucketVersioningState
		return ret
	}).(BucketVersioningStateOutput)
}

func (o BucketVersioningStatePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o BucketVersioningStatePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *BucketVersioningState) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// BucketVersioningStateInput is an input type that accepts BucketVersioningStateArgs and BucketVersioningStateOutput values.
// You can construct a concrete instance of `BucketVersioningStateInput` via:
//
//	BucketVersioningStateArgs{...}
type BucketVersioningStateInput interface {
	pulumi.Input

	ToBucketVersioningStateOutput() BucketVersioningStateOutput
	ToBucketVersioningStateOutputWithContext(context.Context) BucketVersioningStateOutput
}

var bucketVersioningStatePtrType = reflect.TypeOf((**BucketVersioningState)(nil)).Elem()

type BucketVersioningStatePtrInput interface {
	pulumi.Input

	ToBucketVersioningStatePtrOutput() BucketVersioningStatePtrOutput
	ToBucketVersioningStatePtrOutputWithContext(context.Context) BucketVersioningStatePtrOutput
}

type bucketVersioningStatePtr string

func BucketVersioningStatePtr(v string) BucketVersioningStatePtrInput {
	return (*bucketVersioningStatePtr)(&v)
}

func (*bucketVersioningStatePtr) ElementType() reflect.Type {
	return bucketVersioningStatePtrType
}

func (in *bucketVersioningStatePtr) ToBucketVersioningStatePtrOutput() BucketVersioningStatePtrOutput {
	return pulumi.ToOutput(in).(BucketVersioningStatePtrOutput)
}

func (in *bucketVersioningStatePtr) ToBucketVersioningStatePtrOutputWithContext(ctx context.Context) BucketVersioningStatePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(BucketVersioningStatePtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*BucketVersioningStateInput)(nil)).Elem(), BucketVersioningState("Enabled"))
	pulumi.RegisterInputType(reflect.TypeOf((*BucketVersioningStatePtrInput)(nil)).Elem(), BucketVersioningState("Enabled"))
	pulumi.RegisterOutputType(BucketVersioningStateOutput{})
	pulumi.RegisterOutputType(BucketVersioningStatePtrOutput{})
}