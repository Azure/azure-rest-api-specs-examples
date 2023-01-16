using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.ServiceBus;
using Azure.ResourceManager.ServiceBus.Models;

// Generated from example definition: specification/servicebus/resource-manager/Microsoft.ServiceBus/preview/2022-10-01-preview/examples/Rules/RuleCreate_SqlFilter.json
// this example is just showing the usage of "Rules_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ServiceBusSubscriptionResource created on azure
// for more information of creating ServiceBusSubscriptionResource, please refer to the document of ServiceBusSubscriptionResource
string subscriptionId = "subscriptionId";
string resourceGroupName = "resourceGroupName";
string namespaceName = "sdk-Namespace-1319";
string topicName = "sdk-Topics-2081";
string subscriptionName = "sdk-Subscriptions-8691";
ResourceIdentifier serviceBusSubscriptionResourceId = ServiceBusSubscriptionResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, namespaceName, topicName, subscriptionName);
ServiceBusSubscriptionResource serviceBusSubscription = client.GetServiceBusSubscriptionResource(serviceBusSubscriptionResourceId);

// get the collection of this ServiceBusRuleResource
ServiceBusRuleCollection collection = serviceBusSubscription.GetServiceBusRules();

// invoke the operation
string ruleName = "sdk-Rules-6571";
ServiceBusRuleData data = new ServiceBusRuleData()
{
    FilterType = ServiceBusFilterType.SqlFilter,
    SqlFilter = new ServiceBusSqlFilter()
    {
        SqlExpression = "myproperty=test",
    },
};
ArmOperation<ServiceBusRuleResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, ruleName, data);
ServiceBusRuleResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
ServiceBusRuleData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
