using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.ServiceBus.Models;
using Azure.ResourceManager.ServiceBus;

// Generated from example definition: specification/servicebus/resource-manager/Microsoft.ServiceBus/stable/2024-01-01/examples/Queues/SBQueueAuthorizationRuleCreate.json
// this example is just showing the usage of "QueueAuthorizationRules_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ServiceBusQueueAuthorizationRuleResource created on azure
// for more information of creating ServiceBusQueueAuthorizationRuleResource, please refer to the document of ServiceBusQueueAuthorizationRuleResource
string subscriptionId = "5f750a97-50d9-4e36-8081-c9ee4c0210d4";
string resourceGroupName = "ArunMonocle";
string namespaceName = "sdk-Namespace-7982";
string queueName = "sdk-Queues-2317";
string authorizationRuleName = "sdk-AuthRules-5800";
ResourceIdentifier serviceBusQueueAuthorizationRuleResourceId = ServiceBusQueueAuthorizationRuleResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, namespaceName, queueName, authorizationRuleName);
ServiceBusQueueAuthorizationRuleResource serviceBusQueueAuthorizationRule = client.GetServiceBusQueueAuthorizationRuleResource(serviceBusQueueAuthorizationRuleResourceId);

// invoke the operation
ServiceBusAuthorizationRuleData data = new ServiceBusAuthorizationRuleData
{
    Rights = { ServiceBusAccessRight.Listen, ServiceBusAccessRight.Send },
};
ArmOperation<ServiceBusQueueAuthorizationRuleResource> lro = await serviceBusQueueAuthorizationRule.UpdateAsync(WaitUntil.Completed, data);
ServiceBusQueueAuthorizationRuleResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
ServiceBusAuthorizationRuleData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
