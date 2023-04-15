// Code generated by Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package databases

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/cloudwatch"
	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/ec2"
	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/rds"
	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/secretsmanager"
	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/sns"
	"github.com/pulumi/pulumi-random/sdk/v4/go/random"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AuroraMysql struct {
	pulumi.ResourceState

	// CloudWatch alarms that monitor the RDS Cluter
	Alarms cloudwatch.MetricAlarmArrayOutput `pulumi:"alarms"`
	// The RDS Cluster
	Cluster rds.ClusterOutput `pulumi:"cluster"`
	// Cluster instances associated to the cluster
	Instances rds.ClusterInstanceArrayOutput `pulumi:"instances"`
	// Random password generated for admin user
	Password random.RandomPasswordOutput `pulumi:"password"`
	// Component that protects and stores admin password in AWS
	Secret secretsmanager.SecretOutput `pulumi:"secret"`
	// Component that protects and stores admin password in AWS
	SecretVersion secretsmanager.SecretVersionOutput `pulumi:"secretVersion"`
	// The SecurityGroup associated to the cluster to manage traffic
	SecurityGroup ec2.SecurityGroupOutput `pulumi:"securityGroup"`
	// The rules associated SecurityGroup to allow incoming traffic
	SecurityGroupRule ec2.SecurityGroupRuleOutput `pulumi:"securityGroupRule"`
	// The SubnetGroup that reprents the list of subnets where the cluster is deployed
	SubnetGroup rds.SubnetGroupOutput `pulumi:"subnetGroup"`
	// SNS Topic used for CloudWatch alerts
	Topic sns.TopicOutput `pulumi:"topic"`
	// TopicSubscriptions to alerts by email
	TopicSubscriptions sns.TopicSubscriptionArrayOutput `pulumi:"topicSubscriptions"`
}

// NewAuroraMysql registers a new resource with the given unique name, arguments, and options.
func NewAuroraMysql(ctx *pulumi.Context,
	name string, args *AuroraMysqlArgs, opts ...pulumi.ResourceOption) (*AuroraMysql, error) {
	if args == nil {
		args = &AuroraMysqlArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource AuroraMysql
	err := ctx.RegisterRemoteComponentResource("cloud-toolkit-aws:databases:AuroraMysql", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

type auroraMysqlArgs struct {
	// Backup configuration parameters for Aurora cluster
	Backup *AuroraMysqlBackup `pulumi:"backup"`
	// Configuration parameters for the database
	Database *AuroraMysqlDatabase `pulumi:"database"`
	// The instance type for the cluster
	InstanceType *string `pulumi:"instanceType"`
	// The number of instances to be created for Aurora cluster
	InstancesCount *float64 `pulumi:"instancesCount"`
	// Logging configuration parameters for Aurora cluster
	Logging *AuroraMysqlLogging `pulumi:"logging"`
	// Monitoring configuration parameters for Aurora cluster
	Monitoring *AuroraMysqlMonitoring `pulumi:"monitoring"`
	// Networking configuration parameters for Aurora cluster
	Networking *AuroraMysqlNetworking `pulumi:"networking"`
	// Version for database
	Version *string `pulumi:"version"`
}

// The set of arguments for constructing a AuroraMysql resource.
type AuroraMysqlArgs struct {
	// Backup configuration parameters for Aurora cluster
	Backup AuroraMysqlBackupPtrInput
	// Configuration parameters for the database
	Database AuroraMysqlDatabasePtrInput
	// The instance type for the cluster
	InstanceType pulumi.StringPtrInput
	// The number of instances to be created for Aurora cluster
	InstancesCount pulumi.Float64PtrInput
	// Logging configuration parameters for Aurora cluster
	Logging AuroraMysqlLoggingPtrInput
	// Monitoring configuration parameters for Aurora cluster
	Monitoring AuroraMysqlMonitoringPtrInput
	// Networking configuration parameters for Aurora cluster
	Networking AuroraMysqlNetworkingPtrInput
	// Version for database
	Version pulumi.StringPtrInput
}

func (AuroraMysqlArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*auroraMysqlArgs)(nil)).Elem()
}

type AuroraMysqlInput interface {
	pulumi.Input

	ToAuroraMysqlOutput() AuroraMysqlOutput
	ToAuroraMysqlOutputWithContext(ctx context.Context) AuroraMysqlOutput
}

func (*AuroraMysql) ElementType() reflect.Type {
	return reflect.TypeOf((**AuroraMysql)(nil)).Elem()
}

func (i *AuroraMysql) ToAuroraMysqlOutput() AuroraMysqlOutput {
	return i.ToAuroraMysqlOutputWithContext(context.Background())
}

func (i *AuroraMysql) ToAuroraMysqlOutputWithContext(ctx context.Context) AuroraMysqlOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AuroraMysqlOutput)
}

