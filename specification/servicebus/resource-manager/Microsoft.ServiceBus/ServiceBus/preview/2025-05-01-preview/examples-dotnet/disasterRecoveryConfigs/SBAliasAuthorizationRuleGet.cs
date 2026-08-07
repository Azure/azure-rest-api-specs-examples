using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.ServiceBus;

// Generated from example definition: specification/servicebus/resource-manager/Microsoft.ServiceBus/ServiceBus/preview/2025-05-01-preview/examples/disasterRecoveryConfigs/SBAliasAuthorizationRuleGet.json
// this example is just showing the usage of "DisasterRecoveryAuthorizationRules_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ServiceBusDisasterRecoveryResource created on azure
// for more information of creating ServiceBusDisasterRecoveryResource, please refer to the document of ServiceBusDisasterRecoveryResource
string subscriptionId = "5f750a97-50d9-4e36-8081-c9ee4c0210d4";
string resourceGroupName = "exampleResourceGroup";
string namespaceName = "sdk-Namespace-9080";
string @alias = "sdk-DisasterRecovery-4879";
ResourceIdentifier serviceBusDisasterRecoveryResourceId = ServiceBusDisasterRecoveryResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, namespaceName, @alias);
ServiceBusDisasterRecoveryResource serviceBusDisasterRecovery = client.GetServiceBusDisasterRecoveryResource(serviceBusDisasterRecoveryResourceId);

// get the collection of this ServiceBusDisasterRecoveryAuthorizationRuleResource
ServiceBusDisasterRecoveryAuthorizationRuleCollection collection = serviceBusDisasterRecovery.GetServiceBusDisasterRecoveryAuthorizationRules();

// invoke the operation
string authorizationRuleName = "sdk-Authrules-4879";
NullableResponse<ServiceBusDisasterRecoveryAuthorizationRuleResource> response = await collection.GetIfExistsAsync(authorizationRuleName);
ServiceBusDisasterRecoveryAuthorizationRuleResource result = response.HasValue ? response.Value : null;

if (result == null)
{
    Console.WriteLine("Succeeded with null as result");
}
else
{
    // the variable result is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    ServiceBusAuthorizationRuleData resourceData = result.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}
