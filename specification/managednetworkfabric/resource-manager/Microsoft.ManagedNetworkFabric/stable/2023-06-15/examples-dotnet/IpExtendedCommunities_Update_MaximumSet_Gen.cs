using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.ManagedNetworkFabric.Models;
using Azure.ResourceManager.ManagedNetworkFabric;

// Generated from example definition: specification/managednetworkfabric/resource-manager/Microsoft.ManagedNetworkFabric/stable/2023-06-15/examples/IpExtendedCommunities_Update_MaximumSet_Gen.json
// this example is just showing the usage of "IpExtendedCommunities_Update" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this NetworkFabricIPExtendedCommunityResource created on azure
// for more information of creating NetworkFabricIPExtendedCommunityResource, please refer to the document of NetworkFabricIPExtendedCommunityResource
string subscriptionId = "1234ABCD-0A1B-1234-5678-123456ABCDEF";
string resourceGroupName = "example-rg";
string ipExtendedCommunityName = "example-ipExtendedCommunity";
ResourceIdentifier networkFabricIPExtendedCommunityResourceId = NetworkFabricIPExtendedCommunityResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, ipExtendedCommunityName);
NetworkFabricIPExtendedCommunityResource networkFabricIPExtendedCommunity = client.GetNetworkFabricIPExtendedCommunityResource(networkFabricIPExtendedCommunityResourceId);

// invoke the operation
NetworkFabricIPExtendedCommunityPatch patch = new NetworkFabricIPExtendedCommunityPatch
{
    IPExtendedCommunityRules = { new IPExtendedCommunityRule(CommunityActionType.Permit, 4155123341L, new string[] { "1234:2345" }) },
    Tags =
    {
    ["keyID"] = "KeyValue"
    },
};
ArmOperation<NetworkFabricIPExtendedCommunityResource> lro = await networkFabricIPExtendedCommunity.UpdateAsync(WaitUntil.Completed, patch);
NetworkFabricIPExtendedCommunityResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
NetworkFabricIPExtendedCommunityData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
