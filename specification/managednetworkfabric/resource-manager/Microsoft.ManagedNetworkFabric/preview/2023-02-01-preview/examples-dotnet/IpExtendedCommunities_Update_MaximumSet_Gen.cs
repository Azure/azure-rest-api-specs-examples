using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.ManagedNetworkFabric;
using Azure.ResourceManager.ManagedNetworkFabric.Models;
using Azure.ResourceManager.Resources;

// Generated from example definition: specification/managednetworkfabric/resource-manager/Microsoft.ManagedNetworkFabric/preview/2023-02-01-preview/examples/IpExtendedCommunities_Update_MaximumSet_Gen.json
// this example is just showing the usage of "IpExtendedCommunities_Update" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this IPExtendedCommunityResource created on azure
// for more information of creating IPExtendedCommunityResource, please refer to the document of IPExtendedCommunityResource
string subscriptionId = "subscriptionId";
string resourceGroupName = "rgIpExtendedCommunityLists";
string ipExtendedCommunityName = "example_ipExtendedCommunity";
ResourceIdentifier ipExtendedCommunityResourceId = IPExtendedCommunityResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, ipExtendedCommunityName);
IPExtendedCommunityResource ipExtendedCommunity = client.GetIPExtendedCommunityResource(ipExtendedCommunityResourceId);

// invoke the operation
IPExtendedCommunityPatch patch = new IPExtendedCommunityPatch()
{
    Tags =
    {
    ["key5054"] = "key1",
    },
};
ArmOperation<IPExtendedCommunityResource> lro = await ipExtendedCommunity.UpdateAsync(WaitUntil.Completed, patch);
IPExtendedCommunityResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
IPExtendedCommunityData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
