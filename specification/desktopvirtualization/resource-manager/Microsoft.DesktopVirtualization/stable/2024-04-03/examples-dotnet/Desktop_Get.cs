using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.DesktopVirtualization;

// Generated from example definition: specification/desktopvirtualization/resource-manager/Microsoft.DesktopVirtualization/stable/2024-04-03/examples/Desktop_Get.json
// this example is just showing the usage of "Desktops_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this VirtualApplicationGroupResource created on azure
// for more information of creating VirtualApplicationGroupResource, please refer to the document of VirtualApplicationGroupResource
string subscriptionId = "daefabc0-95b4-48b3-b645-8a753a63c4fa";
string resourceGroupName = "resourceGroup1";
string applicationGroupName = "applicationGroup1";
ResourceIdentifier virtualApplicationGroupResourceId = VirtualApplicationGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, applicationGroupName);
VirtualApplicationGroupResource virtualApplicationGroup = client.GetVirtualApplicationGroupResource(virtualApplicationGroupResourceId);

// get the collection of this VirtualDesktopResource
VirtualDesktopCollection collection = virtualApplicationGroup.GetVirtualDesktops();

// invoke the operation
string desktopName = "SessionDesktop";
NullableResponse<VirtualDesktopResource> response = await collection.GetIfExistsAsync(desktopName);
VirtualDesktopResource result = response.HasValue ? response.Value : null;

if (result == null)
{
    Console.WriteLine("Succeeded with null as result");
}
else
{
    // the variable result is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    VirtualDesktopData resourceData = result.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}
