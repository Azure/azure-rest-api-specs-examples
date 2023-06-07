using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.EventGrid;
using Azure.ResourceManager.EventGrid.Models;

// Generated from example definition: specification/eventgrid/resource-manager/Microsoft.EventGrid/stable/2022-06-15/examples/PartnerTopicEventSubscriptions_Get.json
// this example is just showing the usage of "PartnerTopicEventSubscriptions_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this PartnerTopicEventSubscriptionResource created on azure
// for more information of creating PartnerTopicEventSubscriptionResource, please refer to the document of PartnerTopicEventSubscriptionResource
string subscriptionId = "5b4b650e-28b9-4790-b3ab-ddbd88d727c4";
string resourceGroupName = "examplerg";
string partnerTopicName = "examplePartnerTopic1";
string eventSubscriptionName = "examplesubscription1";
ResourceIdentifier partnerTopicEventSubscriptionResourceId = PartnerTopicEventSubscriptionResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, partnerTopicName, eventSubscriptionName);
PartnerTopicEventSubscriptionResource partnerTopicEventSubscription = client.GetPartnerTopicEventSubscriptionResource(partnerTopicEventSubscriptionResourceId);

// invoke the operation
PartnerTopicEventSubscriptionResource result = await partnerTopicEventSubscription.GetAsync();

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
EventGridSubscriptionData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
