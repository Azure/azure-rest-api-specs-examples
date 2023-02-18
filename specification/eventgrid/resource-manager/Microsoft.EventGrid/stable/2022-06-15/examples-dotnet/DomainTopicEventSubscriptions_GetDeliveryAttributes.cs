using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.EventGrid;
using Azure.ResourceManager.EventGrid.Models;

// Generated from example definition: specification/eventgrid/resource-manager/Microsoft.EventGrid/stable/2022-06-15/examples/DomainTopicEventSubscriptions_GetDeliveryAttributes.json
// this example is just showing the usage of "DomainTopicEventSubscriptions_GetDeliveryAttributes" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this DomainTopicEventSubscriptionResource created on azure
// for more information of creating DomainTopicEventSubscriptionResource, please refer to the document of DomainTopicEventSubscriptionResource
string subscriptionId = "5b4b650e-28b9-4790-b3ab-ddbd88d727c4";
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
