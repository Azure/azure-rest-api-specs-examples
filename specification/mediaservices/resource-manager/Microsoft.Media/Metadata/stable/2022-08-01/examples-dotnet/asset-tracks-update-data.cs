using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Media;
using Azure.ResourceManager.Media.Models;

// Generated from example definition: specification/mediaservices/resource-manager/Microsoft.Media/Metadata/stable/2022-08-01/examples/asset-tracks-update-data.json
// this example is just showing the usage of "Tracks_UpdateTrackData" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this MediaAssetTrackResource created on azure
// for more information of creating MediaAssetTrackResource, please refer to the document of MediaAssetTrackResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "contoso";
string accountName = "contosomedia";
string assetName = "ClimbingMountRainer";
string trackName = "text2";
ResourceIdentifier mediaAssetTrackResourceId = MediaAssetTrackResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, accountName, assetName, trackName);
MediaAssetTrackResource mediaAssetTrack = client.GetMediaAssetTrackResource(mediaAssetTrackResourceId);

// invoke the operation
await mediaAssetTrack.UpdateTrackDataAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
