using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.NetworkCloud;
using Azure.ResourceManager.NetworkCloud.Models;
using Azure.ResourceManager.Resources;

// Generated from example definition: specification/networkcloud/resource-manager/Microsoft.NetworkCloud/preview/2022-12-12-preview/examples/DefaultCniNetworks_Create.json
// this example is just showing the usage of "DefaultCniNetworks_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ResourceGroupResource created on azure
// for more information of creating ResourceGroupResource, please refer to the document of ResourceGroupResource
string subscriptionId = "subscriptionId";
string resourceGroupName = "resourceGroupName";
ResourceIdentifier resourceGroupResourceId = ResourceGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName);
ResourceGroupResource resourceGroupResource = client.GetResourceGroupResource(resourceGroupResourceId);

// get the collection of this DefaultCniNetworkResource
DefaultCniNetworkCollection collection = resourceGroupResource.GetDefaultCniNetworks();

// invoke the operation
string defaultCniNetworkName = "defaultCniNetworkName";
DefaultCniNetworkData data = new DefaultCniNetworkData(new AzureLocation("location"), new ExtendedLocation("/subscriptions/subscriptionId/resourceGroups/resourceGroupName/providers/Microsoft.ExtendedLocation/customLocations/clusterExtendedLocationName", "CustomLocation"), "/subscriptions/subscriptionId/resourceGroups/resourceGroupName/providers/Microsoft.ManagedNetworkFabric/l3IsolationDomains/l3IsolationDomainName", 12)
{
    CniBgpConfiguration = new CniBgpConfiguration()
    {
        BgpPeers =
        {
        new BgpPeer(64497,"203.0.113.254")
        },
        CommunityAdvertisements =
        {
        new CommunityAdvertisement(new string[]
        {
        "64512:100"
        },"192.0.2.0/27")
        },
        ServiceExternalPrefixes =
        {
        "192.0.2.0/28"
        },
        ServiceLoadBalancerPrefixes =
        {
        "192.0.2.16/28"
        },
    },
    IPAllocationType = IPAllocationType.DualStack,
    IPv4ConnectedPrefix = "203.0.113.0/24",
    IPv6ConnectedPrefix = "2001:db8:0:3::/64",
    Tags =
    {
    ["key1"] = "myvalue1",
    ["key2"] = "myvalue2",
    },
};
ArmOperation<DefaultCniNetworkResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, defaultCniNetworkName, data);
DefaultCniNetworkResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
DefaultCniNetworkData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
