using System;
using System.Threading.Tasks;
using System.Xml;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Media;

// Generated from example definition: specification/mediaservices/resource-manager/Microsoft.Media/Streaming/stable/2022-08-01/examples/liveoutput-list-all.json
// this example is just showing the usage of "LiveOutputs_List" operation, for the dependent resources, they will have to be created separately.

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

// invoke the operation and iterate over the result
await foreach (MediaLiveOutputResource item in collection.GetAllAsync())
{
    // the variable item is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    MediaLiveOutputData resourceData = item.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}

Console.WriteLine($"Succeeded");
