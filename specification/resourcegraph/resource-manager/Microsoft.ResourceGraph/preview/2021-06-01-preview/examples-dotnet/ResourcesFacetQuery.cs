using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.ResourceGraph;
using Azure.ResourceManager.ResourceGraph.Models;

// Generated from example definition: specification/resourcegraph/resource-manager/Microsoft.ResourceGraph/preview/2021-06-01-preview/examples/ResourcesFacetQuery.json
// this example is just showing the usage of "Resources" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this TenantResource created on azure
// for more information of creating TenantResource, please refer to the document of TenantResource
var tenantResource = client.GetTenants().GetAllAsync().GetAsyncEnumerator().Current;

// invoke the operation
ResourceQueryContent content = new ResourceQueryContent("Resources | where type =~ 'Microsoft.Compute/virtualMachines' | project id, name, location, resourceGroup, properties.storageProfile.osDisk.osType | limit 5")
{
    Subscriptions =
    {
    "cfbbd179-59d2-4052-aa06-9270a38aa9d6"
    },
    Facets =
    {
    new FacetRequest("location")
    {
    Options = new FacetRequestOptions()
    {
    SortOrder = FacetSortOrder.Desc,
    Top = 3,
    },
    },new FacetRequest("properties.storageProfile.osDisk.osType")
    {
    Options = new FacetRequestOptions()
    {
    SortOrder = FacetSortOrder.Desc,
    Top = 3,
    },
    },new FacetRequest("nonExistingColumn")
    {
    Options = new FacetRequestOptions()
    {
    SortOrder = FacetSortOrder.Desc,
    Top = 3,
    },
    },new FacetRequest("resourceGroup")
    {
    Options = new FacetRequestOptions()
    {
    SortBy = "tolower(resourceGroup)",
    SortOrder = FacetSortOrder.Asc,
    Top = 3,
    },
    },new FacetRequest("resourceGroup")
    {
    Options = new FacetRequestOptions()
    {
    Filter = "resourceGroup contains 'test'",
    Top = 3,
    },
    }
    },
};
ResourceQueryResult result = await tenantResource.GetResourcesAsync(content);

Console.WriteLine($"Succeeded: {result}");
