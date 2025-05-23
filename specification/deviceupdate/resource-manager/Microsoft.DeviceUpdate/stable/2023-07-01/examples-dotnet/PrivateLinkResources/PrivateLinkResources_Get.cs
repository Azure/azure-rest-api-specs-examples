using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.DeviceUpdate;

// Generated from example definition: specification/deviceupdate/resource-manager/Microsoft.DeviceUpdate/stable/2023-07-01/examples/PrivateLinkResources/PrivateLinkResources_Get.json
// this example is just showing the usage of "PrivateLinkResources_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this DeviceUpdatePrivateLinkResource created on azure
// for more information of creating DeviceUpdatePrivateLinkResource, please refer to the document of DeviceUpdatePrivateLinkResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "test-rg";
string accountName = "contoso";
string groupId = "adu";
ResourceIdentifier deviceUpdatePrivateLinkResourceId = DeviceUpdatePrivateLinkResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, accountName, groupId);
DeviceUpdatePrivateLinkResource deviceUpdatePrivateLink = client.GetDeviceUpdatePrivateLinkResource(deviceUpdatePrivateLinkResourceId);

// invoke the operation
DeviceUpdatePrivateLinkResource result = await deviceUpdatePrivateLink.GetAsync();

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
DeviceUpdatePrivateLinkData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
