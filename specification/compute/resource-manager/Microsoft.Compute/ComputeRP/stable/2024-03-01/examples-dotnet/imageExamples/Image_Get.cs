using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Compute.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.Compute;

// Generated from example definition: specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2024-03-01/examples/imageExamples/Image_Get.json
// this example is just showing the usage of "Images_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this DiskImageResource created on azure
// for more information of creating DiskImageResource, please refer to the document of DiskImageResource
string subscriptionId = "{subscription-id}";
string resourceGroupName = "myResourceGroup";
string imageName = "myImage";
ResourceIdentifier diskImageResourceId = DiskImageResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, imageName);
DiskImageResource diskImage = client.GetDiskImageResource(diskImageResourceId);

// invoke the operation
DiskImageResource result = await diskImage.GetAsync();

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
DiskImageData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
