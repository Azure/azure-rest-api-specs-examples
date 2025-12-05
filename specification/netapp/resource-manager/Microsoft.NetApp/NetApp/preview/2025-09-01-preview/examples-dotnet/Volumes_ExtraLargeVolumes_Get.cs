using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.NetApp.Models;
using Azure.ResourceManager.NetApp;

// Generated from example definition: specification/netapp/resource-manager/Microsoft.NetApp/NetApp/preview/2025-09-01-preview/examples/Volumes_ExtraLargeVolumes_Get.json
// this example is just showing the usage of "Volumes_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this CapacityPoolResource created on azure
// for more information of creating CapacityPoolResource, please refer to the document of CapacityPoolResource
string subscriptionId = "D633CC2E-722B-4AE1-B636-BBD9E4C60ED9";
string resourceGroupName = "myRG";
string accountName = "account1";
string poolName = "pool1";
ResourceIdentifier capacityPoolResourceId = CapacityPoolResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, accountName, poolName);
CapacityPoolResource capacityPool = client.GetCapacityPoolResource(capacityPoolResourceId);

// get the collection of this NetAppVolumeResource
NetAppVolumeCollection collection = capacityPool.GetNetAppVolumes();

// invoke the operation
string volumeName = "volume1";
NullableResponse<NetAppVolumeResource> response = await collection.GetIfExistsAsync(volumeName);
NetAppVolumeResource result = response.HasValue ? response.Value : null;

if (result == null)
{
    Console.WriteLine("Succeeded with null as result");
}
else
{
    // the variable result is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    NetAppVolumeData resourceData = result.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}
