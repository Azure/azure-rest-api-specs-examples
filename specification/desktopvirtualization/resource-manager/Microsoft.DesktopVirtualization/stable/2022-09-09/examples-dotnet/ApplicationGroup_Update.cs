using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.DesktopVirtualization;
using Azure.ResourceManager.DesktopVirtualization.Models;
using Azure.ResourceManager.Resources;

// Generated from example definition: specification/desktopvirtualization/resource-manager/Microsoft.DesktopVirtualization/stable/2022-09-09/examples/ApplicationGroup_Update.json
// this example is just showing the usage of "ApplicationGroups_Update" operation, for the dependent resources, they will have to be created separately.

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

// invoke the operation
VirtualApplicationGroupPatch patch = new VirtualApplicationGroupPatch()
{
    Tags =
    {
    ["tag1"] = "value1",
    ["tag2"] = "value2",
    },
    Description = "des1",
    FriendlyName = "friendly",
};
VirtualApplicationGroupResource result = await virtualApplicationGroup.UpdateAsync(patch);

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
VirtualApplicationGroupData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
