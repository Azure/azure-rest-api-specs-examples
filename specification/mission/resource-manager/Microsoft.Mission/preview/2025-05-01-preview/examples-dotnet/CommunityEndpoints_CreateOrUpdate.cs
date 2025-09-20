using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.VirtualEnclaves.Models;
using Azure.ResourceManager.VirtualEnclaves;

// Generated from example definition: 2025-05-01-preview/CommunityEndpoints_CreateOrUpdate.json
// this example is just showing the usage of "CommunityEndpointResource_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this VirtualEnclaveCommunityResource created on azure
// for more information of creating VirtualEnclaveCommunityResource, please refer to the document of VirtualEnclaveCommunityResource
string subscriptionId = "73CEECEF-2C30-488E-946F-D20F414D99BA";
string resourceGroupName = "rgopenapi";
string communityName = "TestMyCommunity";
ResourceIdentifier virtualEnclaveCommunityResourceId = VirtualEnclaveCommunityResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, communityName);
VirtualEnclaveCommunityResource virtualEnclaveCommunity = client.GetVirtualEnclaveCommunityResource(virtualEnclaveCommunityResourceId);

// get the collection of this VirtualEnclaveCommunityEndpointResource
VirtualEnclaveCommunityEndpointCollection collection = virtualEnclaveCommunity.GetVirtualEnclaveCommunityEndpoints();

// invoke the operation
string communityEndpointName = "TestMyCommunityEndpoint";
VirtualEnclaveCommunityEndpointData data = new VirtualEnclaveCommunityEndpointData(new AzureLocation("West US"))
{
    Properties = new VirtualEnclaveCommunityEndpointProperties(new CommunityEndpointDestinationRule[]
{
new CommunityEndpointDestinationRule
{
DestinationType = CommunityEndpointDestinationType.FqdnTag,
Protocols = {CommunityEndpointProtocol.Tcp},
TransitHubResourceId = new ResourceIdentifier("/subscriptions/c64f6eca-bdc5-4bc2-88d6-f8f1dc23f86c/resourceGroups/testrg/providers/Microsoft.Mission/communities/TestMyCommunity/transitHubs/TestThName"),
Destination = "foo.example.com",
Ports = "443",
}
}),
    Tags =
    {
    ["sampletag"] = "samplevalue"
    },
};
ArmOperation<VirtualEnclaveCommunityEndpointResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, communityEndpointName, data);
VirtualEnclaveCommunityEndpointResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
VirtualEnclaveCommunityEndpointData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
