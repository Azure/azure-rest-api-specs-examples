using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using System.Xml;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.EventGrid.Models;
using Azure.ResourceManager.EventGrid;

// Generated from example definition: specification/eventgrid/resource-manager/Microsoft.EventGrid/preview/2025-04-01-preview/examples/NamespaceTopicEventSubscriptions_CreateOrUpdate.json
// this example is just showing the usage of "NamespaceTopicEventSubscriptions_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this NamespaceTopicResource created on azure
// for more information of creating NamespaceTopicResource, please refer to the document of NamespaceTopicResource
string subscriptionId = "8f6b6269-84f2-4d09-9e31-1127efcd1e40";
string resourceGroupName = "examplerg";
string namespaceName = "examplenamespace2";
string topicName = "examplenamespacetopic2";
ResourceIdentifier namespaceTopicResourceId = NamespaceTopicResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, namespaceName, topicName);
NamespaceTopicResource namespaceTopic = client.GetNamespaceTopicResource(namespaceTopicResourceId);

// get the collection of this NamespaceTopicEventSubscriptionResource
NamespaceTopicEventSubscriptionCollection collection = namespaceTopic.GetNamespaceTopicEventSubscriptions();

// invoke the operation
string eventSubscriptionName = "examplenamespacetopicEventSub2";
NamespaceTopicEventSubscriptionData data = new NamespaceTopicEventSubscriptionData
{
    DeliveryConfiguration = new DeliveryConfiguration
    {
        DeliveryMode = DeliveryMode.Queue,
        Queue = new QueueInfo
        {
            ReceiveLockDurationInSeconds = 60,
            MaxDeliveryCount = 4,
            EventTimeToLive = XmlConvert.ToTimeSpan("P1D"),
        },
    },
    EventDeliverySchema = DeliverySchema.CloudEventSchemaV10,
};
ArmOperation<NamespaceTopicEventSubscriptionResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, eventSubscriptionName, data);
NamespaceTopicEventSubscriptionResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
NamespaceTopicEventSubscriptionData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
