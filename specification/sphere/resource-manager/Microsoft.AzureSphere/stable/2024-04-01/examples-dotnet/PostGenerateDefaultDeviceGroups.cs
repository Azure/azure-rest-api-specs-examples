using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Sphere;
using Azure.ResourceManager.Sphere.Models;

// Generated from example definition: specification/sphere/resource-manager/Microsoft.AzureSphere/stable/2024-04-01/examples/PostGenerateDefaultDeviceGroups.json
// this example is just showing the usage of "Products_GenerateDefaultDeviceGroups" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SphereProductResource created on azure
// for more information of creating SphereProductResource, please refer to the document of SphereProductResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "MyResourceGroup1";
string catalogName = "MyCatalog1";
string productName = "MyProduct1";
ResourceIdentifier sphereProductResourceId = SphereProductResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, catalogName, productName);
SphereProductResource sphereProduct = client.GetSphereProductResource(sphereProductResourceId);

// invoke the operation and iterate over the result
await foreach (SphereDeviceGroupResource item in sphereProduct.GenerateDefaultDeviceGroupsAsync())
{
    // the variable item is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    SphereDeviceGroupData resourceData = item.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}

Console.WriteLine($"Succeeded");
