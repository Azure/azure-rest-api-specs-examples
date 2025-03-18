using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Peering.Models;
using Azure.ResourceManager.Peering;

// Generated from example definition: specification/peering/resource-manager/Microsoft.Peering/stable/2022-10-01/examples/GetPeerAsn.json
// this example is just showing the usage of "PeerAsns_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this PeerAsnResource created on azure
// for more information of creating PeerAsnResource, please refer to the document of PeerAsnResource
string subscriptionId = "subId";
string peerAsnName = "peerAsnName";
ResourceIdentifier peerAsnResourceId = PeerAsnResource.CreateResourceIdentifier(subscriptionId, peerAsnName);
PeerAsnResource peerAsn = client.GetPeerAsnResource(peerAsnResourceId);

// invoke the operation
PeerAsnResource result = await peerAsn.GetAsync();

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
PeerAsnData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
