using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.WebPubSub.Models;
using Azure.ResourceManager.WebPubSub;

// Generated from example definition: specification/webpubsub/resource-manager/Microsoft.SignalRService/stable/2021-10-01/examples/WebPubSubHubs_CreateOrUpdate.json
// this example is just showing the usage of "WebPubSubHubs_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

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
WebPubSubHubData data = new WebPubSubHubData(new WebPubSubHubProperties
{
    EventHandlers = {new WebPubSubEventHandler("http://host.com")
    {
    UserEventPattern = "*",
    SystemEvents = {"connect", "connected"},
    Auth = new UpstreamAuthSettings
    {
    AuthType = UpstreamAuthType.ManagedIdentity,
    ManagedIdentityResource = "abc",
    },
    }},
});
ArmOperation<WebPubSubHubResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, hubName, data);
WebPubSubHubResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
WebPubSubHubData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
