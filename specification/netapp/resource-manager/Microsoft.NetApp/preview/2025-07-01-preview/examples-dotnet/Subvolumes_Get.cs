using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.NetApp.Models;
using Azure.ResourceManager.NetApp;

// Generated from example definition: specification/netapp/resource-manager/Microsoft.NetApp/preview/2025-07-01-preview/examples/Subvolumes_Get.json
// this example is just showing the usage of "Subvolumes_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this NetAppSubvolumeInfoResource created on azure
// for more information of creating NetAppSubvolumeInfoResource, please refer to the document of NetAppSubvolumeInfoResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "myRG";
string accountName = "account1";
string poolName = "pool1";
string volumeName = "volume1";
string subvolumeName = "subvolume1";
ResourceIdentifier netAppSubvolumeInfoResourceId = NetAppSubvolumeInfoResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, accountName, poolName, volumeName, subvolumeName);
NetAppSubvolumeInfoResource netAppSubvolumeInfo = client.GetNetAppSubvolumeInfoResource(netAppSubvolumeInfoResourceId);

// invoke the operation
NetAppSubvolumeInfoResource result = await netAppSubvolumeInfo.GetAsync();

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
NetAppSubvolumeInfoData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
