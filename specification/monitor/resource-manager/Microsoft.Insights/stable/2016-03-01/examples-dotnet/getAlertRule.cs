using System;
using System.Threading.Tasks;
using System.Xml;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Monitor;
using Azure.ResourceManager.Monitor.Models;
using Azure.ResourceManager.Resources;

// Generated from example definition: specification/monitor/resource-manager/Microsoft.Insights/stable/2016-03-01/examples/getAlertRule.json
// this example is just showing the usage of "AlertRules_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this AlertRuleResource created on azure
// for more information of creating AlertRuleResource, please refer to the document of AlertRuleResource
string subscriptionId = "b67f7fec-69fc-4974-9099-a26bd6ffeda3";
string resourceGroupName = "Rac46PostSwapRG";
string ruleName = "chiricutin";
ResourceIdentifier alertRuleResourceId = AlertRuleResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, ruleName);
AlertRuleResource alertRule = client.GetAlertRuleResource(alertRuleResourceId);

// invoke the operation
AlertRuleResource result = await alertRule.GetAsync();

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
AlertRuleData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
