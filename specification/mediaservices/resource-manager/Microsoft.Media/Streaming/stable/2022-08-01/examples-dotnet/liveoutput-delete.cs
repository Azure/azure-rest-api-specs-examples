using System;
using System.Threading.Tasks;
using System.Xml;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Media;

// Generated from example definition: specification/mediaservices/resource-manager/Microsoft.Media/Streaming/stable/2022-08-01/examples/liveoutput-delete.json
// this example is just showing the usage of "LiveOutputs_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this MediaLiveOutputResource created on azure
// for more information of creating MediaLiveOutputResource, please refer to the document of MediaLiveOutputResource
string subscriptionId = "0a6ec948-5a62-437d-b9df-934dc7c1b722";
string resourceGroupName = "mediaresources";
string accountName = "slitestmedia10";
string liveEventName = "myLiveEvent1";
string liveOutputName = "myLiveOutput1";
ResourceIdentifier mediaLiveOutputResourceId = MediaLiveOutputResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, accountName, liveEventName, liveOutputName);
MediaLiveOutputResource mediaLiveOutput = client.GetMediaLiveOutputResource(mediaLiveOutputResourceId);

// invoke the operation
await mediaLiveOutput.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
