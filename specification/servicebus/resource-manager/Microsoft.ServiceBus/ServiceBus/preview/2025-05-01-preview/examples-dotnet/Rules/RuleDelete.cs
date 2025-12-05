using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.ServiceBus.Models;
using Azure.ResourceManager.ServiceBus;

// Generated from example definition: specification/servicebus/resource-manager/Microsoft.ServiceBus/ServiceBus/preview/2025-05-01-preview/examples/Rules/RuleDelete.json
// this example is just showing the usage of "Rules_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ServiceBusRuleResource created on azure
// for more information of creating ServiceBusRuleResource, please refer to the document of ServiceBusRuleResource
string subscriptionId = "5f750a97-50d9-4e36-8081-c9ee4c0210d4";
string resourceGroupName = "ArunMonocle";
string namespaceName = "sdk-Namespace-1319";
string topicName = "sdk-Topics-2081";
string subscriptionName = "sdk-Subscriptions-8691";
string ruleName = "sdk-Rules-6571";
ResourceIdentifier serviceBusRuleResourceId = ServiceBusRuleResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, namespaceName, topicName, subscriptionName, ruleName);
ServiceBusRuleResource serviceBusRule = client.GetServiceBusRuleResource(serviceBusRuleResourceId);

// invoke the operation
await serviceBusRule.DeleteAsync(WaitUntil.Completed);

Console.WriteLine("Succeeded");
