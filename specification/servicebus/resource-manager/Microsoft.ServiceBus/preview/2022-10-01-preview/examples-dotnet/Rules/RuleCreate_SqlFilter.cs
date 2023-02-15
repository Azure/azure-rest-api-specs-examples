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

// this example assumes you already have this ServiceBusRuleResource created on azure
// for more information of creating ServiceBusRuleResource, please refer to the document of ServiceBusRuleResource
string subscriptionId = "subscriptionId";
string resourceGroupName = "resourceGroupName";
string namespaceName = "sdk-Namespace-1319";
string topicName = "sdk-Topics-2081";
string subscriptionName = "sdk-Subscriptions-8691";
string ruleName = "sdk-Rules-6571";
ResourceIdentifier serviceBusRuleResourceId = ServiceBusRuleResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, namespaceName, topicName, subscriptionName, ruleName);
ServiceBusRuleResource serviceBusRule = client.GetServiceBusRuleResource(serviceBusRuleResourceId);

// invoke the operation
ServiceBusRuleData data = new ServiceBusRuleData()
{
    FilterType = ServiceBusFilterType.SqlFilter,
    SqlFilter = new ServiceBusSqlFilter()
    {
        SqlExpression = "myproperty=test",
    },
};
ArmOperation<ServiceBusRuleResource> lro = await serviceBusRule.UpdateAsync(WaitUntil.Completed, data);
ServiceBusRuleResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
ServiceBusRuleData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
