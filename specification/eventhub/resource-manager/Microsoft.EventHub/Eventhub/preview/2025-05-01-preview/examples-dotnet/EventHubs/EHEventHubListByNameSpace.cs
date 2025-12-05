using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.EventHubs.Models;
using Azure.ResourceManager.EventHubs;

// Generated from example definition: specification/eventhub/resource-manager/Microsoft.EventHub/Eventhub/preview/2025-05-01-preview/examples/EventHubs/EHEventHubListByNameSpace.json
// this example is just showing the usage of "EventHubs_ListByNamespace" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this EventHubsNamespaceResource created on azure
// for more information of creating EventHubsNamespaceResource, please refer to the document of EventHubsNamespaceResource
string subscriptionId = "e2f361f0-3b27-4503-a9cc-21cfba380093";
string resourceGroupName = "Default-NotificationHubs-AustraliaEast";
string namespaceName = "sdk-Namespace-5357";
ResourceIdentifier eventHubsNamespaceResourceId = EventHubsNamespaceResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, namespaceName);
EventHubsNamespaceResource eventHubsNamespace = client.GetEventHubsNamespaceResource(eventHubsNamespaceResourceId);

// get the collection of this EventHubResource
EventHubCollection collection = eventHubsNamespace.GetEventHubs();

// invoke the operation and iterate over the result
await foreach (EventHubResource item in collection.GetAllAsync())
{
    // the variable item is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    EventHubData resourceData = item.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}

Console.WriteLine("Succeeded");
