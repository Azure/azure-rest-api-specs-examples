using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Sphere.Models;
using Azure.ResourceManager.Sphere;

// Generated from example definition: specification/sphere/resource-manager/Microsoft.AzureSphere/stable/2024-04-01/examples/DeleteDeviceGroup.json
// this example is just showing the usage of "DeviceGroups_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SphereDeviceGroupResource created on azure
// for more information of creating SphereDeviceGroupResource, please refer to the document of SphereDeviceGroupResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "MyResourceGroup1";
string catalogName = "MyCatalog1";
string productName = "MyProduct1";
string deviceGroupName = "MyDeviceGroup1";
ResourceIdentifier sphereDeviceGroupResourceId = SphereDeviceGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, catalogName, productName, deviceGroupName);
SphereDeviceGroupResource sphereDeviceGroup = client.GetSphereDeviceGroupResource(sphereDeviceGroupResourceId);

// invoke the operation
await sphereDeviceGroup.DeleteAsync(WaitUntil.Completed);

Console.WriteLine("Succeeded");
