using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Media;
using Azure.ResourceManager.Media.Models;

// Generated from example definition: specification/mediaservices/resource-manager/Microsoft.Media/Metadata/stable/2022-08-01/examples/accountFilters-create.json
// this example is just showing the usage of "AccountFilters_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this MediaServicesAccountResource created on azure
// for more information of creating MediaServicesAccountResource, please refer to the document of MediaServicesAccountResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "contoso";
string accountName = "contosomedia";
ResourceIdentifier mediaServicesAccountResourceId = MediaServicesAccountResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, accountName);
MediaServicesAccountResource mediaServicesAccount = client.GetMediaServicesAccountResource(mediaServicesAccountResourceId);

// get the collection of this MediaServicesAccountFilterResource
MediaServicesAccountFilterCollection collection = mediaServicesAccount.GetMediaServicesAccountFilters();

// invoke the operation
string filterName = "newAccountFilter";
MediaServicesAccountFilterData data = new MediaServicesAccountFilterData()
{
    PresentationTimeRange = new PresentationTimeRange()
    {
        StartTimestamp = 0,
        EndTimestamp = 170000000,
        PresentationWindowDuration = 9223372036854775000,
        LiveBackoffDuration = 0,
        Timescale = 10000000,
        ForceEndTimestamp = false,
    },
    FirstQualityBitrate = 128000,
    Tracks =
    {
    new FilterTrackSelection(new FilterTrackPropertyCondition[]
    {
    new FilterTrackPropertyCondition(FilterTrackPropertyType.Type,"Audio",FilterTrackPropertyCompareOperation.Equal),new FilterTrackPropertyCondition(FilterTrackPropertyType.Language,"en",FilterTrackPropertyCompareOperation.NotEqual),new FilterTrackPropertyCondition(FilterTrackPropertyType.FourCC,"EC-3",FilterTrackPropertyCompareOperation.NotEqual)
    }),new FilterTrackSelection(new FilterTrackPropertyCondition[]
    {
    new FilterTrackPropertyCondition(FilterTrackPropertyType.Type,"Video",FilterTrackPropertyCompareOperation.Equal),new FilterTrackPropertyCondition(FilterTrackPropertyType.Bitrate,"3000000-5000000",FilterTrackPropertyCompareOperation.Equal)
    })
    },
};
ArmOperation<MediaServicesAccountFilterResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, filterName, data);
MediaServicesAccountFilterResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
MediaServicesAccountFilterData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
