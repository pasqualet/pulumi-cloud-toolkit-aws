// Code generated by Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package kubernetes

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi-kubernetes/sdk/v3/go/kubernetes"
	rbacv1 "github.com/pulumi/pulumi-kubernetes/sdk/v3/go/kubernetes/rbac/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ClusterAdmins is a component that create the resources in the Cluster for a set of AWS IAM Users and Roles, managing the access with the integration with AWS IAM.
type ClusterAdmins struct {
	pulumi.ResourceState

	// The Kubernetes ClusterRoleBinding to associate the ClusterRole to the project.
	ClusterRoleBinding rbacv1.ClusterRoleBindingOutput `pulumi:"clusterRoleBinding"`
	// The Kubernetes provider used to provision Kubernetes resources.
	Provider kubernetes.ProviderOutput `pulumi:"provider"`
}

// NewClusterAdmins registers a new resource with the given unique name, arguments, and options.
func NewClusterAdmins(ctx *pulumi.Context,
	name string, args *ClusterAdminsArgs, opts ...pulumi.ResourceOption) (*ClusterAdmins, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Kubeconfig == nil {
		return nil, errors.New("invalid value for required argument 'Kubeconfig'")
	}
	if args.Name == nil {
		return nil, errors.New("invalid value for required argument 'Name'")
	}
	if args.UserArns == nil {
		return nil, errors.New("invalid value for required argument 'UserArns'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource ClusterAdmins
	err := ctx.RegisterRemoteComponentResource("cloud-toolkit-aws:kubernetes:ClusterAdmins", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

type clusterAdminsArgs struct {
	// Kubernetes provider used by Pulumi.
	Kubeconfig string `pulumi:"kubeconfig"`
	// The name for the group of Cluster Admins.
	Name string `pulumi:"name"`
	// The list of AWS IAM User arns.
	UserArns []string `pulumi:"userArns"`
}

// The set of arguments for constructing a ClusterAdmins resource.
type ClusterAdminsArgs struct {
	// Kubernetes provider used by Pulumi.
	Kubeconfig pulumi.StringInput
	// The name for the group of Cluster Admins.
	Name pulumi.StringInput
	// The list of AWS IAM User arns.
	UserArns pulumi.StringArrayInput
}

func (ClusterAdminsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*clusterAdminsArgs)(nil)).Elem()
}

type ClusterAdminsInput interface {
	pulumi.Input

	ToClusterAdminsOutput() ClusterAdminsOutput
	ToClusterAdminsOutputWithContext(ctx context.Context) ClusterAdminsOutput
}

func (*ClusterAdmins) ElementType() reflect.Type {
	return reflect.TypeOf((**ClusterAdmins)(nil)).Elem()
}

func (i *ClusterAdmins) ToClusterAdminsOutput() ClusterAdminsOutput {
	return i.ToClusterAdminsOutputWithContext(context.Background())
}

func (i *ClusterAdmins) ToClusterAdminsOutputWithContext(ctx context.Context) ClusterAdminsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterAdminsOutput)
}

// ClusterAdminsArrayInput is an input type that accepts ClusterAdminsArray and ClusterAdminsArrayOutput values.
// You can construct a concrete instance of `ClusterAdminsArrayInput` via:
//
//	ClusterAdminsArray{ ClusterAdminsArgs{...} }
type ClusterAdminsArrayInput interface {
	pulumi.Input

	ToClusterAdminsArrayOutput() ClusterAdminsArrayOutput
	ToClusterAdminsArrayOutputWithContext(context.Context) ClusterAdminsArrayOutput
}

type ClusterAdminsArray []ClusterAdminsInput

func (ClusterAdminsArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ClusterAdmins)(nil)).Elem()
}

func (i ClusterAdminsArray) ToClusterAdminsArrayOutput() ClusterAdminsArrayOutput {
	return i.ToClusterAdminsArrayOutputWithContext(context.Background())
}

func (i ClusterAdminsArray) ToClusterAdminsArrayOutputWithContext(ctx context.Context) ClusterAdminsArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterAdminsArrayOutput)
}

