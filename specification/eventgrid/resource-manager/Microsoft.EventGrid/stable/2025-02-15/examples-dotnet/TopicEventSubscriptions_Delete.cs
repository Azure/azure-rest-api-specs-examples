using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.EventGrid.Models;
using Azure.ResourceManager.EventGrid;

// Generated from example definition: specification/eventgrid/resource-manager/Microsoft.EventGrid/stable/2025-02-15/examples/TopicEventSubscriptions_Delete.json
// this example is just showing the usage of "TopicEventSubscriptions_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this TopicEventSubscriptionResource created on azure
// for more information of creating TopicEventSubscriptionResource, please refer to the document of TopicEventSubscriptionResource
string subscriptionId = "5b4b650e-28b9-4790-b3ab-ddbd88d727c4";
string resourceGroupName = "examplerg";
string topicName = "exampleTopic";
string eventSubscriptionName = "examplesubscription";
ResourceIdentifier topicEventSubscriptionResourceId = TopicEventSubscriptionResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, topicName, eventSubscriptionName);
TopicEventSubscriptionResource topicEventSubscription = client.GetTopicEventSubscriptionResource(topicEventSubscriptionResourceId);

// invoke the operation
await topicEventSubscription.DeleteAsync(WaitUntil.Completed);

Console.WriteLine("Succeeded");
