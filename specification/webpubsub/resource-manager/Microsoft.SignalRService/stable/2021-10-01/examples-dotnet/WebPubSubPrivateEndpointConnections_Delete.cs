using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.WebPubSub.Models;
using Azure.ResourceManager.WebPubSub;

// Generated from example definition: specification/webpubsub/resource-manager/Microsoft.SignalRService/stable/2021-10-01/examples/WebPubSubPrivateEndpointConnections_Delete.json
// this example is just showing the usage of "WebPubSubPrivateEndpointConnections_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this WebPubSubPrivateEndpointConnectionResource created on azure
// for more information of creating WebPubSubPrivateEndpointConnectionResource, please refer to the document of WebPubSubPrivateEndpointConnectionResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "myResourceGroup";
string resourceName = "myWebPubSubService";
string privateEndpointConnectionName = "mywebpubsubservice.1fa229cd-bf3f-47f0-8c49-afb36723997e";
ResourceIdentifier webPubSubPrivateEndpointConnectionResourceId = WebPubSubPrivateEndpointConnectionResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, resourceName, privateEndpointConnectionName);
WebPubSubPrivateEndpointConnectionResource webPubSubPrivateEndpointConnection = client.GetWebPubSubPrivateEndpointConnectionResource(webPubSubPrivateEndpointConnectionResourceId);

// invoke the operation
await webPubSubPrivateEndpointConnection.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
