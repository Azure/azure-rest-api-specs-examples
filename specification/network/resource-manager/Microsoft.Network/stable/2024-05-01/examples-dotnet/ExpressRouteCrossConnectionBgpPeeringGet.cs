using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Network.Models;
using Azure.ResourceManager.Network;

// Generated from example definition: specification/network/resource-manager/Microsoft.Network/stable/2024-05-01/examples/ExpressRouteCrossConnectionBgpPeeringGet.json
// this example is just showing the usage of "ExpressRouteCrossConnectionPeerings_Get" operation, for the dependent resources, they will have to be created separately.

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
ExpressRouteCrossConnectionPeeringResource result = await expressRouteCrossConnectionPeering.GetAsync();

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
ExpressRouteCrossConnectionPeeringData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
