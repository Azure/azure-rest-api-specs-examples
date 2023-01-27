using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Media;
using Azure.ResourceManager.Media.Models;

// Generated from example definition: specification/mediaservices/resource-manager/Microsoft.Media/Metadata/stable/2022-08-01/examples/asset-tracks-list-all.json
// this example is just showing the usage of "Tracks_List" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this MediaAssetResource created on azure
// for more information of creating MediaAssetResource, please refer to the document of MediaAssetResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "contoso";
string accountName = "contosomedia";
string assetName = "ClimbingMountRainer";
ResourceIdentifier mediaAssetResourceId = MediaAssetResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, accountName, assetName);
MediaAssetResource mediaAsset = client.GetMediaAssetResource(mediaAssetResourceId);

// get the collection of this MediaAssetTrackResource
MediaAssetTrackCollection collection = mediaAsset.GetMediaAssetTracks();

// invoke the operation and iterate over the result
await foreach (MediaAssetTrackResource item in collection.GetAllAsync())
{
    // the variable item is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    MediaAssetTrackData resourceData = item.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}

Console.WriteLine($"Succeeded");
