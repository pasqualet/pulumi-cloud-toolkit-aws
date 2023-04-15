// Code generated by Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package kubernetes

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi-kubernetes/sdk/v3/go/kubernetes"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ClusterAddons is a component that manages the Lubernetes addons to setup a production-ready cluster.
type ClusterAddons struct {
	pulumi.ResourceState

	// The IngressNginx addon used for admin access.
	AdminIngressNginx IngressNginxOutput `pulumi:"adminIngressNginx"`
	// Route53 Zone arn used for admin IngressController.
	AdminZoneArn pulumi.StringPtrOutput `pulumi:"adminZoneArn"`
	// Route53 Zone id used for admin IngressController.
	AdminZoneId pulumi.StringPtrOutput `pulumi:"adminZoneId"`
	// The OpenTelemetry (ADOT) application that sends metrics to CloudWatch.
	AdotApplication AdotApplicationOutput `pulumi:"adotApplication"`
	// The OpenTelemetry (ADOT) operator that sends logs to CloudWatch.
	AdotOperator AdotOperatorOutput `pulumi:"adotOperator"`
	// The ArgoCD addon.
	Argocd ArgoCDOutput `pulumi:"argocd"`
	// The AWS LoadBalancer Controller.
	AwsLoadBalancerController AwsLoadBalancerControllerOutput `pulumi:"awsLoadBalancerController"`
	// The Calico addon used to manage network policies.
	Calico CalicoOutput `pulumi:"calico"`
	// The CertManager addon.
	CertManager CertManagerOutput `pulumi:"certManager"`
	// The Kubernetes ClusterAutoscaler addon.
	ClusterAutoscaler ClusterAutoscalerOutput `pulumi:"clusterAutoscaler"`
	// The Kubernetes dashboard addon.
	Dashboard DashboardOutput `pulumi:"dashboard"`
	// The IngressNginx addon used for default access.
	DefaultIngressNginx IngressNginxOutput `pulumi:"defaultIngressNginx"`
	// Route53 Zone arn used for default IngressController.
	DefaultZoneArn pulumi.StringPtrOutput `pulumi:"defaultZoneArn"`
	// Route53 Zone id used for default IngressController.
	DefaultZoneId pulumi.StringPtrOutput `pulumi:"defaultZoneId"`
	// The EBS CSI driver that allows to create volumes using the block storage service of AWS.
	EbsCsiDriver AwsEbsCsiDriverOutput `pulumi:"ebsCsiDriver"`
	// The ExternalDns addon.
	ExternalDns ExternalDnsOutput `pulumi:"externalDns"`
	// The OpenTelemetry (ADOT) application that sends metrics to CloudWatch.
	Fluentbit FluentbitOutput `pulumi:"fluentbit"`
}

// NewClusterAddons registers a new resource with the given unique name, arguments, and options.
func NewClusterAddons(ctx *pulumi.Context,
	name string, args *ClusterAddonsArgs, opts ...pulumi.ResourceOption) (*ClusterAddons, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ClusterName == nil {
		return nil, errors.New("invalid value for required argument 'ClusterName'")
	}
	if args.IdentityProvidersArn == nil {
		return nil, errors.New("invalid value for required argument 'IdentityProvidersArn'")
	}
	if args.IssuerUrl == nil {
		return nil, errors.New("invalid value for required argument 'IssuerUrl'")
	}
	if args.K8sProvider == nil {
		return nil, errors.New("invalid value for required argument 'K8sProvider'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource ClusterAddons
	err := ctx.RegisterRemoteComponentResource("cloud-toolkit-aws:kubernetes:ClusterAddons", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

type clusterAddonsArgs struct {
	// The EKS Cluster name.
	ClusterName string `pulumi:"clusterName"`
	// The OIDC Identity Provider arn.
	IdentityProvidersArn []string `pulumi:"identityProvidersArn"`
	// The configuration for Ingress Controller.
	Ingress *ClusterAddonsIngress `pulumi:"ingress"`
	// The OIDC Identity Provider url.
	IssuerUrl string `pulumi:"issuerUrl"`
	// The Pulumi provider used for Kubernetes resources.
	K8sProvider *kubernetes.Provider `pulumi:"k8sProvider"`
	// Configure the cluster observability for logging.
	Logging *FluentbitLogging `pulumi:"logging"`
	// Configure the cluster observability for metrics.
	Metrics *AdotApplicationMetrics `pulumi:"metrics"`
}

// The set of arguments for constructing a ClusterAddons resource.
type ClusterAddonsArgs struct {
	// The EKS Cluster name.
	ClusterName pulumi.StringInput
	// The OIDC Identity Provider arn.
	IdentityProvidersArn pulumi.StringArrayInput
	// The configuration for Ingress Controller.
	Ingress ClusterAddonsIngressPtrInput
	// The OIDC Identity Provider url.
	IssuerUrl pulumi.StringInput
	// The Pulumi provider used for Kubernetes resources.
	K8sProvider kubernetes.ProviderInput
	// Configure the cluster observability for logging.
	Logging FluentbitLoggingPtrInput
	// Configure the cluster observability for metrics.
	Metrics AdotApplicationMetricsPtrInput
}

func (ClusterAddonsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*clusterAddonsArgs)(nil)).Elem()
}

type ClusterAddonsInput interface {
	pulumi.Input

	ToClusterAddonsOutput() ClusterAddonsOutput
	ToClusterAddonsOutputWithContext(ctx context.Context) ClusterAddonsOutput
}

func (*ClusterAddons) ElementType() reflect.Type {
	return reflect.TypeOf((**ClusterAddons)(nil)).Elem()
}

func (i *ClusterAddons) ToClusterAddonsOutput() ClusterAddonsOutput {
	return i.ToClusterAddonsOutputWithContext(context.Background())
}

func (i *ClusterAddons) ToClusterAddonsOutputWithContext(ctx context.Context) ClusterAddonsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterAddonsOutput)
}