// AuroraMysqlArrayInput is an input type that accepts AuroraMysqlArray and AuroraMysqlArrayOutput values.
// You can construct a concrete instance of `AuroraMysqlArrayInput` via:
//
//	AuroraMysqlArray{ AuroraMysqlArgs{...} }
type AuroraMysqlArrayInput interface {
	pulumi.Input

	ToAuroraMysqlArrayOutput() AuroraMysqlArrayOutput
	ToAuroraMysqlArrayOutputWithContext(context.Context) AuroraMysqlArrayOutput
}

type AuroraMysqlArray []AuroraMysqlInput

func (AuroraMysqlArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AuroraMysql)(nil)).Elem()
}

func (i AuroraMysqlArray) ToAuroraMysqlArrayOutput() AuroraMysqlArrayOutput {
	return i.ToAuroraMysqlArrayOutputWithContext(context.Background())
}

func (i AuroraMysqlArray) ToAuroraMysqlArrayOutputWithContext(ctx context.Context) AuroraMysqlArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AuroraMysqlArrayOutput)
}

// AuroraMysqlMapInput is an input type that accepts AuroraMysqlMap and AuroraMysqlMapOutput values.
// You can construct a concrete instance of `AuroraMysqlMapInput` via:
//
//	AuroraMysqlMap{ "key": AuroraMysqlArgs{...} }
type AuroraMysqlMapInput interface {
	pulumi.Input

	ToAuroraMysqlMapOutput() AuroraMysqlMapOutput
	ToAuroraMysqlMapOutputWithContext(context.Context) AuroraMysqlMapOutput
}

type AuroraMysqlMap map[string]AuroraMysqlInput

func (AuroraMysqlMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AuroraMysql)(nil)).Elem()
}

func (i AuroraMysqlMap) ToAuroraMysqlMapOutput() AuroraMysqlMapOutput {
	return i.ToAuroraMysqlMapOutputWithContext(context.Background())
}

func (i AuroraMysqlMap) ToAuroraMysqlMapOutputWithContext(ctx context.Context) AuroraMysqlMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AuroraMysqlMapOutput)
}

type AuroraMysqlOutput struct{ *pulumi.OutputState }

func (AuroraMysqlOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AuroraMysql)(nil)).Elem()
}

func (o AuroraMysqlOutput) ToAuroraMysqlOutput() AuroraMysqlOutput {
	return o
}

func (o AuroraMysqlOutput) ToAuroraMysqlOutputWithContext(ctx context.Context) AuroraMysqlOutput {
	return o
}

// CloudWatch alarms that monitor the RDS Cluter
func (o AuroraMysqlOutput) Alarms() cloudwatch.MetricAlarmArrayOutput {
	return o.ApplyT(func(v *AuroraMysql) cloudwatch.MetricAlarmArrayOutput { return v.Alarms }).(cloudwatch.MetricAlarmArrayOutput)
}

// The RDS Cluster
func (o AuroraMysqlOutput) Cluster() rds.ClusterOutput {
	return o.ApplyT(func(v *AuroraMysql) rds.ClusterOutput { return v.Cluster }).(rds.ClusterOutput)
}

// Cluster instances associated to the cluster
func (o AuroraMysqlOutput) Instances() rds.ClusterInstanceArrayOutput {
	return o.ApplyT(func(v *AuroraMysql) rds.ClusterInstanceArrayOutput { return v.Instances }).(rds.ClusterInstanceArrayOutput)
}

// Random password generated for admin user
func (o AuroraMysqlOutput) Password() random.RandomPasswordOutput {
	return o.ApplyT(func(v *AuroraMysql) random.RandomPasswordOutput { return v.Password }).(random.RandomPasswordOutput)
}

