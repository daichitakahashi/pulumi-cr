// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Cr.Inputs
{

    public sealed class KVNamespaceBindingArgs : global::Pulumi.ResourceArgs
    {
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        [Input("namespaceId", required: true)]
        public Input<string> NamespaceId { get; set; } = null!;

        public KVNamespaceBindingArgs()
        {
        }
        public static new KVNamespaceBindingArgs Empty => new KVNamespaceBindingArgs();
    }
}
