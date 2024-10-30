using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.NetApp.Models;
using Azure.ResourceManager.NetApp;

// Generated from example definition: specification/netapp/resource-manager/Microsoft.NetApp/stable/2024-07-01/examples/VolumeGroups_Delete.json
// this example is just showing the usage of "VolumeGroups_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this NetAppVolumeGroupResource created on azure
// for more information of creating NetAppVolumeGroupResource, please refer to the document of NetAppVolumeGroupResource
string subscriptionId = "D633CC2E-722B-4AE1-B636-BBD9E4C60ED9";
string resourceGroupName = "myRG";
string accountName = "account1";
string volumeGroupName = "group1";
ResourceIdentifier netAppVolumeGroupResourceId = NetAppVolumeGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, accountName, volumeGroupName);
NetAppVolumeGroupResource netAppVolumeGroup = client.GetNetAppVolumeGroupResource(netAppVolumeGroupResourceId);

// invoke the operation
await netAppVolumeGroup.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
