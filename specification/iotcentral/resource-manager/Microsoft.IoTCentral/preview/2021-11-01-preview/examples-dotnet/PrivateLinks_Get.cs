using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.IotCentral;

// Generated from example definition: specification/iotcentral/resource-manager/Microsoft.IoTCentral/preview/2021-11-01-preview/examples/PrivateLinks_Get.json
// this example is just showing the usage of "PrivateLinks_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this IotCentralPrivateLinkResource created on azure
// for more information of creating IotCentralPrivateLinkResource, please refer to the document of IotCentralPrivateLinkResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "resRg";
string resourceName = "myIoTCentralApp";
string groupId = "iotApp";
ResourceIdentifier iotCentralPrivateLinkResourceId = IotCentralPrivateLinkResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, resourceName, groupId);
IotCentralPrivateLinkResource iotCentralPrivateLinkResource = client.GetIotCentralPrivateLinkResource(iotCentralPrivateLinkResourceId);

// invoke the operation
IotCentralPrivateLinkResource result = await iotCentralPrivateLinkResource.GetAsync();

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
IotCentralPrivateLinkResourceData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
