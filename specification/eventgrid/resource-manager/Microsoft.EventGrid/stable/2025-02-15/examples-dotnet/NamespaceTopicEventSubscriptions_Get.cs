using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using System.Xml;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.EventGrid.Models;
using Azure.ResourceManager.EventGrid;

// Generated from example definition: specification/eventgrid/resource-manager/Microsoft.EventGrid/stable/2025-02-15/examples/NamespaceTopicEventSubscriptions_Get.json
// this example is just showing the usage of "NamespaceTopicEventSubscriptions_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this NamespaceTopicEventSubscriptionResource created on azure
// for more information of creating NamespaceTopicEventSubscriptionResource, please refer to the document of NamespaceTopicEventSubscriptionResource
string subscriptionId = "8f6b6269-84f2-4d09-9e31-1127efcd1e40";
string resourceGroupName = "examplerg";
string namespaceName = "examplenamespace2";
string topicName = "examplenamespacetopic2";
string eventSubscriptionName = "examplenamespacetopicEventSub1";
ResourceIdentifier namespaceTopicEventSubscriptionResourceId = NamespaceTopicEventSubscriptionResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, namespaceName, topicName, eventSubscriptionName);
NamespaceTopicEventSubscriptionResource namespaceTopicEventSubscription = client.GetNamespaceTopicEventSubscriptionResource(namespaceTopicEventSubscriptionResourceId);

// invoke the operation
NamespaceTopicEventSubscriptionResource result = await namespaceTopicEventSubscription.GetAsync();

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
NamespaceTopicEventSubscriptionData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