// ClusterAddonsArrayInput is an input type that accepts ClusterAddonsArray and ClusterAddonsArrayOutput values.
// You can construct a concrete instance of `ClusterAddonsArrayInput` via:
//
//	ClusterAddonsArray{ ClusterAddonsArgs{...} }
type ClusterAddonsArrayInput interface {
	pulumi.Input

	ToClusterAddonsArrayOutput() ClusterAddonsArrayOutput
	ToClusterAddonsArrayOutputWithContext(context.Context) ClusterAddonsArrayOutput
}

type ClusterAddonsArray []ClusterAddonsInput

func (ClusterAddonsArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ClusterAddons)(nil)).Elem()
}

func (i ClusterAddonsArray) ToClusterAddonsArrayOutput() ClusterAddonsArrayOutput {
	return i.ToClusterAddonsArrayOutputWithContext(context.Background())
}

func (i ClusterAddonsArray) ToClusterAddonsArrayOutputWithContext(ctx context.Context) ClusterAddonsArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterAddonsArrayOutput)
}

// ClusterAddonsMapInput is an input type that accepts ClusterAddonsMap and ClusterAddonsMapOutput values.
// You can construct a concrete instance of `ClusterAddonsMapInput` via:
//
//	ClusterAddonsMap{ "key": ClusterAddonsArgs{...} }
type ClusterAddonsMapInput interface {
	pulumi.Input

	ToClusterAddonsMapOutput() ClusterAddonsMapOutput
	ToClusterAddonsMapOutputWithContext(context.Context) ClusterAddonsMapOutput
}

type ClusterAddonsMap map[string]ClusterAddonsInput

func (ClusterAddonsMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ClusterAddons)(nil)).Elem()
}

func (i ClusterAddonsMap) ToClusterAddonsMapOutput() ClusterAddonsMapOutput {
	return i.ToClusterAddonsMapOutputWithContext(context.Background())
}

func (i ClusterAddonsMap) ToClusterAddonsMapOutputWithContext(ctx context.Context) ClusterAddonsMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterAddonsMapOutput)
}

type ClusterAddonsOutput struct{ *pulumi.OutputState }

func (ClusterAddonsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ClusterAddons)(nil)).Elem()
}

func (o ClusterAddonsOutput) ToClusterAddonsOutput() ClusterAddonsOutput {
	return o
}

func (o ClusterAddonsOutput) ToClusterAddonsOutputWithContext(ctx context.Context) ClusterAddonsOutput {
	return o
}

// The IngressNginx addon used for admin access.
func (o ClusterAddonsOutput) AdminIngressNginx() IngressNginxOutput {
	return o.ApplyT(func(v *ClusterAddons) IngressNginxOutput { return v.AdminIngressNginx }).(IngressNginxOutput)
}

// Route53 Zone arn used for admin IngressController.
func (o ClusterAddonsOutput) AdminZoneArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ClusterAddons) pulumi.StringPtrOutput { return v.AdminZoneArn }).(pulumi.StringPtrOutput)
}

// Route53 Zone id used for admin IngressController.
func (o ClusterAddonsOutput) AdminZoneId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ClusterAddons) pulumi.StringPtrOutput { return v.AdminZoneId }).(pulumi.StringPtrOutput)
}

// The OpenTelemetry (ADOT) application that sends metrics to CloudWatch.
func (o ClusterAddonsOutput) AdotApplication() AdotApplicationOutput {
	return o.ApplyT(func(v *ClusterAddons) AdotApplicationOutput { return v.AdotApplication }).(AdotApplicationOutput)
}

// The OpenTelemetry (ADOT) operator that sends logs to CloudWatch.
func (o ClusterAddonsOutput) AdotOperator() AdotOperatorOutput {
	return o.ApplyT(func(v *ClusterAddons) AdotOperatorOutput { return v.AdotOperator }).(AdotOperatorOutput)
}

// The ArgoCD addon.
func (o ClusterAddonsOutput) Argocd() ArgoCDOutput {
	return o.ApplyT(func(v *ClusterAddons) ArgoCDOutput { return v.Argocd }).(ArgoCDOutput)
}

