using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.ServiceBus.Models;
using Azure.ResourceManager.ServiceBus;

// Generated from example definition: specification/servicebus/resource-manager/Microsoft.ServiceBus/ServiceBus/preview/2025-05-01-preview/examples/Topics/SBTopicAuthorizationRuleListKey.json
// this example is just showing the usage of "Topics_ListKeys" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ServiceBusTopicAuthorizationRuleResource created on azure
// for more information of creating ServiceBusTopicAuthorizationRuleResource, please refer to the document of ServiceBusTopicAuthorizationRuleResource
string subscriptionId = "e2f361f0-3b27-4503-a9cc-21cfba380093";
string resourceGroupName = "Default-ServiceBus-WestUS";
string namespaceName = "sdk-Namespace8408";
string topicName = "sdk-Topics2075";
string authorizationRuleName = "sdk-Authrules5067";
ResourceIdentifier serviceBusTopicAuthorizationRuleResourceId = ServiceBusTopicAuthorizationRuleResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, namespaceName, topicName, authorizationRuleName);
ServiceBusTopicAuthorizationRuleResource serviceBusTopicAuthorizationRule = client.GetServiceBusTopicAuthorizationRuleResource(serviceBusTopicAuthorizationRuleResourceId);

// invoke the operation
ServiceBusAccessKeys result = await serviceBusTopicAuthorizationRule.GetKeysAsync();

Console.WriteLine($"Succeeded: {result}");
