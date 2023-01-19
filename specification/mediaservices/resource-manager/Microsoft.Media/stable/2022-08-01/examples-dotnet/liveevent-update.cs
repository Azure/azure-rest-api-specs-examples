using System;
using System.Net;
using System.Threading.Tasks;
using System.Xml;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Media;
using Azure.ResourceManager.Media.Models;

// Generated from example definition: specification/mediaservices/resource-manager/Microsoft.Media/stable/2022-08-01/examples/liveevent-update.json
// this example is just showing the usage of "LiveEvents_Update" operation, for the dependent resources, they will have to be created separately.

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

// invoke the operation
MediaLiveEventData data = new MediaLiveEventData(new AzureLocation("West US"))
{
    Description = "test event updated",
    Input = new LiveEventInput(LiveEventInputProtocol.FragmentedMp4)
    {
        IPAllowedIPs =
        {
        new IPRange()
        {
        Name = "AllowOne",
        Address = IPAddress.Parse("192.1.1.0"),
        }
        },
        KeyFrameIntervalDuration = XmlConvert.ToTimeSpan("PT6S"),
    },
    Preview = new LiveEventPreview()
    {
        IPAllowedIPs =
        {
        new IPRange()
        {
        Name = "AllowOne",
        Address = IPAddress.Parse("192.1.1.0"),
        }
        },
    },
    Tags =
    {
    ["tag1"] = "value1",
    ["tag2"] = "value2",
    ["tag3"] = "value3",
    },
};
ArmOperation<MediaLiveEventResource> lro = await mediaLiveEvent.UpdateAsync(WaitUntil.Completed, data);
MediaLiveEventResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
MediaLiveEventData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
