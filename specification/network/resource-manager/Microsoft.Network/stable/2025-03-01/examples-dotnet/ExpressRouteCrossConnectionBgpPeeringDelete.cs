using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Network.Models;
using Azure.ResourceManager.Network;

// Generated from example definition: specification/network/resource-manager/Microsoft.Network/stable/2025-03-01/examples/ExpressRouteCrossConnectionBgpPeeringDelete.json
// this example is just showing the usage of "ExpressRouteCrossConnectionPeerings_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ExpressRouteCrossConnectionPeeringResource created on azure
// for more information of creating ExpressRouteCrossConnectionPeeringResource, please refer to the document of ExpressRouteCrossConnectionPeeringResource
string subscriptionId = "subid";
string resourceGroupName = "CrossConnection-SiliconValley";
string crossConnectionName = "<circuitServiceKey>";
string peeringName = "AzurePrivatePeering";
ResourceIdentifier expressRouteCrossConnectionPeeringResourceId = ExpressRouteCrossConnectionPeeringResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, crossConnectionName, peeringName);
ExpressRouteCrossConnectionPeeringResource expressRouteCrossConnectionPeering = client.GetExpressRouteCrossConnectionPeeringResource(expressRouteCrossConnectionPeeringResourceId);

// invoke the operation
await expressRouteCrossConnectionPeering.DeleteAsync(WaitUntil.Completed);

Console.WriteLine("Succeeded");
