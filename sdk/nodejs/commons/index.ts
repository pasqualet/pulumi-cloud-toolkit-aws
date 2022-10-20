// *** WARNING: this file was generated by Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export { CertificateArgs } from "./certificate";
export type Certificate = import("./certificate").Certificate;
export const Certificate: typeof import("./certificate").Certificate = null as any;

utilities.lazyLoad(exports, ["Certificate"], () => require("./certificate"));

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "cloud-toolkit-aws:commons:Certificate":
                return new Certificate(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("cloud-toolkit-aws", "commons", _module)
