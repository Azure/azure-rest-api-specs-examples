using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Peering;
using Azure.ResourceManager.Peering.Models;
using Azure.ResourceManager.Resources;

// Generated from example definition: specification/peering/resource-manager/Microsoft.Peering/stable/2022-10-01/examples/CreatePeerAsn.json
// this example is just showing the usage of "PeerAsns_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SubscriptionResource created on azure
// for more information of creating SubscriptionResource, please refer to the document of SubscriptionResource
string subscriptionId = "subId";
ResourceIdentifier subscriptionResourceId = SubscriptionResource.CreateResourceIdentifier(subscriptionId);
SubscriptionResource subscriptionResource = client.GetSubscriptionResource(subscriptionResourceId);

// get the collection of this PeerAsnResource
PeerAsnCollection collection = subscriptionResource.GetPeerAsns();

// invoke the operation
string peerAsnName = "peerAsnName";
PeerAsnData data = new PeerAsnData()
{
    PeerAsn = 65000,
    PeerContactDetail =
    {
    new PeerAsnContactDetail()
    {
    Role = PeeringRole.Noc,
    Email = "noc@contoso.com",
    Phone = "+1 (234) 567-8999",
    },new PeerAsnContactDetail()
    {
    Role = PeeringRole.Policy,
    Email = "abc@contoso.com",
    Phone = "+1 (234) 567-8900",
    },new PeerAsnContactDetail()
    {
    Role = PeeringRole.Technical,
    Email = "xyz@contoso.com",
    Phone = "+1 (234) 567-8900",
    }
    },
    PeerName = "Contoso",
};
ArmOperation<PeerAsnResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, peerAsnName, data);
PeerAsnResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
PeerAsnData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
