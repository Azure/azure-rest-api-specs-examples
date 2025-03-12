using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.ResourceGraph.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.ResourceGraph;

// Generated from example definition: specification/resourcegraph/resource-manager/Microsoft.ResourceGraph/preview/2021-06-01-preview/examples/ResourcesFirstPageQuery.json
// this example is just showing the usage of "Resources" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

TenantResource tenantResource = client.GetTenants().GetAllAsync().GetAsyncEnumerator().Current;

// invoke the operation
ResourceQueryContent content = new ResourceQueryContent("Resources | where name contains 'test' | project id, name, type, location")
{
    Subscriptions = { "cfbbd179-59d2-4052-aa06-9270a38aa9d6" },
    Options = new ResourceQueryRequestOptions
    {
        Top = 3,
        Skip = 0,
    },
};
ResourceQueryResult result = await tenantResource.GetResourcesAsync(content);

Console.WriteLine($"Succeeded: {result}");
