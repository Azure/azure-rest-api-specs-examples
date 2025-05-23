using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Peering;

// Generated from example definition: specification/peering/resource-manager/Microsoft.Peering/stable/2022-10-01/examples/GetRegisteredPrefix.json
// this example is just showing the usage of "RegisteredPrefixes_Get" operation, for the dependent resources, they will have to be created separately.

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

// get the collection of this PeeringRegisteredPrefixResource
PeeringRegisteredPrefixCollection collection = peering.GetPeeringRegisteredPrefixes();

// invoke the operation
string registeredPrefixName = "registeredPrefixName";
NullableResponse<PeeringRegisteredPrefixResource> response = await collection.GetIfExistsAsync(registeredPrefixName);
PeeringRegisteredPrefixResource result = response.HasValue ? response.Value : null;

if (result == null)
{
    Console.WriteLine("Succeeded with null as result");
}
else
{
    // the variable result is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    PeeringRegisteredPrefixData resourceData = result.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}
