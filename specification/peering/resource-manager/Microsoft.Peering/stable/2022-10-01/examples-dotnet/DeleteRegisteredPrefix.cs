using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Peering;

// Generated from example definition: specification/peering/resource-manager/Microsoft.Peering/stable/2022-10-01/examples/DeleteRegisteredPrefix.json
// this example is just showing the usage of "RegisteredPrefixes_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this PeeringRegisteredPrefixResource created on azure
// for more information of creating PeeringRegisteredPrefixResource, please refer to the document of PeeringRegisteredPrefixResource
string subscriptionId = "subId";
string resourceGroupName = "rgName";
string peeringName = "peeringName";
string registeredPrefixName = "registeredPrefixName";
ResourceIdentifier peeringRegisteredPrefixResourceId = PeeringRegisteredPrefixResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, peeringName, registeredPrefixName);
PeeringRegisteredPrefixResource peeringRegisteredPrefix = client.GetPeeringRegisteredPrefixResource(peeringRegisteredPrefixResourceId);

// invoke the operation
await peeringRegisteredPrefix.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
