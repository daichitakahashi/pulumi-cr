using System.Collections.Generic;
using System.Linq;
using Pulumi;
using cr = Pulumi.cr;

return await Deployment.RunAsync(() => 
{
    var myRandomResource = new cr.Random("myRandomResource", new()
    {
        Length = 24,
    });

    return new Dictionary<string, object?>
    {
        ["output"] = 
        {
            { "value", myRandomResource.Result },
        },
    };
});

