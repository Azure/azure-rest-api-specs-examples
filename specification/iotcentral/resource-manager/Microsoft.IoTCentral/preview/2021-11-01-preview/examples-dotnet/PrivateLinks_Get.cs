using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.IotCentral;

// Generated from example definition: specification/iotcentral/resource-manager/Microsoft.IoTCentral/preview/2021-11-01-preview/examples/PrivateLinks_Get.json
// this example is just showing the usage of "PrivateLinks_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this IotCentralAppResource created on azure
// for more information of creating IotCentralAppResource, please refer to the document of IotCentralAppResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "resRg";
string resourceName = "myIoTCentralApp";
ResourceIdentifier iotCentralAppResourceId = IotCentralAppResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, resourceName);
IotCentralAppResource iotCentralApp = client.GetIotCentralAppResource(iotCentralAppResourceId);

// get the collection of this IotCentralPrivateLinkResource
IotCentralPrivateLinkResourceCollection collection = iotCentralApp.GetIotCentralPrivateLinkResources();

// invoke the operation
string groupId = "iotApp";
bool result = await collection.ExistsAsync(groupId);

Console.WriteLine($"Succeeded: {result}");
