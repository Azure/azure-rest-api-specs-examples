using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.EventGrid.Models;
using Azure.ResourceManager.EventGrid;

// Generated from example definition: specification/eventgrid/resource-manager/Microsoft.EventGrid/preview/2025-04-01-preview/examples/SystemTopicEventSubscriptions_Get.json
// this example is just showing the usage of "SystemTopicEventSubscriptions_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SystemTopicResource created on azure
// for more information of creating SystemTopicResource, please refer to the document of SystemTopicResource
string subscriptionId = "8f6b6269-84f2-4d09-9e31-1127efcd1e40";
string resourceGroupName = "examplerg";
string systemTopicName = "exampleSystemTopic1";
ResourceIdentifier systemTopicResourceId = SystemTopicResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, systemTopicName);
SystemTopicResource systemTopic = client.GetSystemTopicResource(systemTopicResourceId);

// get the collection of this SystemTopicEventSubscriptionResource
SystemTopicEventSubscriptionCollection collection = systemTopic.GetSystemTopicEventSubscriptions();

// invoke the operation
string eventSubscriptionName = "examplesubscription1";
NullableResponse<SystemTopicEventSubscriptionResource> response = await collection.GetIfExistsAsync(eventSubscriptionName);
SystemTopicEventSubscriptionResource result = response.HasValue ? response.Value : null;

if (result == null)
{
    Console.WriteLine("Succeeded with null as result");
}
else
{
    // the variable result is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    EventGridSubscriptionData resourceData = result.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}
