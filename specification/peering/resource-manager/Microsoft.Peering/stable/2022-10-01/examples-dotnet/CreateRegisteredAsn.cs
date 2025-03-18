using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Peering;

// Generated from example definition: specification/peering/resource-manager/Microsoft.Peering/stable/2022-10-01/examples/CreateRegisteredAsn.json
// this example is just showing the usage of "RegisteredAsns_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this PeeringRegisteredAsnResource created on azure
// for more information of creating PeeringRegisteredAsnResource, please refer to the document of PeeringRegisteredAsnResource
string subscriptionId = "subId";
string resourceGroupName = "rgName";
string peeringName = "peeringName";
string registeredAsnName = "registeredAsnName";
ResourceIdentifier peeringRegisteredAsnResourceId = PeeringRegisteredAsnResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, peeringName, registeredAsnName);
PeeringRegisteredAsnResource peeringRegisteredAsn = client.GetPeeringRegisteredAsnResource(peeringRegisteredAsnResourceId);

// invoke the operation
PeeringRegisteredAsnData data = new PeeringRegisteredAsnData
{
    Asn = 65000,
};
ArmOperation<PeeringRegisteredAsnResource> lro = await peeringRegisteredAsn.UpdateAsync(WaitUntil.Completed, data);
PeeringRegisteredAsnResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
PeeringRegisteredAsnData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
