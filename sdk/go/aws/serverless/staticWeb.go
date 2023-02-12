// Code generated by Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package serverless

import (
	"context"
	"reflect"

	"github.com/cloud-toolkit/cloud-toolkit-aws/sdk/go/aws/storage"
	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws"
	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/acm"
	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/cloudfront"
	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/s3"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type StaticWeb struct {
	pulumi.ResourceState

	// DNS Records to expose staticweb
	DNSRecords DNSRecordsPtrOutput `pulumi:"DNSRecords"`
	// DNS Records to validate the certificate
	DNSRecordsForValidation DNSRecordsPtrOutput `pulumi:"DNSRecordsForValidation"`
	// CloudFront Distribution
	Certificate acm.CertificateOutput `pulumi:"certificate"`
	// AWS certificate validation
	CertificateValidation acm.CertificateValidationOutput `pulumi:"certificateValidation"`
	// Content bucket
	ContentBucket storage.BucketOutput `pulumi:"contentBucket"`
	// Bucket policy to connect Cloud Front distribution
	ContentBucketPolicy s3.BucketPolicyOutput `pulumi:"contentBucketPolicy"`
	// CloudFront Distribution
	Distribution cloudfront.DistributionOutput `pulumi:"distribution"`
	// CloudFront Distribution
	DomainValidationOptions acm.CertificateDomainValidationOptionArrayOutput `pulumi:"domainValidationOptions"`
	// AWS eastRegion provider. It is required to create and validate certificates
	EastRegion aws.ProviderOutput `pulumi:"eastRegion"`
	// Logs bucket
	LogsBucket s3.BucketOutput `pulumi:"logsBucket"`
	// Staticweb name
	Name pulumi.StringOutput `pulumi:"name"`
	// OriginAccessIdentity to have access to content bucket
	OriginAccessIdentity cloudfront.OriginAccessIdentityOutput `pulumi:"originAccessIdentity"`
}

// NewStaticWeb registers a new resource with the given unique name, arguments, and options.
func NewStaticWeb(ctx *pulumi.Context,
	name string, args *StaticWebArgs, opts ...pulumi.ResourceOption) (*StaticWeb, error) {
	if args == nil {
		args = &StaticWebArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource StaticWeb
	err := ctx.RegisterRemoteComponentResource("cloud-toolkit-aws:serverless:StaticWeb", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

type staticWebArgs struct {
	// Cloud Front distribution cache
	Cache *CdnCache `pulumi:"cache"`
	// Set to true to configure DNS
	ConfigureDNS *bool `pulumi:"configureDNS"`
	// DNS configuration
	Dns *CdnDns `pulumi:"dns"`
	// Set to true to add an alias to wwww.<domain>
	Domain *string `pulumi:"domain"`
	// Set to true to add an alias to wwww.<domain>
	IncludeWWW *bool `pulumi:"includeWWW"`
	// Cloud Front distribution priceClass
	PriceClass *string `pulumi:"priceClass"`
}

// The set of arguments for constructing a StaticWeb resource.
type StaticWebArgs struct {
	// Cloud Front distribution cache
	Cache CdnCachePtrInput
	// Set to true to configure DNS
	ConfigureDNS pulumi.BoolPtrInput
	// DNS configuration
	Dns CdnDnsPtrInput
	// Set to true to add an alias to wwww.<domain>
	Domain pulumi.StringPtrInput
	// Set to true to add an alias to wwww.<domain>
	IncludeWWW pulumi.BoolPtrInput
	// Cloud Front distribution priceClass
	PriceClass pulumi.StringPtrInput
}

func (StaticWebArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*staticWebArgs)(nil)).Elem()
}

type StaticWebInput interface {
	pulumi.Input

	ToStaticWebOutput() StaticWebOutput
	ToStaticWebOutputWithContext(ctx context.Context) StaticWebOutput
}

func (*StaticWeb) ElementType() reflect.Type {
	return reflect.TypeOf((**StaticWeb)(nil)).Elem()
}

func (i *StaticWeb) ToStaticWebOutput() StaticWebOutput {
	return i.ToStaticWebOutputWithContext(context.Background())
}

func (i *StaticWeb) ToStaticWebOutputWithContext(ctx context.Context) StaticWebOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StaticWebOutput)
}

// StaticWebArrayInput is an input type that accepts StaticWebArray and StaticWebArrayOutput values.
// You can construct a concrete instance of `StaticWebArrayInput` via:
//
//	StaticWebArray{ StaticWebArgs{...} }
type StaticWebArrayInput interface {
	pulumi.Input

	ToStaticWebArrayOutput() StaticWebArrayOutput
	ToStaticWebArrayOutputWithContext(context.Context) StaticWebArrayOutput
}

type StaticWebArray []StaticWebInput

func (StaticWebArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*StaticWeb)(nil)).Elem()
}

func (i StaticWebArray) ToStaticWebArrayOutput() StaticWebArrayOutput {
	return i.ToStaticWebArrayOutputWithContext(context.Background())
}

func (i StaticWebArray) ToStaticWebArrayOutputWithContext(ctx context.Context) StaticWebArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StaticWebArrayOutput)
}

// StaticWebMapInput is an input type that accepts StaticWebMap and StaticWebMapOutput values.
// You can construct a concrete instance of `StaticWebMapInput` via:
//
//	StaticWebMap{ "key": StaticWebArgs{...} }
type StaticWebMapInput interface {
	pulumi.Input

	ToStaticWebMapOutput() StaticWebMapOutput
	ToStaticWebMapOutputWithContext(context.Context) StaticWebMapOutput
}

type StaticWebMap map[string]StaticWebInput

