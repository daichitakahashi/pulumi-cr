import pulumi
import pulumi_cr as cr

my_random_resource = cr.Random("myRandomResource", length=24)
pulumi.export("output", {
    "value": my_random_resource.result,
})
