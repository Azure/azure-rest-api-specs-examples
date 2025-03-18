using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Media.Models;
using Azure.ResourceManager.Media;

// Generated from example definition: specification/mediaservices/resource-manager/Microsoft.Media/Streaming/stable/2022-08-01/examples/streamingendpoint-start.json
// this example is just showing the usage of "StreamingEndpoints_Start" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this StreamingEndpointResource created on azure
// for more information of creating StreamingEndpointResource, please refer to the document of StreamingEndpointResource
string subscriptionId = "0a6ec948-5a62-437d-b9df-934dc7c1b722";
string resourceGroupName = "mediaresources";
string accountName = "slitestmedia10";
string streamingEndpointName = "myStreamingEndpoint1";
ResourceIdentifier streamingEndpointResourceId = StreamingEndpointResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, accountName, streamingEndpointName);
StreamingEndpointResource streamingEndpoint = client.GetStreamingEndpointResource(streamingEndpointResourceId);

// invoke the operation
await streamingEndpoint.StartAsync(WaitUntil.Completed);

Console.WriteLine("Succeeded");
