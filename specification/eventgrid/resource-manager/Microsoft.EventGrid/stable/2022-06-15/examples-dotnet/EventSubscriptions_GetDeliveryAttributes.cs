using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.EventGrid;
using Azure.ResourceManager.EventGrid.Models;

// Generated from example definition: specification/eventgrid/resource-manager/Microsoft.EventGrid/stable/2022-06-15/examples/EventSubscriptions_GetDeliveryAttributes.json
// this example is just showing the usage of "EventSubscriptions_GetDeliveryAttributes" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this EventSubscriptionResource created on azure
// for more information of creating EventSubscriptionResource, please refer to the document of EventSubscriptionResource
string scope = "aaaaaaaaaaaaaaaaaaaaaaaaa";
string eventSubscriptionName = "aaaaaaaaaaaaaaaaaa";
ResourceIdentifier eventSubscriptionResourceId = EventSubscriptionResource.CreateResourceIdentifier(scope, eventSubscriptionName);
EventSubscriptionResource eventSubscription = client.GetEventSubscriptionResource(eventSubscriptionResourceId);

// invoke the operation and iterate over the result
await foreach (DeliveryAttributeMapping item in eventSubscription.GetDeliveryAttributesAsync())
{
    Console.WriteLine($"Succeeded: {item}");
}

Console.WriteLine($"Succeeded");
