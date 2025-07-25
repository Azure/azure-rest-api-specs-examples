using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.NetworkCloud.Models;
using Azure.ResourceManager.NetworkCloud;

// Generated from example definition: specification/networkcloud/resource-manager/Microsoft.NetworkCloud/stable/2025-02-01/examples/Volumes_Patch.json
// this example is just showing the usage of "Volumes_Update" operation, for the dependent resources, they will have to be created separately.

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
NetworkCloudVolumePatch patch = new NetworkCloudVolumePatch
{
    Tags =
    {
    ["key1"] = "myvalue1",
    ["key2"] = "myvalue2"
    },
};
NetworkCloudVolumeResource result = await networkCloudVolume.UpdateAsync(patch);

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
NetworkCloudVolumeData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
