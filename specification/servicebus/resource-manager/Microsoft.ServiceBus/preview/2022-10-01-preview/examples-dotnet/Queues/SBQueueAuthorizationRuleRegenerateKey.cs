using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.ServiceBus;
using Azure.ResourceManager.ServiceBus.Models;

// Generated from example definition: specification/servicebus/resource-manager/Microsoft.ServiceBus/preview/2022-10-01-preview/examples/Queues/SBQueueAuthorizationRuleRegenerateKey.json
// this example is just showing the usage of "QueueAuthorizationRules_RegenerateKeys" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ServiceBusQueueAuthorizationRuleResource created on azure
// for more information of creating ServiceBusQueueAuthorizationRuleResource, please refer to the document of ServiceBusQueueAuthorizationRuleResource
string subscriptionId = "5f750a97-50d9-4e36-8081-c9ee4c0210d4";
string resourceGroupName = "ArunMonocle";
string namespaceName = "sdk-namespace-7982";
string queueName = "sdk-Queues-2317";
string authorizationRuleName = "sdk-AuthRules-5800";
ResourceIdentifier serviceBusQueueAuthorizationRuleResourceId = ServiceBusQueueAuthorizationRuleResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, namespaceName, queueName, authorizationRuleName);
ServiceBusQueueAuthorizationRuleResource serviceBusQueueAuthorizationRule = client.GetServiceBusQueueAuthorizationRuleResource(serviceBusQueueAuthorizationRuleResourceId);

// invoke the operation
ServiceBusRegenerateAccessKeyContent content = new ServiceBusRegenerateAccessKeyContent(ServiceBusAccessKeyType.PrimaryKey);
ServiceBusAccessKeys result = await serviceBusQueueAuthorizationRule.RegenerateKeysAsync(content);

Console.WriteLine($"Succeeded: {result}");
