using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Peering;

// Generated from example definition: specification/peering/resource-manager/Microsoft.Peering/stable/2022-10-01/examples/CreatePeeringServicePrefix.json
// this example is just showing the usage of "Prefixes_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this PeeringServiceResource created on azure
// for more information of creating PeeringServiceResource, please refer to the document of PeeringServiceResource
string subscriptionId = "subId";
string resourceGroupName = "rgName";
string peeringServiceName = "peeringServiceName";
ResourceIdentifier peeringServiceResourceId = PeeringServiceResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, peeringServiceName);
PeeringServiceResource peeringService = client.GetPeeringServiceResource(peeringServiceResourceId);

// get the collection of this PeeringServicePrefixResource
PeeringServicePrefixCollection collection = peeringService.GetPeeringServicePrefixes();

// invoke the operation
string prefixName = "peeringServicePrefixName";
PeeringServicePrefixData data = new PeeringServicePrefixData
{
    Prefix = "192.168.1.0/24",
    PeeringServicePrefixKey = "00000000-0000-0000-0000-000000000000",
};
ArmOperation<PeeringServicePrefixResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, prefixName, data);
PeeringServicePrefixResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
PeeringServicePrefixData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
