using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Media;
using Azure.ResourceManager.Media.Models;

// Generated from example definition: specification/mediaservices/resource-manager/Microsoft.Media/Metadata/stable/2023-01-01/examples/assets-list-sas-urls.json
// this example is just showing the usage of "Assets_ListContainerSas" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this MediaAssetResource created on azure
// for more information of creating MediaAssetResource, please refer to the document of MediaAssetResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "contosorg";
string accountName = "contosomedia";
string assetName = "ClimbingMountBaker";
ResourceIdentifier mediaAssetResourceId = MediaAssetResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, accountName, assetName);
MediaAssetResource mediaAsset = client.GetMediaAssetResource(mediaAssetResourceId);

// invoke the operation and iterate over the result
MediaAssetStorageContainerSasContent content = new MediaAssetStorageContainerSasContent()
{
    Permissions = MediaAssetContainerPermission.ReadWrite,
    ExpireOn = DateTimeOffset.Parse("2018-01-01T10:00:00.007Z"),
};
await foreach (Uri item in mediaAsset.GetStorageContainerUrisAsync(content))
{
    Console.WriteLine($"Succeeded: {item}");
}

Console.WriteLine($"Succeeded");
