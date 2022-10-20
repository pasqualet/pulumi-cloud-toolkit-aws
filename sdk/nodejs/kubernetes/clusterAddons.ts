// *** WARNING: this file was generated by Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as enums from "../types/enums";
import * as utilities from "../utilities";

import * as pulumiKubernetes from "@pulumi/kubernetes";

import {ArgoCD, Calico, CertManager, Dashboard, ExternalDns, IngressNginx} from "./index";

/**
 * ClusterAddons is a component that manages the Lubernetes addons to setup a production-ready cluster.
 */
export class ClusterAddons extends pulumi.ComponentResource {
    /** @internal */
    public static readonly __pulumiType = 'cloud-toolkit-aws:kubernetes:ClusterAddons';

    /**
     * Returns true if the given object is an instance of ClusterAddons.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ClusterAddons {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ClusterAddons.__pulumiType;
    }

    /**
     * The IngressNginx addon used for admin access.
     */
    public /*out*/ readonly adminIngressNginx!: pulumi.Output<IngressNginx>;
    /**
     * The ArgoCD addon.
     */
    public /*out*/ readonly argocd!: pulumi.Output<ArgoCD>;
    /**
     * The Calico addon used to manage network policies.
     */
    public /*out*/ readonly calico!: pulumi.Output<Calico>;
    /**
     * The CertManager addon.
     */
    public /*out*/ readonly certManager!: pulumi.Output<CertManager>;
    /**
     * The Kubernetes dashboard addon.
     */
    public /*out*/ readonly dashboard!: pulumi.Output<Dashboard>;
    /**
     * The ExternalDns addon.
     */
    public /*out*/ readonly externalDns!: pulumi.Output<ExternalDns>;

    /**
     * Create a ClusterAddons resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ClusterAddonsArgs, opts?: pulumi.ComponentResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.domain === undefined) && !opts.urn) {
                throw new Error("Missing required property 'domain'");
            }
            if ((!args || args.identityProvidersArn === undefined) && !opts.urn) {
                throw new Error("Missing required property 'identityProvidersArn'");
            }
            if ((!args || args.issuerUrl === undefined) && !opts.urn) {
                throw new Error("Missing required property 'issuerUrl'");
            }
            if ((!args || args.k8sProvider === undefined) && !opts.urn) {
                throw new Error("Missing required property 'k8sProvider'");
            }
            if ((!args || args.zoneArns === undefined) && !opts.urn) {
                throw new Error("Missing required property 'zoneArns'");
            }
            if ((!args || args.zoneId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'zoneId'");
            }
            resourceInputs["domain"] = args ? args.domain : undefined;
            resourceInputs["identityProvidersArn"] = args ? args.identityProvidersArn : undefined;
            resourceInputs["ingress"] = args ? args.ingress : undefined;
            resourceInputs["issuerUrl"] = args ? args.issuerUrl : undefined;
            resourceInputs["k8sProvider"] = args ? args.k8sProvider : undefined;
            resourceInputs["zoneArns"] = args ? args.zoneArns : undefined;
            resourceInputs["zoneId"] = args ? args.zoneId : undefined;
            resourceInputs["adminIngressNginx"] = undefined /*out*/;
            resourceInputs["argocd"] = undefined /*out*/;
            resourceInputs["calico"] = undefined /*out*/;
            resourceInputs["certManager"] = undefined /*out*/;
            resourceInputs["dashboard"] = undefined /*out*/;
            resourceInputs["externalDns"] = undefined /*out*/;
        } else {
            resourceInputs["adminIngressNginx"] = undefined /*out*/;
            resourceInputs["argocd"] = undefined /*out*/;
            resourceInputs["calico"] = undefined /*out*/;
            resourceInputs["certManager"] = undefined /*out*/;
            resourceInputs["dashboard"] = undefined /*out*/;
            resourceInputs["externalDns"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ClusterAddons.__pulumiType, name, resourceInputs, opts, true /*remote*/);
    }
}

/**
 * The set of arguments for constructing a ClusterAddons resource.
 */
export interface ClusterAddonsArgs {
    /**
     * The domain used by the cluster.
     */
    domain: pulumi.Input<string>;
    /**
     * The OIDC Identity Provider arn.
     */
    identityProvidersArn: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The configuration for Ingress Controller.
     */
    ingress?: pulumi.Input<inputs.kubernetes.ClusterAddonsIngressArgsArgs>;
    /**
     * The OIDC Identity Provider url.
     */
    issuerUrl: pulumi.Input<string>;
    /**
     * The Pulumi provider used for Kubernetes resources.
     */
    k8sProvider: pulumi.Input<pulumiKubernetes.Provider>;
    /**
     * The list of DNS Zone arns to be used by CertManager and ExternalDNS.
     */
    zoneArns: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The main DNS Zone id.
     */
    zoneId: pulumi.Input<string>;
}