func (StaticWebMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*StaticWeb)(nil)).Elem()
}

func (i StaticWebMap) ToStaticWebMapOutput() StaticWebMapOutput {
	return i.ToStaticWebMapOutputWithContext(context.Background())
}

func (i StaticWebMap) ToStaticWebMapOutputWithContext(ctx context.Context) StaticWebMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StaticWebMapOutput)
}

type StaticWebOutput struct{ *pulumi.OutputState }

func (StaticWebOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**StaticWeb)(nil)).Elem()
}

func (o StaticWebOutput) ToStaticWebOutput() StaticWebOutput {
	return o
}

func (o StaticWebOutput) ToStaticWebOutputWithContext(ctx context.Context) StaticWebOutput {
	return o
}

// DNS Records to expose staticweb
func (o StaticWebOutput) DNSRecords() DNSRecordsPtrOutput {
	return o.ApplyT(func(v *StaticWeb) DNSRecordsPtrOutput { return v.DNSRecords }).(DNSRecordsPtrOutput)
}

// DNS Records to validate the certificate
func (o StaticWebOutput) DNSRecordsForValidation() DNSRecordsPtrOutput {
	return o.ApplyT(func(v *StaticWeb) DNSRecordsPtrOutput { return v.DNSRecordsForValidation }).(DNSRecordsPtrOutput)
}

// CloudFront Distribution
func (o StaticWebOutput) Certificate() acm.CertificateOutput {
	return o.ApplyT(func(v *StaticWeb) acm.CertificateOutput { return v.Certificate }).(acm.CertificateOutput)
}

// AWS certificate validation
func (o StaticWebOutput) CertificateValidation() acm.CertificateValidationOutput {
	return o.ApplyT(func(v *StaticWeb) acm.CertificateValidationOutput { return v.CertificateValidation }).(acm.CertificateValidationOutput)
}

// Content bucket
func (o StaticWebOutput) ContentBucket() storage.BucketOutput {
	return o.ApplyT(func(v *StaticWeb) storage.BucketOutput { return v.ContentBucket }).(storage.BucketOutput)
}

// Bucket policy to connect Cloud Front distribution
func (o StaticWebOutput) ContentBucketPolicy() s3.BucketPolicyOutput {
	return o.ApplyT(func(v *StaticWeb) s3.BucketPolicyOutput { return v.ContentBucketPolicy }).(s3.BucketPolicyOutput)
}

// CloudFront Distribution
func (o StaticWebOutput) Distribution() cloudfront.DistributionOutput {
	return o.ApplyT(func(v *StaticWeb) cloudfront.DistributionOutput { return v.Distribution }).(cloudfront.DistributionOutput)
}

// CloudFront Distribution
func (o StaticWebOutput) DomainValidationOptions() acm.CertificateDomainValidationOptionArrayOutput {
	return o.ApplyT(func(v *StaticWeb) acm.CertificateDomainValidationOptionArrayOutput { return v.DomainValidationOptions }).(acm.CertificateDomainValidationOptionArrayOutput)
}

// AWS eastRegion provider. It is required to create and validate certificates
func (o StaticWebOutput) EastRegion() aws.ProviderOutput {
	return o.ApplyT(func(v *StaticWeb) aws.ProviderOutput { return v.EastRegion }).(aws.ProviderOutput)
}

// Logs bucket
func (o StaticWebOutput) LogsBucket() s3.BucketOutput {
	return o.ApplyT(func(v *StaticWeb) s3.BucketOutput { return v.LogsBucket }).(s3.BucketOutput)
}

// Staticweb name
func (o StaticWebOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *StaticWeb) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// OriginAccessIdentity to have access to content bucket
func (o StaticWebOutput) OriginAccessIdentity() cloudfront.OriginAccessIdentityOutput {
	return o.ApplyT(func(v *StaticWeb) cloudfront.OriginAccessIdentityOutput { return v.OriginAccessIdentity }).(cloudfront.OriginAccessIdentityOutput)
}

type StaticWebArrayOutput struct{ *pulumi.OutputState }

func (StaticWebArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*StaticWeb)(nil)).Elem()
}

func (o StaticWebArrayOutput) ToStaticWebArrayOutput() StaticWebArrayOutput {
	return o
}

func (o StaticWebArrayOutput) ToStaticWebArrayOutputWithContext(ctx context.Context) StaticWebArrayOutput {
	return o
}

func (o StaticWebArrayOutput) Index(i pulumi.IntInput) StaticWebOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *StaticWeb {
		return vs[0].([]*StaticWeb)[vs[1].(int)]
	}).(StaticWebOutput)
}

type StaticWebMapOutput struct{ *pulumi.OutputState }

func (StaticWebMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*StaticWeb)(nil)).Elem()
}

func (o StaticWebMapOutput) ToStaticWebMapOutput() StaticWebMapOutput {
	return o
}

func (o StaticWebMapOutput) ToStaticWebMapOutputWithContext(ctx context.Context) StaticWebMapOutput {
	return o
}

func (o StaticWebMapOutput) MapIndex(k pulumi.StringInput) StaticWebOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *StaticWeb {
		return vs[0].(map[string]*StaticWeb)[vs[1].(string)]
	}).(StaticWebOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*StaticWebInput)(nil)).Elem(), &StaticWeb{})
	pulumi.RegisterInputType(reflect.TypeOf((*StaticWebArrayInput)(nil)).Elem(), StaticWebArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*StaticWebMapInput)(nil)).Elem(), StaticWebMap{})
	pulumi.RegisterOutputType(StaticWebOutput{})
	pulumi.RegisterOutputType(StaticWebArrayOutput{})
	pulumi.RegisterOutputType(StaticWebMapOutput{})
}
