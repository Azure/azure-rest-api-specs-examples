using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.ServiceBus.Models;
using Azure.ResourceManager.ServiceBus;

// Generated from example definition: specification/servicebus/resource-manager/Microsoft.ServiceBus/stable/2024-01-01/examples/Rules/RuleListBySubscription.json
// this example is just showing the usage of "Rules_ListBySubscriptions" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ServiceBusSubscriptionResource created on azure
// for more information of creating ServiceBusSubscriptionResource, please refer to the document of ServiceBusSubscriptionResource
string subscriptionId = "5f750a97-50d9-4e36-8081-c9ee4c0210d4";
string resourceGroupName = "ArunMonocle";
string namespaceName = "sdk-Namespace-1319";
string topicName = "sdk-Topics-2081";
string subscriptionName = "sdk-Subscriptions-8691";
ResourceIdentifier serviceBusSubscriptionResourceId = ServiceBusSubscriptionResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, namespaceName, topicName, subscriptionName);
ServiceBusSubscriptionResource serviceBusSubscription = client.GetServiceBusSubscriptionResource(serviceBusSubscriptionResourceId);

// get the collection of this ServiceBusRuleResource
ServiceBusRuleCollection collection = serviceBusSubscription.GetServiceBusRules();

// invoke the operation and iterate over the result
await foreach (ServiceBusRuleResource item in collection.GetAllAsync())
{
    // the variable item is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    ServiceBusRuleData resourceData = item.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}

Console.WriteLine("Succeeded");
