using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Compute.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.Compute;
// Generated from example definition: specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2022-03-01/examples/imageExamples/Images_Delete_MinimumSet_Gen.json
// this example is just showing the usage of "Images_Delete" operation, for the dependent resources, they will have to be created separately.
            
// authenticate your client
ArmClient client = new ArmClient(new DefaultAzureCredential());
            
// this example assumes you already have this DiskImageResource created on azure
// for more information of creating DiskImageResource, please refer to the document of DiskImageResource
string subscriptionId = "{subscription-id}";
string resourceGroupName = "rgcompute";
string imageName = "aaaaaaaaaaaaaaaaaaaaaaaaaaaa";
ResourceIdentifier diskImageResourceId = DiskImageResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, imageName);
DiskImageResource diskImage = client.GetDiskImageResource(diskImageResourceId);
            
// invoke the operation
await diskImage.DeleteAsync(WaitUntil.Completed);
            
Console.WriteLine($"Succeeded");
