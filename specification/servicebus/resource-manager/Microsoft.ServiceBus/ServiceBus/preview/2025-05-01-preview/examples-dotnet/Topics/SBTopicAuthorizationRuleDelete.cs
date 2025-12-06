using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.ServiceBus.Models;
using Azure.ResourceManager.ServiceBus;

// Generated from example definition: specification/servicebus/resource-manager/Microsoft.ServiceBus/ServiceBus/preview/2025-05-01-preview/examples/Topics/SBTopicAuthorizationRuleDelete.json
// this example is just showing the usage of "TopicAuthorizationRules_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ServiceBusTopicAuthorizationRuleResource created on azure
// for more information of creating ServiceBusTopicAuthorizationRuleResource, please refer to the document of ServiceBusTopicAuthorizationRuleResource
string subscriptionId = "5f750a97-50d9-4e36-8081-c9ee4c0210d4";
string resourceGroupName = "ArunMonocle";
string namespaceName = "sdk-Namespace-6261";
string topicName = "sdk-Topics-1984";
string authorizationRuleName = "sdk-AuthRules-4310";
ResourceIdentifier serviceBusTopicAuthorizationRuleResourceId = ServiceBusTopicAuthorizationRuleResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, namespaceName, topicName, authorizationRuleName);
ServiceBusTopicAuthorizationRuleResource serviceBusTopicAuthorizationRule = client.GetServiceBusTopicAuthorizationRuleResource(serviceBusTopicAuthorizationRuleResourceId);

// invoke the operation
await serviceBusTopicAuthorizationRule.DeleteAsync(WaitUntil.Completed);

Console.WriteLine("Succeeded");
