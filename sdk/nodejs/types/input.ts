// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";

export interface AnalyticsEngineBindingArgs {
    dataset: pulumi.Input<string>;
    name: pulumi.Input<string>;
}

export interface D1DatabaseBindingArgs {
    databaseId: pulumi.Input<string>;
    name: pulumi.Input<string>;
}

export interface KVNamespaceBindingArgs {
    name: pulumi.Input<string>;
    namespaceId: pulumi.Input<string>;
}

export interface PlainTextBindingArgs {
    name: pulumi.Input<string>;
    text: pulumi.Input<string>;
}

export interface QueueBindingArgs {
    name: pulumi.Input<string>;
    queueName: pulumi.Input<string>;
}

export interface R2BucketBindingArgs {
    bucketName: pulumi.Input<string>;
    name: pulumi.Input<string>;
}

export interface SecretTextBindingArgs {
    name: pulumi.Input<string>;
    text: pulumi.Input<string>;
}

export interface ServiceBindingArgs {
    name: pulumi.Input<string>;
    service: pulumi.Input<string>;
}