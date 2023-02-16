using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.WebPubSub;
using Azure.ResourceManager.WebPubSub.Models;

// Generated from example definition: specification/webpubsub/resource-manager/Microsoft.SignalRService/stable/2021-10-01/examples/WebPubSubHubs_Get.json
// this example is just showing the usage of "WebPubSubHubs_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this WebPubSubResource created on azure
// for more information of creating WebPubSubResource, please refer to the document of WebPubSubResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "myResourceGroup";
string resourceName = "myWebPubSubService";
ResourceIdentifier webPubSubResourceId = WebPubSubResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, resourceName);
WebPubSubResource webPubSub = client.GetWebPubSubResource(webPubSubResourceId);

// get the collection of this WebPubSubHubResource
WebPubSubHubCollection collection = webPubSub.GetWebPubSubHubs();

// invoke the operation
string hubName = "exampleHub";
bool result = await collection.ExistsAsync(hubName);

Console.WriteLine($"Succeeded: {result}");
