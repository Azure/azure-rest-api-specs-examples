using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.ServiceBus.Models;
using Azure.ResourceManager.ServiceBus;

// Generated from example definition: specification/servicebus/resource-manager/Microsoft.ServiceBus/stable/2024-01-01/examples/disasterRecoveryConfigs/SBAliasAuthorizationRuleListKey.json
// this example is just showing the usage of "DisasterRecoveryAuthorizationRules_ListKeys" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ServiceBusDisasterRecoveryAuthorizationRuleResource created on azure
// for more information of creating ServiceBusDisasterRecoveryAuthorizationRuleResource, please refer to the document of ServiceBusDisasterRecoveryAuthorizationRuleResource
string subscriptionId = "exampleSubscriptionId";
string resourceGroupName = "exampleResourceGroup";
string namespaceName = "sdk-Namespace-2702";
string @alias = "sdk-DisasterRecovery-4047";
string authorizationRuleName = "sdk-Authrules-1746";
ResourceIdentifier serviceBusDisasterRecoveryAuthorizationRuleResourceId = ServiceBusDisasterRecoveryAuthorizationRuleResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, namespaceName, @alias, authorizationRuleName);
ServiceBusDisasterRecoveryAuthorizationRuleResource serviceBusDisasterRecoveryAuthorizationRule = client.GetServiceBusDisasterRecoveryAuthorizationRuleResource(serviceBusDisasterRecoveryAuthorizationRuleResourceId);

// invoke the operation
ServiceBusAccessKeys result = await serviceBusDisasterRecoveryAuthorizationRule.GetKeysAsync();

Console.WriteLine($"Succeeded: {result}");
