// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

declare var exports: any;
const __config = new pulumi.Config("cr");

export declare const cloudflareApiToken: string | undefined;
Object.defineProperty(exports, "cloudflareApiToken", {
    get() {
        return __config.get("cloudflareApiToken");
    },
    enumerable: true,
});

