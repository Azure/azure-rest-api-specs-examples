using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.ResourceGraph.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.ResourceGraph;

// Generated from example definition: specification/resourcegraph/resource-manager/Microsoft.ResourceGraph/preview/2021-06-01-preview/examples/ResourcesHistoryMgsGet.json
// this example is just showing the usage of "ResourcesHistory" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

TenantResource tenantResource = client.GetTenants().GetAllAsync().GetAsyncEnumerator().Current;

// invoke the operation
ResourcesHistoryContent content = new ResourcesHistoryContent
{
    Query = "where name =~ 'cpu-utilization' | project id, name, properties",
    Options = new ResourcesHistoryRequestOptions
    {
        Interval = new DateTimeInterval(DateTimeOffset.Parse("2020-11-12T01:00:00.0000000Z"), DateTimeOffset.Parse("2020-11-12T01:25:00.0000000Z")),
    },
    ManagementGroups = { "e927f598-c1d4-4f72-8541-95d83a6a4ac8", "ProductionMG" },
};
BinaryData result = await tenantResource.GetResourceHistoryAsync(content);

Console.WriteLine($"Succeeded: {result}");
