using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.Search.Models;
using Azure.ResourceManager.Search;

// Generated from example definition: specification/search/resource-manager/Microsoft.Search/stable/2025-05-01/examples/SearchCreateOrUpdateServiceToAllowAccessFromPublicCustomIPsAndBypass.json
// this example is just showing the usage of "Services_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ResourceGroupResource created on azure
// for more information of creating ResourceGroupResource, please refer to the document of ResourceGroupResource
string subscriptionId = "subid";
string resourceGroupName = "rg1";
ResourceIdentifier resourceGroupResourceId = ResourceGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName);
ResourceGroupResource resourceGroupResource = client.GetResourceGroupResource(resourceGroupResourceId);

// get the collection of this SearchServiceResource
SearchServiceCollection collection = resourceGroupResource.GetSearchServices();

// invoke the operation
string searchServiceName = "mysearchservice";
SearchServiceData data = new SearchServiceData(new AzureLocation("westus"))
{
    SearchSkuName = SearchServiceSkuName.Standard,
    ReplicaCount = 1,
    PartitionCount = 1,
    HostingMode = SearchServiceHostingMode.Default,
    ComputeType = SearchServiceComputeType.Default,
    NetworkRuleSet = new SearchServiceNetworkRuleSet
    {
        IPRules = {new SearchServiceIPRule
        {
        Value = "123.4.5.6",
        }, new SearchServiceIPRule
        {
        Value = "123.4.6.0/18",
        }},
        Bypass = SearchBypass.AzureServices,
    },
    Tags =
    {
    ["app-name"] = "My e-commerce app"
    },
};
ArmOperation<SearchServiceResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, searchServiceName, data);
SearchServiceResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
SearchServiceData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
