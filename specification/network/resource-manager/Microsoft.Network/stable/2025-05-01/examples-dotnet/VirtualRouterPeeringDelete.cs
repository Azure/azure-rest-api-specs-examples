using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Network;

// Generated from example definition: specification/network/resource-manager/Microsoft.Network/stable/2025-05-01/examples/VirtualRouterPeeringDelete.json
// this example is just showing the usage of "VirtualRouterPeerings_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this VirtualRouterPeeringResource created on azure
// for more information of creating VirtualRouterPeeringResource, please refer to the document of VirtualRouterPeeringResource
string subscriptionId = "subid";
string resourceGroupName = "rg1";
string virtualRouterName = "virtualRouter";
string peeringName = "peering1";
ResourceIdentifier virtualRouterPeeringResourceId = VirtualRouterPeeringResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, virtualRouterName, peeringName);
VirtualRouterPeeringResource virtualRouterPeering = client.GetVirtualRouterPeeringResource(virtualRouterPeeringResourceId);

// invoke the operation
await virtualRouterPeering.DeleteAsync(WaitUntil.Completed);

Console.WriteLine("Succeeded");
