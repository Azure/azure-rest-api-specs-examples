using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Peering;

// Generated from example definition: specification/peering/resource-manager/Microsoft.Peering/stable/2022-10-01/examples/CreateRegisteredPrefix.json
// this example is just showing the usage of "RegisteredPrefixes_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

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
PeeringRegisteredPrefixData data = new PeeringRegisteredPrefixData()
{
    Prefix = "10.22.20.0/24",
};
ArmOperation<PeeringRegisteredPrefixResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, registeredPrefixName, data);
PeeringRegisteredPrefixResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
PeeringRegisteredPrefixData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
