using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.ServiceBus.Models;
using Azure.ResourceManager.ServiceBus;

// Generated from example definition: specification/servicebus/resource-manager/Microsoft.ServiceBus/preview/2022-10-01-preview/examples/Queues/SBQueueAuthorizationRuleGet.json
// this example is just showing the usage of "QueueAuthorizationRules_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ServiceBusQueueResource created on azure
// for more information of creating ServiceBusQueueResource, please refer to the document of ServiceBusQueueResource
string subscriptionId = "5f750a97-50d9-4e36-8081-c9ee4c0210d4";
string resourceGroupName = "ArunMonocle";
string namespaceName = "sdk-Namespace-7982";
string queueName = "sdk-Queues-2317";
ResourceIdentifier serviceBusQueueResourceId = ServiceBusQueueResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, namespaceName, queueName);
ServiceBusQueueResource serviceBusQueue = client.GetServiceBusQueueResource(serviceBusQueueResourceId);

// get the collection of this ServiceBusQueueAuthorizationRuleResource
ServiceBusQueueAuthorizationRuleCollection collection = serviceBusQueue.GetServiceBusQueueAuthorizationRules();

// invoke the operation
string authorizationRuleName = "sdk-AuthRules-5800";
NullableResponse<ServiceBusQueueAuthorizationRuleResource> response = await collection.GetIfExistsAsync(authorizationRuleName);
ServiceBusQueueAuthorizationRuleResource result = response.HasValue ? response.Value : null;

if (result == null)
{
    Console.WriteLine($"Succeeded with null as result");
}
else
{
    // the variable result is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    ServiceBusAuthorizationRuleData resourceData = result.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}
