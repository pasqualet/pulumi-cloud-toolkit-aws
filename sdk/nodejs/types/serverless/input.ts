// *** WARNING: this file was generated by Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

import * as pulumiAws from "@pulumi/aws";
import * as pulumiKubernetes from "@pulumi/kubernetes";

export interface CdnCacheArgsArgs {
    /**
     * Cloud Front distribution cache time to live
     */
    ttl: pulumi.Input<number>;
}

export interface CdnDnsArgsArgs {
    /**
     * DNS time to live
     */
    ttl: pulumi.Input<number>;
}

export interface DeadLetterQueueTypeArgsArgs {
    /**
     * Enables the feature.
     */
    enable: pulumi.Input<boolean>;
    /**
     * Placing a Queue ARN will set said already existing Queue as a Dead Letter Queue for the new one.
     */
    existingDeadLetterQueueArn?: pulumi.Input<string>;
    /**
     * The amount of time that a message will be stored in the Dead Letter Queue without being deleted. Minimum is 60 seconds (1 minutes) and Maximum 1,209,600 (14 days) seconds. By default a message is retained 4 days.
     */
    messageRetentionSeconds?: pulumi.Input<number>;
    /**
     * Dead Letter Queue type attached to the component to create.
     */
    type: pulumi.Input<enums.serverless.DeadLetterQueueTypes>;
}

export interface QueueArgsArgs {
    /**
     * Dead Letter Queue attached to the component to create.
     */
    DeadLetterQueueTypeArgs?: pulumi.Input<inputs.serverless.DeadLetterQueueTypeArgsArgs>;
    /**
     * Set to true to create the Queue as FiFo. False for a Standard Queue.
     */
    isFifo?: pulumi.Input<boolean>;
    /**
     * The limit for a Queue message size in bytes. Minimum is 1 byte (1 character) and Maximum 262,144 bytes (256 KiB). By default a message can be 256 KiB large.
     */
    maxMessageSize?: pulumi.Input<number>;
    /**
     * The amount of time that a message will be stored in the Queue without being deleted. Minimum is 60 seconds (1 minutes) and Maximum 1,209,600 (14 days) seconds. By default a message is retained 4 days.
     */
    messageRetentionSeconds?: pulumi.Input<number>;
    /**
     * Custom policy for the Queue.
     */
    policy?: pulumi.Input<string>;
}

