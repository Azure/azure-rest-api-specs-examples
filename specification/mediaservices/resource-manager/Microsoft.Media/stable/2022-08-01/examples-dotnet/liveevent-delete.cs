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

// Generated from example definition: specification/mediaservices/resource-manager/Microsoft.Media/stable/2022-08-01/examples/liveevent-delete.json
// this example is just showing the usage of "LiveEvents_Delete" operation, for the dependent resources, they will have to be created separately.

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
await mediaLiveEvent.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
