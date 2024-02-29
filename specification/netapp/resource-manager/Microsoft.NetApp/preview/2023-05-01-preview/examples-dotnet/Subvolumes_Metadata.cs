using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.NetApp;
using Azure.ResourceManager.NetApp.Models;

// Generated from example definition: specification/netapp/resource-manager/Microsoft.NetApp/preview/2023-05-01-preview/examples/Subvolumes_Metadata.json
// this example is just showing the usage of "Subvolumes_GetMetadata" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this NetAppSubvolumeInfoResource created on azure
// for more information of creating NetAppSubvolumeInfoResource, please refer to the document of NetAppSubvolumeInfoResource
string subscriptionId = "D633CC2E-722B-4AE1-B636-BBD9E4C60ED9";
string resourceGroupName = "myRG";
string accountName = "account1";
string poolName = "pool1";
string volumeName = "volume1";
string subvolumeName = "subvolume1";
ResourceIdentifier netAppSubvolumeInfoResourceId = NetAppSubvolumeInfoResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, accountName, poolName, volumeName, subvolumeName);
NetAppSubvolumeInfoResource netAppSubvolumeInfo = client.GetNetAppSubvolumeInfoResource(netAppSubvolumeInfoResourceId);

// invoke the operation
ArmOperation<NetAppSubvolumeMetadata> lro = await netAppSubvolumeInfo.GetMetadataAsync(WaitUntil.Completed);
NetAppSubvolumeMetadata result = lro.Value;

Console.WriteLine($"Succeeded: {result}");
