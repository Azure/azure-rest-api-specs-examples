using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Media;
using Azure.ResourceManager.Media.Models;

// Generated from example definition: specification/mediaservices/resource-manager/Microsoft.Media/Metadata/stable/2022-08-01/examples/streaming-locators-delete.json
// this example is just showing the usage of "StreamingLocators_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this StreamingLocatorResource created on azure
// for more information of creating StreamingLocatorResource, please refer to the document of StreamingLocatorResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "contoso";
string accountName = "contosomedia";
string streamingLocatorName = "clearStreamingLocator";
ResourceIdentifier streamingLocatorResourceId = StreamingLocatorResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, accountName, streamingLocatorName);
StreamingLocatorResource streamingLocator = client.GetStreamingLocatorResource(streamingLocatorResourceId);

// invoke the operation
await streamingLocator.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
