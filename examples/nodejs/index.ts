import * as pulumi from "@pulumi/pulumi";
import * as cr from "@pulumi/cr";

const myRandomResource = new cr.Random("myRandomResource", {length: 24});
export const output = {
    value: myRandomResource.result,
};
