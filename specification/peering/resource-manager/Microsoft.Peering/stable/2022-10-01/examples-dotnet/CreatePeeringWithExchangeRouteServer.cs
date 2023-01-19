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

// Generated from example definition: specification/peering/resource-manager/Microsoft.Peering/stable/2022-10-01/examples/CreatePeeringWithExchangeRouteServer.json
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
    Name = "Premium_Direct_Free",
}, PeeringKind.Direct)
{
    Direct = new DirectPeeringProperties()
    {
        Connections =
        {
        new PeeringDirectConnection()
        {
        BandwidthInMbps = 10000,
        SessionAddressProvider = PeeringSessionAddressProvider.Peer,
        UseForPeeringService = true,
        PeeringDBFacilityId = 99999,
        BgpSession = new PeeringBgpSession()
        {
        SessionPrefixV4 = "192.168.0.0/24",
        MicrosoftSessionIPv4Address = IPAddress.Parse("192.168.0.123"),
        PeerSessionIPv4Address = IPAddress.Parse("192.168.0.234"),
        MaxPrefixesAdvertisedV4 = 1000,
        MaxPrefixesAdvertisedV6 = 100,
        },
        ConnectionIdentifier = "5F4CB5C7-6B43-4444-9338-9ABC72606C16",
        }
        },
        PeerAsnId = new ResourceIdentifier("/subscriptions/subId/providers/Microsoft.Peering/peerAsns/myAsn1"),
        DirectPeeringType = DirectPeeringType.IxRs,
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
