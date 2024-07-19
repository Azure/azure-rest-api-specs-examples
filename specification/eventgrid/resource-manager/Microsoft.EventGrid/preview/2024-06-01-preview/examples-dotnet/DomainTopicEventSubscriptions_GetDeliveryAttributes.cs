using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.EventGrid.Models;
using Azure.ResourceManager.EventGrid;

// Generated from example definition: specification/eventgrid/resource-manager/Microsoft.EventGrid/preview/2024-06-01-preview/examples/DomainTopicEventSubscriptions_GetDeliveryAttributes.json
// this example is just showing the usage of "DomainTopicEventSubscriptions_GetDeliveryAttributes" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this DomainTopicEventSubscriptionResource created on azure
// for more information of creating DomainTopicEventSubscriptionResource, please refer to the document of DomainTopicEventSubscriptionResource
string subscriptionId = "8f6b6269-84f2-4d09-9e31-1127efcd1e40";
string resourceGroupName = "examplerg";
string domainName = "exampleDomain1";
string topicName = "exampleDomainTopic1";
string eventSubscriptionName = "examplesubscription1";
ResourceIdentifier domainTopicEventSubscriptionResourceId = DomainTopicEventSubscriptionResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, domainName, topicName, eventSubscriptionName);
DomainTopicEventSubscriptionResource domainTopicEventSubscription = client.GetDomainTopicEventSubscriptionResource(domainTopicEventSubscriptionResourceId);

// invoke the operation and iterate over the result
await foreach (DeliveryAttributeMapping item in domainTopicEventSubscription.GetDeliveryAttributesAsync())
{
    Console.WriteLine($"Succeeded: {item}");
}

Console.WriteLine($"Succeeded");
