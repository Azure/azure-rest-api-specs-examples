using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.NetworkCloud;
using Azure.ResourceManager.NetworkCloud.Models;
using Azure.ResourceManager.Resources;

// Generated from example definition: specification/networkcloud/resource-manager/Microsoft.NetworkCloud/stable/2023-07-01/examples/Volumes_Delete.json
// this example is just showing the usage of "Volumes_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this NetworkCloudVolumeResource created on azure
// for more information of creating NetworkCloudVolumeResource, please refer to the document of NetworkCloudVolumeResource
string subscriptionId = "123e4567-e89b-12d3-a456-426655440000";
string resourceGroupName = "resourceGroupName";
string volumeName = "volumeName";
ResourceIdentifier networkCloudVolumeResourceId = NetworkCloudVolumeResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, volumeName);
NetworkCloudVolumeResource networkCloudVolume = client.GetNetworkCloudVolumeResource(networkCloudVolumeResourceId);

// invoke the operation
await networkCloudVolume.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
