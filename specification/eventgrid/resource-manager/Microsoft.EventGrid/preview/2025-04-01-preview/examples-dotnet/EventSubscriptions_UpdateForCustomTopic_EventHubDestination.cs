using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.EventGrid.Models;
using Azure.ResourceManager.EventGrid;

// Generated from example definition: specification/eventgrid/resource-manager/Microsoft.EventGrid/preview/2025-04-01-preview/examples/EventSubscriptions_UpdateForCustomTopic_EventHubDestination.json
// this example is just showing the usage of "EventSubscriptions_Update" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this EventSubscriptionResource created on azure
// for more information of creating EventSubscriptionResource, please refer to the document of EventSubscriptionResource
string scope = "subscriptions/8f6b6269-84f2-4d09-9e31-1127efcd1e40/resourceGroups/examplerg/providers/Microsoft.EventGrid/topics/exampletopic2";
string eventSubscriptionName = "examplesubscription1";
ResourceIdentifier eventSubscriptionResourceId = EventSubscriptionResource.CreateResourceIdentifier(scope, eventSubscriptionName);
EventSubscriptionResource eventSubscription = client.GetEventSubscriptionResource(eventSubscriptionResourceId);

// invoke the operation
EventGridSubscriptionPatch patch = new EventGridSubscriptionPatch
{
    Destination = new EventHubEventSubscriptionDestination
    {
        ResourceId = new ResourceIdentifier("/subscriptions/55f3dcd4-cac7-43b4-990b-a139d62a1eb2/resourceGroups/TestRG/providers/Microsoft.EventHub/namespaces/ContosoNamespace/eventhubs/EH1"),
    },
    Filter = new EventSubscriptionFilter
    {
        SubjectBeginsWith = "existingPrefix",
        SubjectEndsWith = "newSuffix",
        IsSubjectCaseSensitive = true,
    },
    Labels = { "label1", "label2" },
};
ArmOperation<EventSubscriptionResource> lro = await eventSubscription.UpdateAsync(WaitUntil.Completed, patch);
EventSubscriptionResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
EventGridSubscriptionData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
