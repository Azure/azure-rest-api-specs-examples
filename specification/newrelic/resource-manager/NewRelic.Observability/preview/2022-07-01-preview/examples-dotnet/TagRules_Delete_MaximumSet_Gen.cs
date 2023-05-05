using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.NewRelicObservability;
using Azure.ResourceManager.NewRelicObservability.Models;

// Generated from example definition: specification/newrelic/resource-manager/NewRelic.Observability/preview/2022-07-01-preview/examples/TagRules_Delete_MaximumSet_Gen.json
// this example is just showing the usage of "TagRules_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this TagRuleResource created on azure
// for more information of creating TagRuleResource, please refer to the document of TagRuleResource
string subscriptionId = "ddqonpqwjr";
string resourceGroupName = "rgopenapi";
string monitorName = "ipxmlcbonyxtolzejcjshkmlron";
string ruleSetName = "bxcantgzggsepbhqmedjqyrqeezmfb";
ResourceIdentifier tagRuleResourceId = TagRuleResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, monitorName, ruleSetName);
TagRuleResource tagRule = client.GetTagRuleResource(tagRuleResourceId);

// invoke the operation
await tagRule.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