// ClusterAdminsMapInput is an input type that accepts ClusterAdminsMap and ClusterAdminsMapOutput values.
// You can construct a concrete instance of `ClusterAdminsMapInput` via:
//
//	ClusterAdminsMap{ "key": ClusterAdminsArgs{...} }
type ClusterAdminsMapInput interface {
	pulumi.Input

	ToClusterAdminsMapOutput() ClusterAdminsMapOutput
	ToClusterAdminsMapOutputWithContext(context.Context) ClusterAdminsMapOutput
}

type ClusterAdminsMap map[string]ClusterAdminsInput

func (ClusterAdminsMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ClusterAdmins)(nil)).Elem()
}

func (i ClusterAdminsMap) ToClusterAdminsMapOutput() ClusterAdminsMapOutput {
	return i.ToClusterAdminsMapOutputWithContext(context.Background())
}

func (i ClusterAdminsMap) ToClusterAdminsMapOutputWithContext(ctx context.Context) ClusterAdminsMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterAdminsMapOutput)
}

type ClusterAdminsOutput struct{ *pulumi.OutputState }

func (ClusterAdminsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ClusterAdmins)(nil)).Elem()
}

func (o ClusterAdminsOutput) ToClusterAdminsOutput() ClusterAdminsOutput {
	return o
}

func (o ClusterAdminsOutput) ToClusterAdminsOutputWithContext(ctx context.Context) ClusterAdminsOutput {
	return o
}

// The Kubernetes ClusterRoleBinding to associate the ClusterRole to the project.
func (o ClusterAdminsOutput) ClusterRoleBinding() rbacv1.ClusterRoleBindingOutput {
	return o.ApplyT(func(v *ClusterAdmins) rbacv1.ClusterRoleBindingOutput { return v.ClusterRoleBinding }).(rbacv1.ClusterRoleBindingOutput)
}

// The Kubernetes provider used to provision Kubernetes resources.
func (o ClusterAdminsOutput) Provider() kubernetes.ProviderOutput {
	return o.ApplyT(func(v *ClusterAdmins) kubernetes.ProviderOutput { return v.Provider }).(kubernetes.ProviderOutput)
}

type ClusterAdminsArrayOutput struct{ *pulumi.OutputState }

func (ClusterAdminsArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ClusterAdmins)(nil)).Elem()
}

func (o ClusterAdminsArrayOutput) ToClusterAdminsArrayOutput() ClusterAdminsArrayOutput {
	return o
}

func (o ClusterAdminsArrayOutput) ToClusterAdminsArrayOutputWithContext(ctx context.Context) ClusterAdminsArrayOutput {
	return o
}

func (o ClusterAdminsArrayOutput) Index(i pulumi.IntInput) ClusterAdminsOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ClusterAdmins {
		return vs[0].([]*ClusterAdmins)[vs[1].(int)]
	}).(ClusterAdminsOutput)
}

type ClusterAdminsMapOutput struct{ *pulumi.OutputState }

func (ClusterAdminsMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ClusterAdmins)(nil)).Elem()
}

func (o ClusterAdminsMapOutput) ToClusterAdminsMapOutput() ClusterAdminsMapOutput {
	return o
}

func (o ClusterAdminsMapOutput) ToClusterAdminsMapOutputWithContext(ctx context.Context) ClusterAdminsMapOutput {
	return o
}

func (o ClusterAdminsMapOutput) MapIndex(k pulumi.StringInput) ClusterAdminsOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ClusterAdmins {
		return vs[0].(map[string]*ClusterAdmins)[vs[1].(string)]
	}).(ClusterAdminsOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ClusterAdminsInput)(nil)).Elem(), &ClusterAdmins{})
	pulumi.RegisterInputType(reflect.TypeOf((*ClusterAdminsArrayInput)(nil)).Elem(), ClusterAdminsArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ClusterAdminsMapInput)(nil)).Elem(), ClusterAdminsMap{})
	pulumi.RegisterOutputType(ClusterAdminsOutput{})
	pulumi.RegisterOutputType(ClusterAdminsArrayOutput{})
	pulumi.RegisterOutputType(ClusterAdminsMapOutput{})
}