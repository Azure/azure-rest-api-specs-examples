using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Compute.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.Resources.Models;
using Azure.ResourceManager.Compute;
// Generated from example definition: specification/compute/resource-manager/Microsoft.Compute/DiskRP/stable/2022-03-02/examples/diskExamples/Disk_Create_WithDataAccessAuthMode.json
// this example is just showing the usage of "Disks_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.
            
// authenticate your client
ArmClient client = new ArmClient(new DefaultAzureCredential());
            
// this example assumes you already have this ResourceGroupResource created on azure
// for more information of creating ResourceGroupResource, please refer to the document of ResourceGroupResource
string subscriptionId = "{subscription-id}";
string resourceGroupName = "myResourceGroup";
ResourceIdentifier resourceGroupResourceId = ResourceGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName);
ResourceGroupResource resourceGroupResource = client.GetResourceGroupResource(resourceGroupResourceId);
            
// get the collection of this ManagedDiskResource
ManagedDiskCollection collection = resourceGroupResource.GetManagedDisks();
            
// invoke the operation
string diskName = "myDisk";
ManagedDiskData data = new ManagedDiskData(new AzureLocation("West US"))
{
    CreationData = new DiskCreationData(DiskCreateOption.Empty),
    DiskSizeGB = 200,
    DataAccessAuthMode = DataAccessAuthMode.AzureActiveDirectory,
};
ArmOperation<ManagedDiskResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, diskName, data);
ManagedDiskResource result = lro.Value;
            
// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
ManagedDiskData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
