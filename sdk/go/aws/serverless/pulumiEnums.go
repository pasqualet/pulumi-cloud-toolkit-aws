// Code generated by Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package serverless

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Dead Letter Queue type that will receive the faulty messages from the base Queue.
// Permissive - Messages will be sent to the Dead Letter Queue after 10 failed delivery attempts.
// Restrictive - Messages will be sent to the Dead Letter Queue after the first failed delivery attempt.
type DeadLetterQueueTypes string

const (
	DeadLetterQueueTypesPERMISSIVE  = DeadLetterQueueTypes("permissive")
	DeadLetterQueueTypesRESTRICTIVE = DeadLetterQueueTypes("restrictive")
)

func (DeadLetterQueueTypes) ElementType() reflect.Type {
	return reflect.TypeOf((*DeadLetterQueueTypes)(nil)).Elem()
}

func (e DeadLetterQueueTypes) ToDeadLetterQueueTypesOutput() DeadLetterQueueTypesOutput {
	return pulumi.ToOutput(e).(DeadLetterQueueTypesOutput)
}

func (e DeadLetterQueueTypes) ToDeadLetterQueueTypesOutputWithContext(ctx context.Context) DeadLetterQueueTypesOutput {
	return pulumi.ToOutputWithContext(ctx, e).(DeadLetterQueueTypesOutput)
}

func (e DeadLetterQueueTypes) ToDeadLetterQueueTypesPtrOutput() DeadLetterQueueTypesPtrOutput {
	return e.ToDeadLetterQueueTypesPtrOutputWithContext(context.Background())
}

func (e DeadLetterQueueTypes) ToDeadLetterQueueTypesPtrOutputWithContext(ctx context.Context) DeadLetterQueueTypesPtrOutput {
	return DeadLetterQueueTypes(e).ToDeadLetterQueueTypesOutputWithContext(ctx).ToDeadLetterQueueTypesPtrOutputWithContext(ctx)
}

func (e DeadLetterQueueTypes) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e DeadLetterQueueTypes) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e DeadLetterQueueTypes) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e DeadLetterQueueTypes) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type DeadLetterQueueTypesOutput struct{ *pulumi.OutputState }

func (DeadLetterQueueTypesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DeadLetterQueueTypes)(nil)).Elem()
}

func (o DeadLetterQueueTypesOutput) ToDeadLetterQueueTypesOutput() DeadLetterQueueTypesOutput {
	return o
}

func (o DeadLetterQueueTypesOutput) ToDeadLetterQueueTypesOutputWithContext(ctx context.Context) DeadLetterQueueTypesOutput {
	return o
}

func (o DeadLetterQueueTypesOutput) ToDeadLetterQueueTypesPtrOutput() DeadLetterQueueTypesPtrOutput {
	return o.ToDeadLetterQueueTypesPtrOutputWithContext(context.Background())
}

func (o DeadLetterQueueTypesOutput) ToDeadLetterQueueTypesPtrOutputWithContext(ctx context.Context) DeadLetterQueueTypesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DeadLetterQueueTypes) *DeadLetterQueueTypes {
		return &v
	}).(DeadLetterQueueTypesPtrOutput)
}

func (o DeadLetterQueueTypesOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o DeadLetterQueueTypesOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e DeadLetterQueueTypes) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o DeadLetterQueueTypesOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o DeadLetterQueueTypesOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e DeadLetterQueueTypes) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type DeadLetterQueueTypesPtrOutput struct{ *pulumi.OutputState }

func (DeadLetterQueueTypesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DeadLetterQueueTypes)(nil)).Elem()
}

func (o DeadLetterQueueTypesPtrOutput) ToDeadLetterQueueTypesPtrOutput() DeadLetterQueueTypesPtrOutput {
	return o
}

func (o DeadLetterQueueTypesPtrOutput) ToDeadLetterQueueTypesPtrOutputWithContext(ctx context.Context) DeadLetterQueueTypesPtrOutput {
	return o
}

func (o DeadLetterQueueTypesPtrOutput) Elem() DeadLetterQueueTypesOutput {
	return o.ApplyT(func(v *DeadLetterQueueTypes) DeadLetterQueueTypes {
		if v != nil {
			return *v
		}
		var ret DeadLetterQueueTypes
		return ret
	}).(DeadLetterQueueTypesOutput)
}

func (o DeadLetterQueueTypesPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o DeadLetterQueueTypesPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *DeadLetterQueueTypes) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}

// DeadLetterQueueTypesInput is an input type that accepts DeadLetterQueueTypesArgs and DeadLetterQueueTypesOutput values.
// You can construct a concrete instance of `DeadLetterQueueTypesInput` via:
//
//	DeadLetterQueueTypesArgs{...}
type DeadLetterQueueTypesInput interface {
	pulumi.Input

	ToDeadLetterQueueTypesOutput() DeadLetterQueueTypesOutput
	ToDeadLetterQueueTypesOutputWithContext(context.Context) DeadLetterQueueTypesOutput
}

var deadLetterQueueTypesPtrType = reflect.TypeOf((**DeadLetterQueueTypes)(nil)).Elem()

type DeadLetterQueueTypesPtrInput interface {
	pulumi.Input

	ToDeadLetterQueueTypesPtrOutput() DeadLetterQueueTypesPtrOutput
	ToDeadLetterQueueTypesPtrOutputWithContext(context.Context) DeadLetterQueueTypesPtrOutput
}

type deadLetterQueueTypesPtr string

func DeadLetterQueueTypesPtr(v string) DeadLetterQueueTypesPtrInput {
	return (*deadLetterQueueTypesPtr)(&v)
}

func (*deadLetterQueueTypesPtr) ElementType() reflect.Type {
	return deadLetterQueueTypesPtrType
}

func (in *deadLetterQueueTypesPtr) ToDeadLetterQueueTypesPtrOutput() DeadLetterQueueTypesPtrOutput {
	return pulumi.ToOutput(in).(DeadLetterQueueTypesPtrOutput)
}

func (in *deadLetterQueueTypesPtr) ToDeadLetterQueueTypesPtrOutputWithContext(ctx context.Context) DeadLetterQueueTypesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(DeadLetterQueueTypesPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DeadLetterQueueTypesInput)(nil)).Elem(), DeadLetterQueueTypes("permissive"))
	pulumi.RegisterInputType(reflect.TypeOf((*DeadLetterQueueTypesPtrInput)(nil)).Elem(), DeadLetterQueueTypes("permissive"))
	pulumi.RegisterOutputType(DeadLetterQueueTypesOutput{})
	pulumi.RegisterOutputType(DeadLetterQueueTypesPtrOutput{})
}