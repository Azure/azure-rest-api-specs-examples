using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Peering;
using Azure.ResourceManager.Resources;

// Generated from example definition: specification/peering/resource-manager/Microsoft.Peering/stable/2022-10-01/examples/CreatePeeringService.json
// this example is just showing the usage of "PeeringServices_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ResourceGroupResource created on azure
// for more information of creating ResourceGroupResource, please refer to the document of ResourceGroupResource
string subscriptionId = "subId";
string resourceGroupName = "rgName";
ResourceIdentifier resourceGroupResourceId = ResourceGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName);
ResourceGroupResource resourceGroupResource = client.GetResourceGroupResource(resourceGroupResourceId);

// get the collection of this PeeringServiceResource
PeeringServiceCollection collection = resourceGroupResource.GetPeeringServices();

// invoke the operation
string peeringServiceName = "peeringServiceName";
PeeringServiceData data = new PeeringServiceData(new AzureLocation("eastus"))
{
    PeeringServiceLocation = "state1",
    PeeringServiceProvider = "serviceProvider1",
    ProviderPrimaryPeeringLocation = "peeringLocation1",
    ProviderBackupPeeringLocation = "peeringLocation2",
};
ArmOperation<PeeringServiceResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, peeringServiceName, data);
PeeringServiceResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
PeeringServiceData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
