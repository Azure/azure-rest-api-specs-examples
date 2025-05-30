using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using System.Xml;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Media;

// Generated from example definition: specification/mediaservices/resource-manager/Microsoft.Media/Streaming/stable/2022-08-01/examples/liveoutput-create.json
// this example is just showing the usage of "LiveOutputs_Create" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this MediaLiveEventResource created on azure
// for more information of creating MediaLiveEventResource, please refer to the document of MediaLiveEventResource
string subscriptionId = "0a6ec948-5a62-437d-b9df-934dc7c1b722";
string resourceGroupName = "mediaresources";
string accountName = "slitestmedia10";
string liveEventName = "myLiveEvent1";
ResourceIdentifier mediaLiveEventResourceId = MediaLiveEventResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, accountName, liveEventName);
MediaLiveEventResource mediaLiveEvent = client.GetMediaLiveEventResource(mediaLiveEventResourceId);

// get the collection of this MediaLiveOutputResource
MediaLiveOutputCollection collection = mediaLiveEvent.GetMediaLiveOutputs();

// invoke the operation
string liveOutputName = "myLiveOutput1";
MediaLiveOutputData data = new MediaLiveOutputData
{
    Description = "test live output 1",
    AssetName = "6f3264f5-a189-48b4-a29a-a40f22575212",
    ArchiveWindowLength = XmlConvert.ToTimeSpan("PT5M"),
    RewindWindowLength = XmlConvert.ToTimeSpan("PT4M"),
    ManifestName = "testmanifest",
    HlsFragmentsPerTsSegment = 5,
};
ArmOperation<MediaLiveOutputResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, liveOutputName, data);
MediaLiveOutputResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
MediaLiveOutputData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
