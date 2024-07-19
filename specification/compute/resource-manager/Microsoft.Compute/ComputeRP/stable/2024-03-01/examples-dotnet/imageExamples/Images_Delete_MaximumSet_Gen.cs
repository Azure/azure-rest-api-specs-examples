using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Compute.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.Compute;

// Generated from example definition: specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2024-03-01/examples/imageExamples/Images_Delete_MaximumSet_Gen.json
// this example is just showing the usage of "Images_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this DiskImageResource created on azure
// for more information of creating DiskImageResource, please refer to the document of DiskImageResource
string subscriptionId = "{subscription-id}";
string resourceGroupName = "rgcompute";
string imageName = "aaaaaaaaaaaaaaaaaaaaaaaaaaa";
ResourceIdentifier diskImageResourceId = DiskImageResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, imageName);
DiskImageResource diskImage = client.GetDiskImageResource(diskImageResourceId);

// invoke the operation
await diskImage.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
