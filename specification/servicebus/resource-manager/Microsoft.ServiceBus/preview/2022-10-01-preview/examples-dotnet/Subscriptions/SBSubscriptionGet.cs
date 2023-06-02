using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.ServiceBus;

// Generated from example definition: specification/servicebus/resource-manager/Microsoft.ServiceBus/preview/2022-10-01-preview/examples/Subscriptions/SBSubscriptionGet.json
// this example is just showing the usage of "Subscriptions_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ServiceBusTopicResource created on azure
// for more information of creating ServiceBusTopicResource, please refer to the document of ServiceBusTopicResource
string subscriptionId = "Subscriptionid";
string resourceGroupName = "ResourceGroup";
string namespaceName = "sdk-Namespace-1349";
string topicName = "sdk-Topics-8740";
ResourceIdentifier serviceBusTopicResourceId = ServiceBusTopicResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, namespaceName, topicName);
ServiceBusTopicResource serviceBusTopic = client.GetServiceBusTopicResource(serviceBusTopicResourceId);

// get the collection of this ServiceBusSubscriptionResource
ServiceBusSubscriptionCollection collection = serviceBusTopic.GetServiceBusSubscriptions();

// invoke the operation
string subscriptionName = "sdk-Subscriptions-2178";
bool result = await collection.ExistsAsync(subscriptionName);

Console.WriteLine($"Succeeded: {result}");
