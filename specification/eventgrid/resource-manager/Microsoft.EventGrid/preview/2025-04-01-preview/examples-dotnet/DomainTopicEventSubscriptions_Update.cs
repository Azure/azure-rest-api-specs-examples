using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.EventGrid.Models;
using Azure.ResourceManager.EventGrid;

// Generated from example definition: specification/eventgrid/resource-manager/Microsoft.EventGrid/preview/2025-04-01-preview/examples/DomainTopicEventSubscriptions_Update.json
// this example is just showing the usage of "DomainTopicEventSubscriptions_Update" operation, for the dependent resources, they will have to be created separately.

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
string eventSubscriptionName = "exampleEventSubscriptionName1";
ResourceIdentifier domainTopicEventSubscriptionResourceId = DomainTopicEventSubscriptionResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, domainName, topicName, eventSubscriptionName);
DomainTopicEventSubscriptionResource domainTopicEventSubscription = client.GetDomainTopicEventSubscriptionResource(domainTopicEventSubscriptionResourceId);

// invoke the operation
EventGridSubscriptionPatch patch = new EventGridSubscriptionPatch
{
    Destination = new WebHookEventSubscriptionDestination
    {
        Endpoint = new Uri("https://requestb.in/15ksip71"),
    },
    Filter = new EventSubscriptionFilter
    {
        SubjectBeginsWith = "existingPrefix",
        SubjectEndsWith = "newSuffix",
        IsSubjectCaseSensitive = true,
    },
    Labels = { "label1", "label2" },
};
ArmOperation<DomainTopicEventSubscriptionResource> lro = await domainTopicEventSubscription.UpdateAsync(WaitUntil.Completed, patch);
DomainTopicEventSubscriptionResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
EventGridSubscriptionData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
