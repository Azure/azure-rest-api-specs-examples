using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.ResourceGraph;
using Azure.ResourceManager.ResourceGraph.Models;

// Generated from example definition: specification/resourcegraph/resource-manager/Microsoft.ResourceGraph/preview/2021-06-01-preview/examples/ResourcesBasicQuery.json
// this example is just showing the usage of "Resources" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this TenantResource created on azure
// for more information of creating TenantResource, please refer to the document of TenantResource
var tenantResource = client.GetTenants().GetAllAsync().GetAsyncEnumerator().Current;

// invoke the operation
ResourceQueryContent content = new ResourceQueryContent("Resources | project id, name, type, location, tags | limit 3")
{
    Subscriptions =
    {
    "cfbbd179-59d2-4052-aa06-9270a38aa9d6"
    },
};
ResourceQueryResult result = await tenantResource.GetResourcesAsync(content);

Console.WriteLine($"Succeeded: {result}");
