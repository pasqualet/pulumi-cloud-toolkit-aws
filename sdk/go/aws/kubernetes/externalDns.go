// Code generated by Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package kubernetes

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi-kubernetes/sdk/v3/go/kubernetes/apiextensions"
	corev1 "github.com/pulumi/pulumi-kubernetes/sdk/v3/go/kubernetes/core/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ExternalDns is a component to manage DNS records according to the Ingresses created in the cluster.
type ExternalDns struct {
	pulumi.ResourceState

	// The Namespace used to deploy the component.
	Application apiextensions.CustomResourceOutput `pulumi:"application"`
	// The IAM roles for service accounts.
	Irsa IrsaOutput `pulumi:"irsa"`
	// The Namespace used to deploy the component.
	Namespace corev1.NamespaceOutput `pulumi:"namespace"`
}

// NewExternalDns registers a new resource with the given unique name, arguments, and options.
func NewExternalDns(ctx *pulumi.Context,
	name string, args *ExternalDnsArgs, opts ...pulumi.ResourceOption) (*ExternalDns, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ZoneArns == nil {
		return nil, errors.New("invalid value for required argument 'ZoneArns'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource ExternalDns
	err := ctx.RegisterRemoteComponentResource("cloud-toolkit-aws:kubernetes:ExternalDns", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

type externalDnsArgs struct {
	// The list of DNS Zone arn to be used by ExternalDns.
	ZoneArns []string `pulumi:"zoneArns"`
}

// The set of arguments for constructing a ExternalDns resource.
type ExternalDnsArgs struct {
	// The list of DNS Zone arn to be used by ExternalDns.
	ZoneArns pulumi.StringArrayInput
}

func (ExternalDnsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*externalDnsArgs)(nil)).Elem()
}

type ExternalDnsInput interface {
	pulumi.Input

	ToExternalDnsOutput() ExternalDnsOutput
	ToExternalDnsOutputWithContext(ctx context.Context) ExternalDnsOutput
}

func (*ExternalDns) ElementType() reflect.Type {
	return reflect.TypeOf((**ExternalDns)(nil)).Elem()
}

func (i *ExternalDns) ToExternalDnsOutput() ExternalDnsOutput {
	return i.ToExternalDnsOutputWithContext(context.Background())
}

func (i *ExternalDns) ToExternalDnsOutputWithContext(ctx context.Context) ExternalDnsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExternalDnsOutput)
}

// ExternalDnsArrayInput is an input type that accepts ExternalDnsArray and ExternalDnsArrayOutput values.
// You can construct a concrete instance of `ExternalDnsArrayInput` via:
//
//	ExternalDnsArray{ ExternalDnsArgs{...} }
type ExternalDnsArrayInput interface {
	pulumi.Input

	ToExternalDnsArrayOutput() ExternalDnsArrayOutput
	ToExternalDnsArrayOutputWithContext(context.Context) ExternalDnsArrayOutput
}

type ExternalDnsArray []ExternalDnsInput

func (ExternalDnsArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ExternalDns)(nil)).Elem()
}

func (i ExternalDnsArray) ToExternalDnsArrayOutput() ExternalDnsArrayOutput {
	return i.ToExternalDnsArrayOutputWithContext(context.Background())
}

func (i ExternalDnsArray) ToExternalDnsArrayOutputWithContext(ctx context.Context) ExternalDnsArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExternalDnsArrayOutput)
}

// ExternalDnsMapInput is an input type that accepts ExternalDnsMap and ExternalDnsMapOutput values.
// You can construct a concrete instance of `ExternalDnsMapInput` via:
//
//	ExternalDnsMap{ "key": ExternalDnsArgs{...} }
type ExternalDnsMapInput interface {
	pulumi.Input

	ToExternalDnsMapOutput() ExternalDnsMapOutput
	ToExternalDnsMapOutputWithContext(context.Context) ExternalDnsMapOutput
}

type ExternalDnsMap map[string]ExternalDnsInput

func (ExternalDnsMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ExternalDns)(nil)).Elem()
}

func (i ExternalDnsMap) ToExternalDnsMapOutput() ExternalDnsMapOutput {
	return i.ToExternalDnsMapOutputWithContext(context.Background())
}

func (i ExternalDnsMap) ToExternalDnsMapOutputWithContext(ctx context.Context) ExternalDnsMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExternalDnsMapOutput)
}

type ExternalDnsOutput struct{ *pulumi.OutputState }

func (ExternalDnsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ExternalDns)(nil)).Elem()
}

func (o ExternalDnsOutput) ToExternalDnsOutput() ExternalDnsOutput {
	return o
}

func (o ExternalDnsOutput) ToExternalDnsOutputWithContext(ctx context.Context) ExternalDnsOutput {
	return o
}

// The Namespace used to deploy the component.
func (o ExternalDnsOutput) Application() apiextensions.CustomResourceOutput {
	return o.ApplyT(func(v *ExternalDns) apiextensions.CustomResourceOutput { return v.Application }).(apiextensions.CustomResourceOutput)
}

// The IAM roles for service accounts.
func (o ExternalDnsOutput) Irsa() IrsaOutput {
	return o.ApplyT(func(v *ExternalDns) IrsaOutput { return v.Irsa }).(IrsaOutput)
}

// The Namespace used to deploy the component.
func (o ExternalDnsOutput) Namespace() corev1.NamespaceOutput {
	return o.ApplyT(func(v *ExternalDns) corev1.NamespaceOutput { return v.Namespace }).(corev1.NamespaceOutput)
}

type ExternalDnsArrayOutput struct{ *pulumi.OutputState }

func (ExternalDnsArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ExternalDns)(nil)).Elem()
}

func (o ExternalDnsArrayOutput) ToExternalDnsArrayOutput() ExternalDnsArrayOutput {
	return o
}

func (o ExternalDnsArrayOutput) ToExternalDnsArrayOutputWithContext(ctx context.Context) ExternalDnsArrayOutput {
	return o
}

func (o ExternalDnsArrayOutput) Index(i pulumi.IntInput) ExternalDnsOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ExternalDns {
		return vs[0].([]*ExternalDns)[vs[1].(int)]
	}).(ExternalDnsOutput)
}

type ExternalDnsMapOutput struct{ *pulumi.OutputState }

func (ExternalDnsMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ExternalDns)(nil)).Elem()
}

func (o ExternalDnsMapOutput) ToExternalDnsMapOutput() ExternalDnsMapOutput {
	return o
}

func (o ExternalDnsMapOutput) ToExternalDnsMapOutputWithContext(ctx context.Context) ExternalDnsMapOutput {
	return o
}

func (o ExternalDnsMapOutput) MapIndex(k pulumi.StringInput) ExternalDnsOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ExternalDns {
		return vs[0].(map[string]*ExternalDns)[vs[1].(string)]
	}).(ExternalDnsOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ExternalDnsInput)(nil)).Elem(), &ExternalDns{})
	pulumi.RegisterInputType(reflect.TypeOf((*ExternalDnsArrayInput)(nil)).Elem(), ExternalDnsArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ExternalDnsMapInput)(nil)).Elem(), ExternalDnsMap{})
	pulumi.RegisterOutputType(ExternalDnsOutput{})
	pulumi.RegisterOutputType(ExternalDnsArrayOutput{})
	pulumi.RegisterOutputType(ExternalDnsMapOutput{})
}