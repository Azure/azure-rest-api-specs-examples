using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Peering;
using Azure.ResourceManager.Peering.Models;

// Generated from example definition: specification/peering/resource-manager/Microsoft.Peering/stable/2022-10-01/examples/DeletePeerAsn.json
// this example is just showing the usage of "PeerAsns_Delete" operation, for the dependent resources, they will have to be created separately.

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
await peerAsn.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
