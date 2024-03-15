using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Sphere;

// Generated from example definition: specification/sphere/resource-manager/Microsoft.AzureSphere/stable/2024-04-01/examples/DeleteImage.json
// this example is just showing the usage of "Images_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SphereImageResource created on azure
// for more information of creating SphereImageResource, please refer to the document of SphereImageResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "MyResourceGroup1";
string catalogName = "MyCatalog1";
string imageName = "00000000-0000-0000-0000-000000000000";
ResourceIdentifier sphereImageResourceId = SphereImageResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, catalogName, imageName);
SphereImageResource sphereImage = client.GetSphereImageResource(sphereImageResourceId);

// invoke the operation
await sphereImage.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
