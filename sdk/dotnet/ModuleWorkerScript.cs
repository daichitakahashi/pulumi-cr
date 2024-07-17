// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Cr
{
    [CrResourceType("cr:index:ModuleWorkerScript")]
    public partial class ModuleWorkerScript : global::Pulumi.CustomResource
    {
        [Output("accountId")]
        public Output<string> AccountId { get; private set; } = null!;

        [Output("analyticsEngineBindings")]
        public Output<ImmutableArray<Outputs.AnalyticsEngineBinding>> AnalyticsEngineBindings { get; private set; } = null!;

        [Output("compatibilityDate")]
        public Output<string?> CompatibilityDate { get; private set; } = null!;

        [Output("compatibilityFlags")]
        public Output<ImmutableArray<string>> CompatibilityFlags { get; private set; } = null!;

        [Output("d1DatabaseBindings")]
        public Output<ImmutableArray<Outputs.D1DatabaseBinding>> D1DatabaseBindings { get; private set; } = null!;

        [Output("dir")]
        public Output<string> Dir { get; private set; } = null!;

        [Output("eTag")]
        public Output<string> ETag { get; private set; } = null!;

        [Output("kvNamespaceBindings")]
        public Output<ImmutableArray<Outputs.KVNamespaceBinding>> KvNamespaceBindings { get; private set; } = null!;

        [Output("logpush")]
        public Output<bool?> Logpush { get; private set; } = null!;

        [Output("mainModule")]
        public Output<string> MainModule { get; private set; } = null!;

        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        [Output("plainTextBindings")]
        public Output<ImmutableArray<Outputs.PlainTextBinding>> PlainTextBindings { get; private set; } = null!;

        [Output("queueBindings")]
        public Output<ImmutableArray<Outputs.QueueBinding>> QueueBindings { get; private set; } = null!;

        [Output("r2BucketBindings")]
        public Output<ImmutableArray<Outputs.R2BucketBinding>> R2BucketBindings { get; private set; } = null!;

        [Output("secretTextBindings")]
        public Output<ImmutableArray<Outputs.SecretTextBinding>> SecretTextBindings { get; private set; } = null!;

        [Output("serviceBindings")]
        public Output<ImmutableArray<Outputs.ServiceBinding>> ServiceBindings { get; private set; } = null!;


        /// <summary>
        /// Create a ModuleWorkerScript resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ModuleWorkerScript(string name, ModuleWorkerScriptArgs args, CustomResourceOptions? options = null)
            : base("cr:index:ModuleWorkerScript", name, args ?? new ModuleWorkerScriptArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ModuleWorkerScript(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("cr:index:ModuleWorkerScript", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing ModuleWorkerScript resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ModuleWorkerScript Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new ModuleWorkerScript(name, id, options);
        }
    }

    public sealed class ModuleWorkerScriptArgs : global::Pulumi.ResourceArgs
    {
        [Input("accountId", required: true)]
        public Input<string> AccountId { get; set; } = null!;

        [Input("analyticsEngineBindings")]
        private InputList<Inputs.AnalyticsEngineBindingArgs>? _analyticsEngineBindings;
        public InputList<Inputs.AnalyticsEngineBindingArgs> AnalyticsEngineBindings
        {
            get => _analyticsEngineBindings ?? (_analyticsEngineBindings = new InputList<Inputs.AnalyticsEngineBindingArgs>());
            set => _analyticsEngineBindings = value;
        }

        [Input("compatibilityDate")]
        public Input<string>? CompatibilityDate { get; set; }

        [Input("compatibilityFlags")]
        private InputList<string>? _compatibilityFlags;
        public InputList<string> CompatibilityFlags
        {
            get => _compatibilityFlags ?? (_compatibilityFlags = new InputList<string>());
            set => _compatibilityFlags = value;
        }

        [Input("d1DatabaseBindings")]
        private InputList<Inputs.D1DatabaseBindingArgs>? _d1DatabaseBindings;
        public InputList<Inputs.D1DatabaseBindingArgs> D1DatabaseBindings
        {
            get => _d1DatabaseBindings ?? (_d1DatabaseBindings = new InputList<Inputs.D1DatabaseBindingArgs>());
            set => _d1DatabaseBindings = value;
        }

        [Input("dir", required: true)]
        public Input<string> Dir { get; set; } = null!;

        [Input("kvNamespaceBindings")]
        private InputList<Inputs.KVNamespaceBindingArgs>? _kvNamespaceBindings;
        public InputList<Inputs.KVNamespaceBindingArgs> KvNamespaceBindings
        {
            get => _kvNamespaceBindings ?? (_kvNamespaceBindings = new InputList<Inputs.KVNamespaceBindingArgs>());
            set => _kvNamespaceBindings = value;
        }

        [Input("logpush")]
        public Input<bool>? Logpush { get; set; }

        [Input("mainModule", required: true)]
        public Input<string> MainModule { get; set; } = null!;

        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        [Input("plainTextBindings")]
        private InputList<Inputs.PlainTextBindingArgs>? _plainTextBindings;
        public InputList<Inputs.PlainTextBindingArgs> PlainTextBindings
        {
            get => _plainTextBindings ?? (_plainTextBindings = new InputList<Inputs.PlainTextBindingArgs>());
            set => _plainTextBindings = value;
        }

        [Input("queueBindings")]
        private InputList<Inputs.QueueBindingArgs>? _queueBindings;
        public InputList<Inputs.QueueBindingArgs> QueueBindings
        {
            get => _queueBindings ?? (_queueBindings = new InputList<Inputs.QueueBindingArgs>());
            set => _queueBindings = value;
        }

        [Input("r2BucketBindings")]
        private InputList<Inputs.R2BucketBindingArgs>? _r2BucketBindings;
        public InputList<Inputs.R2BucketBindingArgs> R2BucketBindings
        {
            get => _r2BucketBindings ?? (_r2BucketBindings = new InputList<Inputs.R2BucketBindingArgs>());
            set => _r2BucketBindings = value;
        }

        [Input("secretTextBindings")]
        private InputList<Inputs.SecretTextBindingArgs>? _secretTextBindings;
        public InputList<Inputs.SecretTextBindingArgs> SecretTextBindings
        {
            get => _secretTextBindings ?? (_secretTextBindings = new InputList<Inputs.SecretTextBindingArgs>());
            set => _secretTextBindings = value;
        }

        [Input("serviceBindings")]
        private InputList<Inputs.ServiceBindingArgs>? _serviceBindings;
        public InputList<Inputs.ServiceBindingArgs> ServiceBindings
        {
            get => _serviceBindings ?? (_serviceBindings = new InputList<Inputs.ServiceBindingArgs>());
            set => _serviceBindings = value;
        }

        public ModuleWorkerScriptArgs()
        {
        }
        public static new ModuleWorkerScriptArgs Empty => new ModuleWorkerScriptArgs();
    }
}
