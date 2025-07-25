using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.EventGrid.Models;
using Azure.ResourceManager.EventGrid;

// Generated from example definition: specification/eventgrid/resource-manager/Microsoft.EventGrid/preview/2025-04-01-preview/examples/PartnerTopicEventSubscriptions_GetFullUrl.json
// this example is just showing the usage of "PartnerTopicEventSubscriptions_GetFullUri" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this PartnerTopicEventSubscriptionResource created on azure
// for more information of creating PartnerTopicEventSubscriptionResource, please refer to the document of PartnerTopicEventSubscriptionResource
string subscriptionId = "8f6b6269-84f2-4d09-9e31-1127efcd1e40";
string resourceGroupName = "examplerg";
string partnerTopicName = "examplePartnerTopic1";
string eventSubscriptionName = "examplesubscription1";
ResourceIdentifier partnerTopicEventSubscriptionResourceId = PartnerTopicEventSubscriptionResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, partnerTopicName, eventSubscriptionName);
PartnerTopicEventSubscriptionResource partnerTopicEventSubscription = client.GetPartnerTopicEventSubscriptionResource(partnerTopicEventSubscriptionResourceId);

// invoke the operation
EventSubscriptionFullUri result = await partnerTopicEventSubscription.GetFullUriAsync();

Console.WriteLine($"Succeeded: {result}");
