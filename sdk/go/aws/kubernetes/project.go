// Code generated by Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package kubernetes

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi-kubernetes/sdk/v3/go/kubernetes"
	corev1 "github.com/pulumi/pulumi-kubernetes/sdk/v3/go/kubernetes/core/v1"
	rbacv1 "github.com/pulumi/pulumi-kubernetes/sdk/v3/go/kubernetes/rbac/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Project is a component that create the resources in the Cluster for a set of AWS IAM Users and Roles, managing the access with the integration with AWS IAM.
type Project struct {
	pulumi.ResourceState

	// The Kubernetes RoleBinding for admin users.
	AdminRoleBinding rbacv1.RoleBindingOutput `pulumi:"adminRoleBinding"`
	// The Kubernetes ClusterRole used to grant minimal access to the cluster.
	ClusterRole rbacv1.ClusterRoleOutput `pulumi:"clusterRole"`
	// The Kubernetes ClusterRoleBinding to associate the ClusterRole to the project.
	ClusterRoleBinding rbacv1.ClusterRoleBindingOutput `pulumi:"clusterRoleBinding"`
	// The Kubernetes RoleBinding for edit users.
	EditRoleBinding rbacv1.RoleBindingOutput `pulumi:"editRoleBinding"`
	// The Namespace used by the project.
	Namespace corev1.NamespaceOutput `pulumi:"namespace"`
	// The Kubernetes provider used to provision Kubernetes resources.
	Provider kubernetes.ProviderOutput `pulumi:"provider"`
	// ResourceQuota for the provisioned Namespace.
	ResourceQuota corev1.ResourceQuotaOutput `pulumi:"resourceQuota"`
	// The Kubernetes RoleBinding for view users.
	ViewRoleBinding rbacv1.RoleBindingOutput `pulumi:"viewRoleBinding"`
}

// NewProject registers a new resource with the given unique name, arguments, and options.
func NewProject(ctx *pulumi.Context,
	name string, args *ProjectArgs, opts ...pulumi.ResourceOption) (*Project, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Kubeconfig == nil {
		return nil, errors.New("invalid value for required argument 'Kubeconfig'")
	}
	if args.Name == nil {
		return nil, errors.New("invalid value for required argument 'Name'")
	}
	if args.Namespace == nil {
		return nil, errors.New("invalid value for required argument 'Namespace'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource Project
	err := ctx.RegisterRemoteComponentResource("cloud-toolkit-aws:kubernetes:Project", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

type projectArgs struct {
	// The list of AWS IAM User arns that can access to this project with 'admin' role.
	AdminUserArns []string `pulumi:"adminUserArns"`
	// The list of AWS IAM User arns that can access to this project with 'edit' role.
	EditUserArns []string `pulumi:"editUserArns"`
	// The kubeconfig to access the kubernetes cluster.
	Kubeconfig string `pulumi:"kubeconfig"`
	// The Project name.
	Name string `pulumi:"name"`
	// The Namespace name where the addon will be installed.
	Namespace string `pulumi:"namespace"`
	// The cluster resources to be assigned to the project.
	Resources *ProjectResources `pulumi:"resources"`
	// The list of AWS IAM User arns that can access to this project with 'view' role.
	ViewUserArns []string `pulumi:"viewUserArns"`
}

// The set of arguments for constructing a Project resource.
type ProjectArgs struct {
	// The list of AWS IAM User arns that can access to this project with 'admin' role.
	AdminUserArns pulumi.StringArrayInput
	// The list of AWS IAM User arns that can access to this project with 'edit' role.
	EditUserArns pulumi.StringArrayInput
	// The kubeconfig to access the kubernetes cluster.
	Kubeconfig pulumi.StringInput
	// The Project name.
	Name pulumi.StringInput
	// The Namespace name where the addon will be installed.
	Namespace pulumi.StringInput
	// The cluster resources to be assigned to the project.
	Resources ProjectResourcesPtrInput
	// The list of AWS IAM User arns that can access to this project with 'view' role.
	ViewUserArns pulumi.StringArrayInput
}

func (ProjectArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*projectArgs)(nil)).Elem()
}

type ProjectInput interface {
	pulumi.Input

	ToProjectOutput() ProjectOutput
	ToProjectOutputWithContext(ctx context.Context) ProjectOutput
}

func (*Project) ElementType() reflect.Type {
	return reflect.TypeOf((**Project)(nil)).Elem()
}

func (i *Project) ToProjectOutput() ProjectOutput {
	return i.ToProjectOutputWithContext(context.Background())
}

func (i *Project) ToProjectOutputWithContext(ctx context.Context) ProjectOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProjectOutput)
}

// ProjectArrayInput is an input type that accepts ProjectArray and ProjectArrayOutput values.
// You can construct a concrete instance of `ProjectArrayInput` via:
//
//	ProjectArray{ ProjectArgs{...} }
type ProjectArrayInput interface {
	pulumi.Input

	ToProjectArrayOutput() ProjectArrayOutput
	ToProjectArrayOutputWithContext(context.Context) ProjectArrayOutput
}

type ProjectArray []ProjectInput

func (ProjectArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Project)(nil)).Elem()
}

func (i ProjectArray) ToProjectArrayOutput() ProjectArrayOutput {
	return i.ToProjectArrayOutputWithContext(context.Background())
}

