using System;
using System.Net;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Peering;
using Azure.ResourceManager.Peering.Models;
using Azure.ResourceManager.Resources;

// Generated from example definition: specification/peering/resource-manager/Microsoft.Peering/stable/2022-10-01/examples/CreateExchangePeering.json
// this example is just showing the usage of "Peerings_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ResourceGroupResource created on azure
// for more information of creating ResourceGroupResource, please refer to the document of ResourceGroupResource
string subscriptionId = "subId";
string resourceGroupName = "rgName";
ResourceIdentifier resourceGroupResourceId = ResourceGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName);
ResourceGroupResource resourceGroupResource = client.GetResourceGroupResource(resourceGroupResourceId);

// get the collection of this PeeringResource
PeeringCollection collection = resourceGroupResource.GetPeerings();

// invoke the operation
string peeringName = "peeringName";
PeeringData data = new PeeringData(new AzureLocation("eastus"), new PeeringSku()
{
    Name = "Basic_Exchange_Free",
}, PeeringKind.Exchange)
{
    Exchange = new ExchangePeeringProperties()
    {
        Connections =
        {
        new PeeringExchangeConnection()
        {
        PeeringDBFacilityId = 99999,
        BgpSession = new PeeringBgpSession()
        {
        PeerSessionIPv4Address = IPAddress.Parse("192.168.2.1"),
        PeerSessionIPv6Address = IPAddress.Parse("fd00::1"),
        MaxPrefixesAdvertisedV4 = 1000,
        MaxPrefixesAdvertisedV6 = 100,
        Md5AuthenticationKey = "test-md5-auth-key",
        },
        ConnectionIdentifier = Guid.Parse("CE495334-0E94-4E51-8164-8116D6CD284D"),
        },new PeeringExchangeConnection()
        {
        PeeringDBFacilityId = 99999,
        BgpSession = new PeeringBgpSession()
        {
        PeerSessionIPv4Address = IPAddress.Parse("192.168.2.2"),
        PeerSessionIPv6Address = IPAddress.Parse("fd00::2"),
        MaxPrefixesAdvertisedV4 = 1000,
        MaxPrefixesAdvertisedV6 = 100,
        Md5AuthenticationKey = "test-md5-auth-key",
        },
        ConnectionIdentifier = Guid.Parse("CDD8E673-CB07-47E6-84DE-3739F778762B"),
        }
        },
        PeerAsnId = new ResourceIdentifier("/subscriptions/subId/providers/Microsoft.Peering/peerAsns/myAsn1"),
    },
    PeeringLocation = "peeringLocation0",
};
ArmOperation<PeeringResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, peeringName, data);
PeeringResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
PeeringData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
