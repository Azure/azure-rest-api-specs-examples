using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.EventHubs;

// Generated from example definition: specification/eventhub/resource-manager/Microsoft.EventHub/stable/2024-01-01/examples/ConsumerGroup/EHConsumerGroupCreate.json
// this example is just showing the usage of "ConsumerGroups_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this EventHubsConsumerGroupResource created on azure
// for more information of creating EventHubsConsumerGroupResource, please refer to the document of EventHubsConsumerGroupResource
string subscriptionId = "5f750a97-50d9-4e36-8081-c9ee4c0210d4";
string resourceGroupName = "ArunMonocle";
string namespaceName = "sdk-Namespace-2661";
string eventHubName = "sdk-EventHub-6681";
string consumerGroupName = "sdk-ConsumerGroup-5563";
ResourceIdentifier eventHubsConsumerGroupResourceId = EventHubsConsumerGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, namespaceName, eventHubName, consumerGroupName);
EventHubsConsumerGroupResource eventHubsConsumerGroup = client.GetEventHubsConsumerGroupResource(eventHubsConsumerGroupResourceId);

// invoke the operation
EventHubsConsumerGroupData data = new EventHubsConsumerGroupData
{
    UserMetadata = "New consumergroup",
};
ArmOperation<EventHubsConsumerGroupResource> lro = await eventHubsConsumerGroup.UpdateAsync(WaitUntil.Completed, data);
EventHubsConsumerGroupResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
EventHubsConsumerGroupData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
