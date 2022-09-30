// *** WARNING: this file was generated by Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export { AccountIamArgs } from "./accountIam";
export type AccountIam = import("./accountIam").AccountIam;
export const AccountIam: typeof import("./accountIam").AccountIam = null as any;

utilities.lazyLoad(exports, ["AccountIam"], () => require("./accountIam"));

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "cloud-toolkit-aws:landingZone:AccountIam":
                return new AccountIam(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("cloud-toolkit-aws", "landingZone", _module)