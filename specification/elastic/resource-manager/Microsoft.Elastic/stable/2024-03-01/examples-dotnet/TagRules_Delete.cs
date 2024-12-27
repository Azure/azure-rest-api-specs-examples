using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Elastic;

// Generated from example definition: specification/elastic/resource-manager/Microsoft.Elastic/stable/2024-03-01/examples/TagRules_Delete.json
// this example is just showing the usage of "TagRules_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ElasticTagRuleResource created on azure
// for more information of creating ElasticTagRuleResource, please refer to the document of ElasticTagRuleResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "myResourceGroup";
string monitorName = "myMonitor";
string ruleSetName = "default";
ResourceIdentifier elasticTagRuleResourceId = ElasticTagRuleResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, monitorName, ruleSetName);
ElasticTagRuleResource elasticTagRule = client.GetElasticTagRuleResource(elasticTagRuleResourceId);

// invoke the operation
await elasticTagRule.DeleteAsync(WaitUntil.Completed);

Console.WriteLine("Succeeded");