// Component that protects and stores admin password in AWS
func (o AuroraMysqlOutput) Secret() secretsmanager.SecretOutput {
	return o.ApplyT(func(v *AuroraMysql) secretsmanager.SecretOutput { return v.Secret }).(secretsmanager.SecretOutput)
}

// Component that protects and stores admin password in AWS
func (o AuroraMysqlOutput) SecretVersion() secretsmanager.SecretVersionOutput {
	return o.ApplyT(func(v *AuroraMysql) secretsmanager.SecretVersionOutput { return v.SecretVersion }).(secretsmanager.SecretVersionOutput)
}

// The SecurityGroup associated to the cluster to manage traffic
func (o AuroraMysqlOutput) SecurityGroup() ec2.SecurityGroupOutput {
	return o.ApplyT(func(v *AuroraMysql) ec2.SecurityGroupOutput { return v.SecurityGroup }).(ec2.SecurityGroupOutput)
}

// The rules associated SecurityGroup to allow incoming traffic
func (o AuroraMysqlOutput) SecurityGroupRule() ec2.SecurityGroupRuleOutput {
	return o.ApplyT(func(v *AuroraMysql) ec2.SecurityGroupRuleOutput { return v.SecurityGroupRule }).(ec2.SecurityGroupRuleOutput)
}

// The SubnetGroup that reprents the list of subnets where the cluster is deployed
func (o AuroraMysqlOutput) SubnetGroup() rds.SubnetGroupOutput {
	return o.ApplyT(func(v *AuroraMysql) rds.SubnetGroupOutput { return v.SubnetGroup }).(rds.SubnetGroupOutput)
}

// SNS Topic used for CloudWatch alerts
func (o AuroraMysqlOutput) Topic() sns.TopicOutput {
	return o.ApplyT(func(v *AuroraMysql) sns.TopicOutput { return v.Topic }).(sns.TopicOutput)
}

// TopicSubscriptions to alerts by email
func (o AuroraMysqlOutput) TopicSubscriptions() sns.TopicSubscriptionArrayOutput {
	return o.ApplyT(func(v *AuroraMysql) sns.TopicSubscriptionArrayOutput { return v.TopicSubscriptions }).(sns.TopicSubscriptionArrayOutput)
}

type AuroraMysqlArrayOutput struct{ *pulumi.OutputState }

func (AuroraMysqlArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AuroraMysql)(nil)).Elem()
}

func (o AuroraMysqlArrayOutput) ToAuroraMysqlArrayOutput() AuroraMysqlArrayOutput {
	return o
}

func (o AuroraMysqlArrayOutput) ToAuroraMysqlArrayOutputWithContext(ctx context.Context) AuroraMysqlArrayOutput {
	return o
}

func (o AuroraMysqlArrayOutput) Index(i pulumi.IntInput) AuroraMysqlOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *AuroraMysql {
		return vs[0].([]*AuroraMysql)[vs[1].(int)]
	}).(AuroraMysqlOutput)
}

type AuroraMysqlMapOutput struct{ *pulumi.OutputState }

func (AuroraMysqlMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AuroraMysql)(nil)).Elem()
}

func (o AuroraMysqlMapOutput) ToAuroraMysqlMapOutput() AuroraMysqlMapOutput {
	return o
}

func (o AuroraMysqlMapOutput) ToAuroraMysqlMapOutputWithContext(ctx context.Context) AuroraMysqlMapOutput {
	return o
}

func (o AuroraMysqlMapOutput) MapIndex(k pulumi.StringInput) AuroraMysqlOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *AuroraMysql {
		return vs[0].(map[string]*AuroraMysql)[vs[1].(string)]
	}).(AuroraMysqlOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AuroraMysqlInput)(nil)).Elem(), &AuroraMysql{})
	pulumi.RegisterInputType(reflect.TypeOf((*AuroraMysqlArrayInput)(nil)).Elem(), AuroraMysqlArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AuroraMysqlMapInput)(nil)).Elem(), AuroraMysqlMap{})
	pulumi.RegisterOutputType(AuroraMysqlOutput{})
	pulumi.RegisterOutputType(AuroraMysqlArrayOutput{})
	pulumi.RegisterOutputType(AuroraMysqlMapOutput{})
}