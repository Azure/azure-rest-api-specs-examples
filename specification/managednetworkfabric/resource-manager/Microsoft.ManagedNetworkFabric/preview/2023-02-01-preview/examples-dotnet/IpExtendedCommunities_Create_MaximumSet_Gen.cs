using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.ManagedNetworkFabric;
using Azure.ResourceManager.ManagedNetworkFabric.Models;
using Azure.ResourceManager.Resources;

// Generated from example definition: specification/managednetworkfabric/resource-manager/Microsoft.ManagedNetworkFabric/preview/2023-02-01-preview/examples/IpExtendedCommunities_Create_MaximumSet_Gen.json
// this example is just showing the usage of "IpExtendedCommunities_Create" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ResourceGroupResource created on azure
// for more information of creating ResourceGroupResource, please refer to the document of ResourceGroupResource
string subscriptionId = "subscriptionId";
string resourceGroupName = "rgIpExtendedCommunityLists";
ResourceIdentifier resourceGroupResourceId = ResourceGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName);
ResourceGroupResource resourceGroupResource = client.GetResourceGroupResource(resourceGroupResourceId);

// get the collection of this IPExtendedCommunityResource
IPExtendedCommunityCollection collection = resourceGroupResource.GetIPExtendedCommunities();

// invoke the operation
string ipExtendedCommunityName = "example_ipExtendedCommunity";
IPExtendedCommunityData data = new IPExtendedCommunityData(new AzureLocation("EastUs"))
{
    Annotation = "annotationValue",
    Action = CommunityActionType.Permit,
    RouteTargets =
    {
    "1234:5678"
    },
    Tags =
    {
    ["key5054"] = "key",
    },
};
ArmOperation<IPExtendedCommunityResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, ipExtendedCommunityName, data);
IPExtendedCommunityResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
IPExtendedCommunityData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
