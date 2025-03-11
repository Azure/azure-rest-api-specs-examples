using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.ResourceGraph.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.ResourceGraph;

// Generated from example definition: specification/resourcegraph/resource-manager/Microsoft.ResourceGraph/preview/2021-06-01-preview/examples/ResourcesMgBasicQuery.json
// this example is just showing the usage of "Resources" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

TenantResource tenantResource = client.GetTenants().GetAllAsync().GetAsyncEnumerator().Current;

// invoke the operation
ResourceQueryContent content = new ResourceQueryContent("Resources | project id, name, type, location, tags | limit 3")
{
    ManagementGroups = { "e927f598-c1d4-4f72-8541-95d83a6a4ac8", "ProductionMG" },
};
ResourceQueryResult result = await tenantResource.GetResourcesAsync(content);

Console.WriteLine($"Succeeded: {result}");