// The AWS LoadBalancer Controller.
func (o ClusterAddonsOutput) AwsLoadBalancerController() AwsLoadBalancerControllerOutput {
	return o.ApplyT(func(v *ClusterAddons) AwsLoadBalancerControllerOutput { return v.AwsLoadBalancerController }).(AwsLoadBalancerControllerOutput)
}

// The Calico addon used to manage network policies.
func (o ClusterAddonsOutput) Calico() CalicoOutput {
	return o.ApplyT(func(v *ClusterAddons) CalicoOutput { return v.Calico }).(CalicoOutput)
}

// The CertManager addon.
func (o ClusterAddonsOutput) CertManager() CertManagerOutput {
	return o.ApplyT(func(v *ClusterAddons) CertManagerOutput { return v.CertManager }).(CertManagerOutput)
}

// The Kubernetes ClusterAutoscaler addon.
func (o ClusterAddonsOutput) ClusterAutoscaler() ClusterAutoscalerOutput {
	return o.ApplyT(func(v *ClusterAddons) ClusterAutoscalerOutput { return v.ClusterAutoscaler }).(ClusterAutoscalerOutput)
}

// The Kubernetes dashboard addon.
func (o ClusterAddonsOutput) Dashboard() DashboardOutput {
	return o.ApplyT(func(v *ClusterAddons) DashboardOutput { return v.Dashboard }).(DashboardOutput)
}

// The IngressNginx addon used for default access.
func (o ClusterAddonsOutput) DefaultIngressNginx() IngressNginxOutput {
	return o.ApplyT(func(v *ClusterAddons) IngressNginxOutput { return v.DefaultIngressNginx }).(IngressNginxOutput)
}

// Route53 Zone arn used for default IngressController.
func (o ClusterAddonsOutput) DefaultZoneArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ClusterAddons) pulumi.StringPtrOutput { return v.DefaultZoneArn }).(pulumi.StringPtrOutput)
}

// Route53 Zone id used for default IngressController.
func (o ClusterAddonsOutput) DefaultZoneId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ClusterAddons) pulumi.StringPtrOutput { return v.DefaultZoneId }).(pulumi.StringPtrOutput)
}

// The EBS CSI driver that allows to create volumes using the block storage service of AWS.
func (o ClusterAddonsOutput) EbsCsiDriver() AwsEbsCsiDriverOutput {
	return o.ApplyT(func(v *ClusterAddons) AwsEbsCsiDriverOutput { return v.EbsCsiDriver }).(AwsEbsCsiDriverOutput)
}

// The ExternalDns addon.
func (o ClusterAddonsOutput) ExternalDns() ExternalDnsOutput {
	return o.ApplyT(func(v *ClusterAddons) ExternalDnsOutput { return v.ExternalDns }).(ExternalDnsOutput)
}

// The OpenTelemetry (ADOT) application that sends metrics to CloudWatch.
func (o ClusterAddonsOutput) Fluentbit() FluentbitOutput {
	return o.ApplyT(func(v *ClusterAddons) FluentbitOutput { return v.Fluentbit }).(FluentbitOutput)
}

type ClusterAddonsArrayOutput struct{ *pulumi.OutputState }

func (ClusterAddonsArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ClusterAddons)(nil)).Elem()
}

func (o ClusterAddonsArrayOutput) ToClusterAddonsArrayOutput() ClusterAddonsArrayOutput {
	return o
}

func (o ClusterAddonsArrayOutput) ToClusterAddonsArrayOutputWithContext(ctx context.Context) ClusterAddonsArrayOutput {
	return o
}

func (o ClusterAddonsArrayOutput) Index(i pulumi.IntInput) ClusterAddonsOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ClusterAddons {
		return vs[0].([]*ClusterAddons)[vs[1].(int)]
	}).(ClusterAddonsOutput)
}

type ClusterAddonsMapOutput struct{ *pulumi.OutputState }

func (ClusterAddonsMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ClusterAddons)(nil)).Elem()
}

func (o ClusterAddonsMapOutput) ToClusterAddonsMapOutput() ClusterAddonsMapOutput {
	return o
}

func (o ClusterAddonsMapOutput) ToClusterAddonsMapOutputWithContext(ctx context.Context) ClusterAddonsMapOutput {
	return o
}

func (o ClusterAddonsMapOutput) MapIndex(k pulumi.StringInput) ClusterAddonsOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ClusterAddons {
		return vs[0].(map[string]*ClusterAddons)[vs[1].(string)]
	}).(ClusterAddonsOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ClusterAddonsInput)(nil)).Elem(), &ClusterAddons{})
	pulumi.RegisterInputType(reflect.TypeOf((*ClusterAddonsArrayInput)(nil)).Elem(), ClusterAddonsArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ClusterAddonsMapInput)(nil)).Elem(), ClusterAddonsMap{})
	pulumi.RegisterOutputType(ClusterAddonsOutput{})
	pulumi.RegisterOutputType(ClusterAddonsArrayOutput{})
	pulumi.RegisterOutputType(ClusterAddonsMapOutput{})
}