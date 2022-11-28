// *** WARNING: this file was generated by Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export { AdotApplicationArgs } from "./adotApplication";
export type AdotApplication = import("./adotApplication").AdotApplication;
export const AdotApplication: typeof import("./adotApplication").AdotApplication = null as any;
utilities.lazyLoad(exports, ["AdotApplication"], () => require("./adotApplication"));

export { AdotOperatorArgs } from "./adotOperator";
export type AdotOperator = import("./adotOperator").AdotOperator;
export const AdotOperator: typeof import("./adotOperator").AdotOperator = null as any;
utilities.lazyLoad(exports, ["AdotOperator"], () => require("./adotOperator"));

export { ApplicationAddonArgs } from "./applicationAddon";
export type ApplicationAddon = import("./applicationAddon").ApplicationAddon;
export const ApplicationAddon: typeof import("./applicationAddon").ApplicationAddon = null as any;
utilities.lazyLoad(exports, ["ApplicationAddon"], () => require("./applicationAddon"));

export { ArgoCDArgs } from "./argoCD";
export type ArgoCD = import("./argoCD").ArgoCD;
export const ArgoCD: typeof import("./argoCD").ArgoCD = null as any;
utilities.lazyLoad(exports, ["ArgoCD"], () => require("./argoCD"));

export { AwsEbsCsiDriverArgs } from "./awsEbsCsiDriver";
export type AwsEbsCsiDriver = import("./awsEbsCsiDriver").AwsEbsCsiDriver;
export const AwsEbsCsiDriver: typeof import("./awsEbsCsiDriver").AwsEbsCsiDriver = null as any;
utilities.lazyLoad(exports, ["AwsEbsCsiDriver"], () => require("./awsEbsCsiDriver"));

export { AwsLoadBalancerControllerArgs } from "./awsLoadBalancerController";
export type AwsLoadBalancerController = import("./awsLoadBalancerController").AwsLoadBalancerController;
export const AwsLoadBalancerController: typeof import("./awsLoadBalancerController").AwsLoadBalancerController = null as any;
utilities.lazyLoad(exports, ["AwsLoadBalancerController"], () => require("./awsLoadBalancerController"));

export { CalicoArgs } from "./calico";
export type Calico = import("./calico").Calico;
export const Calico: typeof import("./calico").Calico = null as any;
utilities.lazyLoad(exports, ["Calico"], () => require("./calico"));

export { CertManagerArgs } from "./certManager";
export type CertManager = import("./certManager").CertManager;
export const CertManager: typeof import("./certManager").CertManager = null as any;
utilities.lazyLoad(exports, ["CertManager"], () => require("./certManager"));

export { ClusterArgs } from "./cluster";
export type Cluster = import("./cluster").Cluster;
export const Cluster: typeof import("./cluster").Cluster = null as any;
utilities.lazyLoad(exports, ["Cluster"], () => require("./cluster"));

export { ClusterAddonsArgs } from "./clusterAddons";
export type ClusterAddons = import("./clusterAddons").ClusterAddons;
export const ClusterAddons: typeof import("./clusterAddons").ClusterAddons = null as any;
utilities.lazyLoad(exports, ["ClusterAddons"], () => require("./clusterAddons"));

export { ClusterAutoscalerArgs } from "./clusterAutoscaler";
export type ClusterAutoscaler = import("./clusterAutoscaler").ClusterAutoscaler;
export const ClusterAutoscaler: typeof import("./clusterAutoscaler").ClusterAutoscaler = null as any;
utilities.lazyLoad(exports, ["ClusterAutoscaler"], () => require("./clusterAutoscaler"));

export { DashboardArgs } from "./dashboard";
export type Dashboard = import("./dashboard").Dashboard;
export const Dashboard: typeof import("./dashboard").Dashboard = null as any;
utilities.lazyLoad(exports, ["Dashboard"], () => require("./dashboard"));

export { ExternalDnsArgs } from "./externalDns";
export type ExternalDns = import("./externalDns").ExternalDns;
export const ExternalDns: typeof import("./externalDns").ExternalDns = null as any;
utilities.lazyLoad(exports, ["ExternalDns"], () => require("./externalDns"));

export { IngressNginxArgs } from "./ingressNginx";
export type IngressNginx = import("./ingressNginx").IngressNginx;
export const IngressNginx: typeof import("./ingressNginx").IngressNginx = null as any;
utilities.lazyLoad(exports, ["IngressNginx"], () => require("./ingressNginx"));

export { IrsaArgs } from "./irsa";
export type Irsa = import("./irsa").Irsa;
export const Irsa: typeof import("./irsa").Irsa = null as any;
utilities.lazyLoad(exports, ["Irsa"], () => require("./irsa"));

export { NodeGroupArgs } from "./nodeGroup";
export type NodeGroup = import("./nodeGroup").NodeGroup;
export const NodeGroup: typeof import("./nodeGroup").NodeGroup = null as any;
utilities.lazyLoad(exports, ["NodeGroup"], () => require("./nodeGroup"));


// Export enums:
export * from "../types/enums/kubernetes";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "cloud-toolkit-aws:kubernetes:AdotApplication":
                return new AdotApplication(name, <any>undefined, { urn })
            case "cloud-toolkit-aws:kubernetes:AdotOperator":
                return new AdotOperator(name, <any>undefined, { urn })
            case "cloud-toolkit-aws:kubernetes:ApplicationAddon":
                return new ApplicationAddon(name, <any>undefined, { urn })
            case "cloud-toolkit-aws:kubernetes:ArgoCD":
                return new ArgoCD(name, <any>undefined, { urn })
            case "cloud-toolkit-aws:kubernetes:AwsEbsCsiDriver":
                return new AwsEbsCsiDriver(name, <any>undefined, { urn })
            case "cloud-toolkit-aws:kubernetes:AwsLoadBalancerController":
                return new AwsLoadBalancerController(name, <any>undefined, { urn })
            case "cloud-toolkit-aws:kubernetes:Calico":
                return new Calico(name, <any>undefined, { urn })
            case "cloud-toolkit-aws:kubernetes:CertManager":
                return new CertManager(name, <any>undefined, { urn })
            case "cloud-toolkit-aws:kubernetes:Cluster":
                return new Cluster(name, <any>undefined, { urn })
            case "cloud-toolkit-aws:kubernetes:ClusterAddons":
                return new ClusterAddons(name, <any>undefined, { urn })
            case "cloud-toolkit-aws:kubernetes:ClusterAutoscaler":
                return new ClusterAutoscaler(name, <any>undefined, { urn })
            case "cloud-toolkit-aws:kubernetes:Dashboard":
                return new Dashboard(name, <any>undefined, { urn })
            case "cloud-toolkit-aws:kubernetes:ExternalDns":
                return new ExternalDns(name, <any>undefined, { urn })
            case "cloud-toolkit-aws:kubernetes:IngressNginx":
                return new IngressNginx(name, <any>undefined, { urn })
            case "cloud-toolkit-aws:kubernetes:Irsa":
                return new Irsa(name, <any>undefined, { urn })
            case "cloud-toolkit-aws:kubernetes:NodeGroup":
                return new NodeGroup(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("cloud-toolkit-aws", "kubernetes", _module)
