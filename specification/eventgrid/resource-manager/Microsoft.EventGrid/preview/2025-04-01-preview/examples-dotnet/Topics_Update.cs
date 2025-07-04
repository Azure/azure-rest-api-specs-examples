using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.EventGrid.Models;
using Azure.ResourceManager.EventGrid;

// Generated from example definition: specification/eventgrid/resource-manager/Microsoft.EventGrid/preview/2025-04-01-preview/examples/Topics_Update.json
// this example is just showing the usage of "Topics_Update" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this EventGridTopicResource created on azure
// for more information of creating EventGridTopicResource, please refer to the document of EventGridTopicResource
string subscriptionId = "8f6b6269-84f2-4d09-9e31-1127efcd1e40";
string resourceGroupName = "examplerg";
string topicName = "exampletopic1";
ResourceIdentifier eventGridTopicResourceId = EventGridTopicResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, topicName);
EventGridTopicResource eventGridTopic = client.GetEventGridTopicResource(eventGridTopicResourceId);

// invoke the operation
EventGridTopicPatch patch = new EventGridTopicPatch
{
    Tags =
    {
    ["tag1"] = "value1",
    ["tag2"] = "value2"
    },
    PublicNetworkAccess = EventGridPublicNetworkAccess.Enabled,
    InboundIPRules = {new EventGridInboundIPRule
    {
    IPMask = "12.18.30.15",
    Action = EventGridIPActionType.Allow,
    }, new EventGridInboundIPRule
    {
    IPMask = "12.18.176.1",
    Action = EventGridIPActionType.Allow,
    }},
};
await eventGridTopic.UpdateAsync(WaitUntil.Completed, patch);

Console.WriteLine("Succeeded");