func (i ProjectArray) ToProjectArrayOutputWithContext(ctx context.Context) ProjectArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProjectArrayOutput)
}

// ProjectMapInput is an input type that accepts ProjectMap and ProjectMapOutput values.
// You can construct a concrete instance of `ProjectMapInput` via:
//
//	ProjectMap{ "key": ProjectArgs{...} }
type ProjectMapInput interface {
	pulumi.Input

	ToProjectMapOutput() ProjectMapOutput
	ToProjectMapOutputWithContext(context.Context) ProjectMapOutput
}

type ProjectMap map[string]ProjectInput

func (ProjectMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Project)(nil)).Elem()
}

func (i ProjectMap) ToProjectMapOutput() ProjectMapOutput {
	return i.ToProjectMapOutputWithContext(context.Background())
}

func (i ProjectMap) ToProjectMapOutputWithContext(ctx context.Context) ProjectMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProjectMapOutput)
}

type ProjectOutput struct{ *pulumi.OutputState }

func (ProjectOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Project)(nil)).Elem()
}

func (o ProjectOutput) ToProjectOutput() ProjectOutput {
	return o
}

func (o ProjectOutput) ToProjectOutputWithContext(ctx context.Context) ProjectOutput {
	return o
}

// The Kubernetes RoleBinding for admin users.
func (o ProjectOutput) AdminRoleBinding() rbacv1.RoleBindingOutput {
	return o.ApplyT(func(v *Project) rbacv1.RoleBindingOutput { return v.AdminRoleBinding }).(rbacv1.RoleBindingOutput)
}

// The Kubernetes ClusterRole used to grant minimal access to the cluster.
func (o ProjectOutput) ClusterRole() rbacv1.ClusterRoleOutput {
	return o.ApplyT(func(v *Project) rbacv1.ClusterRoleOutput { return v.ClusterRole }).(rbacv1.ClusterRoleOutput)
}

// The Kubernetes ClusterRoleBinding to associate the ClusterRole to the project.
func (o ProjectOutput) ClusterRoleBinding() rbacv1.ClusterRoleBindingOutput {
	return o.ApplyT(func(v *Project) rbacv1.ClusterRoleBindingOutput { return v.ClusterRoleBinding }).(rbacv1.ClusterRoleBindingOutput)
}

// The Kubernetes RoleBinding for edit users.
func (o ProjectOutput) EditRoleBinding() rbacv1.RoleBindingOutput {
	return o.ApplyT(func(v *Project) rbacv1.RoleBindingOutput { return v.EditRoleBinding }).(rbacv1.RoleBindingOutput)
}

// The Namespace used by the project.
func (o ProjectOutput) Namespace() corev1.NamespaceOutput {
	return o.ApplyT(func(v *Project) corev1.NamespaceOutput { return v.Namespace }).(corev1.NamespaceOutput)
}

// The Kubernetes provider used to provision Kubernetes resources.
func (o ProjectOutput) Provider() kubernetes.ProviderOutput {
	return o.ApplyT(func(v *Project) kubernetes.ProviderOutput { return v.Provider }).(kubernetes.ProviderOutput)
}

// ResourceQuota for the provisioned Namespace.
func (o ProjectOutput) ResourceQuota() corev1.ResourceQuotaOutput {
	return o.ApplyT(func(v *Project) corev1.ResourceQuotaOutput { return v.ResourceQuota }).(corev1.ResourceQuotaOutput)
}

// The Kubernetes RoleBinding for view users.
func (o ProjectOutput) ViewRoleBinding() rbacv1.RoleBindingOutput {
	return o.ApplyT(func(v *Project) rbacv1.RoleBindingOutput { return v.ViewRoleBinding }).(rbacv1.RoleBindingOutput)
}

type ProjectArrayOutput struct{ *pulumi.OutputState }

func (ProjectArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Project)(nil)).Elem()
}

func (o ProjectArrayOutput) ToProjectArrayOutput() ProjectArrayOutput {
	return o
}

func (o ProjectArrayOutput) ToProjectArrayOutputWithContext(ctx context.Context) ProjectArrayOutput {
	return o
}

func (o ProjectArrayOutput) Index(i pulumi.IntInput) ProjectOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Project {
		return vs[0].([]*Project)[vs[1].(int)]
	}).(ProjectOutput)
}

type ProjectMapOutput struct{ *pulumi.OutputState }

func (ProjectMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Project)(nil)).Elem()
}

func (o ProjectMapOutput) ToProjectMapOutput() ProjectMapOutput {
	return o
}

func (o ProjectMapOutput) ToProjectMapOutputWithContext(ctx context.Context) ProjectMapOutput {
	return o
}

func (o ProjectMapOutput) MapIndex(k pulumi.StringInput) ProjectOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Project {
		return vs[0].(map[string]*Project)[vs[1].(string)]
	}).(ProjectOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ProjectInput)(nil)).Elem(), &Project{})
	pulumi.RegisterInputType(reflect.TypeOf((*ProjectArrayInput)(nil)).Elem(), ProjectArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ProjectMapInput)(nil)).Elem(), ProjectMap{})
	pulumi.RegisterOutputType(ProjectOutput{})
	pulumi.RegisterOutputType(ProjectArrayOutput{})
	pulumi.RegisterOutputType(ProjectMapOutput{})
}
