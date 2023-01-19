using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Peering;
using Azure.ResourceManager.Peering.Models;
using Azure.ResourceManager.Resources;

// Generated from example definition: specification/peering/resource-manager/Microsoft.Peering/stable/2022-10-01/examples/GetPeeringReceivedRoutes.json
// this example is just showing the usage of "ReceivedRoutes_ListByPeering" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this PeeringResource created on azure
// for more information of creating PeeringResource, please refer to the document of PeeringResource
string subscriptionId = "subId";
string resourceGroupName = "rgName";
string peeringName = "peeringName";
ResourceIdentifier peeringResourceId = PeeringResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, peeringName);
PeeringResource peering = client.GetPeeringResource(peeringResourceId);

// invoke the operation and iterate over the result
string prefix = "1.1.1.0/24";
string asPath = "123 456";
string originAsValidationState = "Valid";
string rpkiValidationState = "Valid";
await foreach (PeeringReceivedRoute item in peering.GetReceivedRoutesAsync(prefix: prefix, asPath: asPath, originAsValidationState: originAsValidationState, rpkiValidationState: rpkiValidationState))
{
    Console.WriteLine($"Succeeded: {item}");
}

Console.WriteLine($"Succeeded");
