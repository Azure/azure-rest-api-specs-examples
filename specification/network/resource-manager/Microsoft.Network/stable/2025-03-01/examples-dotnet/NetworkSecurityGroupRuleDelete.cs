using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Network.Models;
using Azure.ResourceManager.Network;

// Generated from example definition: specification/network/resource-manager/Microsoft.Network/stable/2025-03-01/examples/NetworkSecurityGroupRuleDelete.json
// this example is just showing the usage of "SecurityRules_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SecurityRuleResource created on azure
// for more information of creating SecurityRuleResource, please refer to the document of SecurityRuleResource
string subscriptionId = "subid";
string resourceGroupName = "rg1";
string networkSecurityGroupName = "testnsg";
string securityRuleName = "rule1";
ResourceIdentifier securityRuleResourceId = SecurityRuleResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, networkSecurityGroupName, securityRuleName);
SecurityRuleResource securityRule = client.GetSecurityRuleResource(securityRuleResourceId);

// invoke the operation
await securityRule.DeleteAsync(WaitUntil.Completed);

Console.WriteLine("Succeeded");
