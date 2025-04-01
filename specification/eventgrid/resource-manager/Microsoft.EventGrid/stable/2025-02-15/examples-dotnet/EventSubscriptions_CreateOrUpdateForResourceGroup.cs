using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.EventGrid.Models;
using Azure.ResourceManager.EventGrid;

// Generated from example definition: specification/eventgrid/resource-manager/Microsoft.EventGrid/stable/2025-02-15/examples/EventSubscriptions_CreateOrUpdateForResourceGroup.json
// this example is just showing the usage of "EventSubscriptions_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// get the collection of this EventSubscriptionResource
string scope = "subscriptions/5b4b650e-28b9-4790-b3ab-ddbd88d727c4/resourceGroups/examplerg";
EventSubscriptionCollection collection = client.GetEventSubscriptions(new ResourceIdentifier(scope));

// invoke the operation
string eventSubscriptionName = "examplesubscription2";
EventGridSubscriptionData data = new EventGridSubscriptionData
{
    Destination = new WebHookEventSubscriptionDestination
    {
        Endpoint = new Uri("https://requestb.in/15ksip71"),
    },
    Filter = new EventSubscriptionFilter
    {
        SubjectBeginsWith = "ExamplePrefix",
        SubjectEndsWith = "ExampleSuffix",
        IsSubjectCaseSensitive = false,
    },
};
ArmOperation<EventSubscriptionResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, eventSubscriptionName, data);
EventSubscriptionResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
EventGridSubscriptionData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
