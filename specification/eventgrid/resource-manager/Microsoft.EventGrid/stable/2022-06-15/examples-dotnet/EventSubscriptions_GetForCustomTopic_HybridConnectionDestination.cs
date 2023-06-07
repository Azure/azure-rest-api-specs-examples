using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.EventGrid;
using Azure.ResourceManager.EventGrid.Models;

// Generated from example definition: specification/eventgrid/resource-manager/Microsoft.EventGrid/stable/2022-06-15/examples/EventSubscriptions_GetForCustomTopic_HybridConnectionDestination.json
// this example is just showing the usage of "EventSubscriptions_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this EventSubscriptionResource created on azure
// for more information of creating EventSubscriptionResource, please refer to the document of EventSubscriptionResource
string scope = "subscriptions/5b4b650e-28b9-4790-b3ab-ddbd88d727c4/resourceGroups/examplerg/providers/Microsoft.EventGrid/topics/exampletopic2";
string eventSubscriptionName = "examplesubscription1";
ResourceIdentifier eventSubscriptionResourceId = EventSubscriptionResource.CreateResourceIdentifier(scope, eventSubscriptionName);
EventSubscriptionResource eventSubscription = client.GetEventSubscriptionResource(eventSubscriptionResourceId);

// invoke the operation
EventSubscriptionResource result = await eventSubscription.GetAsync();

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
EventGridSubscriptionData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
