using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Avs.Models;
using Azure.ResourceManager.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.Avs;

// Generated from example definition: 2024-09-01/PrivateClouds_CreateOrUpdate.json
// this example is just showing the usage of "PrivateCloud_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ResourceGroupResource created on azure
// for more information of creating ResourceGroupResource, please refer to the document of ResourceGroupResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "group1";
ResourceIdentifier resourceGroupResourceId = ResourceGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName);
ResourceGroupResource resourceGroupResource = client.GetResourceGroupResource(resourceGroupResourceId);

// get the collection of this AvsPrivateCloudResource
AvsPrivateCloudCollection collection = resourceGroupResource.GetAvsPrivateClouds();

// invoke the operation
string privateCloudName = "cloud1";
AvsPrivateCloudData data = new AvsPrivateCloudData(new AzureLocation("eastus2"), new AvsSku("AV36"))
{
    Identity = new ManagedServiceIdentity("SystemAssigned"),
    Tags = { },
};
ArmOperation<AvsPrivateCloudResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, privateCloudName, data);
AvsPrivateCloudResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
AvsPrivateCloudData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
