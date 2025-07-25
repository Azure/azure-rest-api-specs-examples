using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.EventGrid;

// Generated from example definition: specification/eventgrid/resource-manager/Microsoft.EventGrid/preview/2025-04-01-preview/examples/PrivateLinkResources_ListByResource.json
// this example is just showing the usage of "PrivateLinkResources_ListByResource" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this EventGridDomainResource created on azure
// for more information of creating EventGridDomainResource, please refer to the document of EventGridDomainResource
string subscriptionId = "8f6b6269-84f2-4d09-9e31-1127efcd1e40";
string resourceGroupName = "examplerg";
string parentName = "exampletopic1";
ResourceIdentifier eventGridDomainResourceId = EventGridDomainResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, parentName);
EventGridDomainResource eventGridDomain = client.GetEventGridDomainResource(eventGridDomainResourceId);

// get the collection of this EventGridDomainPrivateLinkResource
EventGridDomainPrivateLinkResourceCollection collection = eventGridDomain.GetEventGridDomainPrivateLinkResources();

// invoke the operation and iterate over the result
await foreach (EventGridDomainPrivateLinkResource item in collection.GetAllAsync())
{
    // the variable item is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    EventGridPrivateLinkResourceData resourceData = item.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}

Console.WriteLine("Succeeded");
