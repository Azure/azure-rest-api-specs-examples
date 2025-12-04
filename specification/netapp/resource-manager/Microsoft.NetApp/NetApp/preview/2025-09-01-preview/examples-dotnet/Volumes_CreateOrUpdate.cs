using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.NetApp.Models;
using Azure.ResourceManager.NetApp;

// Generated from example definition: specification/netapp/resource-manager/Microsoft.NetApp/NetApp/preview/2025-09-01-preview/examples/Volumes_CreateOrUpdate.json
// this example is just showing the usage of "Volumes_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this CapacityPoolResource created on azure
// for more information of creating CapacityPoolResource, please refer to the document of CapacityPoolResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "myRG";
string accountName = "account1";
string poolName = "pool1";
ResourceIdentifier capacityPoolResourceId = CapacityPoolResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, accountName, poolName);
CapacityPoolResource capacityPool = client.GetCapacityPoolResource(capacityPoolResourceId);

// get the collection of this NetAppVolumeResource
NetAppVolumeCollection collection = capacityPool.GetNetAppVolumes();

// invoke the operation
string volumeName = "volume1";
NetAppVolumeData data = new NetAppVolumeData(new AzureLocation("eastus"), "my-unique-file-path", 107374182400L, new ResourceIdentifier("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myRP/providers/Microsoft.Network/virtualNetworks/testvnet3/subnets/testsubnet3"))
{
    ServiceLevel = NetAppFileServiceLevel.Premium,
};
ArmOperation<NetAppVolumeResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, volumeName, data);
NetAppVolumeResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
NetAppVolumeData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
