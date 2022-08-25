using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Compute.Models;
using Azure.ResourceManager.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.Compute;
// Generated from example definition: specification/compute/resource-manager/Microsoft.Compute/DiskRP/stable/2022-03-02/examples/diskEncryptionSetExamples/DiskEncryptionSet_ListAssociatedResources.json
// this example is just showing the usage of "DiskEncryptionSets_ListAssociatedResources" operation, for the dependent resources, they will have to be created separately.
            
// authenticate your client
ArmClient client = new ArmClient(new DefaultAzureCredential());
            
// this example assumes you already have this DiskEncryptionSetResource created on azure
// for more information of creating DiskEncryptionSetResource, please refer to the document of DiskEncryptionSetResource
string subscriptionId = "{subscription-id}";
string resourceGroupName = "myResourceGroup";
string diskEncryptionSetName = "myDiskEncryptionSet";
ResourceIdentifier diskEncryptionSetResourceId = DiskEncryptionSetResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, diskEncryptionSetName);
DiskEncryptionSetResource diskEncryptionSet = client.GetDiskEncryptionSetResource(diskEncryptionSetResourceId);
            
// invoke the operation and iterate over the result
await foreach (string item in diskEncryptionSet.GetAssociatedResourcesAsync())
{
    Console.WriteLine($"Succeeded: {item}");
}
            
Console.WriteLine($"Succeeded");
