using Azure;
using Azure.ResourceManager;
using System;
using System.Net;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.NetApp.Models;
using Azure.ResourceManager.NetApp;

// Generated from example definition: specification/netapp/resource-manager/Microsoft.NetApp/stable/2023-11-01/examples/Volumes_Update.json
// this example is just showing the usage of "Volumes_Update" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this NetAppVolumeResource created on azure
// for more information of creating NetAppVolumeResource, please refer to the document of NetAppVolumeResource
string subscriptionId = "D633CC2E-722B-4AE1-B636-BBD9E4C60ED9";
string resourceGroupName = "myRG";
string accountName = "account1";
string poolName = "pool1";
string volumeName = "volume1";
ResourceIdentifier netAppVolumeResourceId = NetAppVolumeResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, accountName, poolName, volumeName);
NetAppVolumeResource netAppVolume = client.GetNetAppVolumeResource(netAppVolumeResourceId);

// invoke the operation
NetAppVolumePatch patch = new NetAppVolumePatch(new AzureLocation("placeholder"));
ArmOperation<NetAppVolumeResource> lro = await netAppVolume.UpdateAsync(WaitUntil.Completed, patch);
NetAppVolumeResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
NetAppVolumeData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
