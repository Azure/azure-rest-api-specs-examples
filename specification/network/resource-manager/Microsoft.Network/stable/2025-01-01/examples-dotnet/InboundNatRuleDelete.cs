using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Network.Models;
using Azure.ResourceManager.Network;

// Generated from example definition: specification/network/resource-manager/Microsoft.Network/stable/2025-01-01/examples/InboundNatRuleDelete.json
// this example is just showing the usage of "InboundNatRules_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this InboundNatRuleResource created on azure
// for more information of creating InboundNatRuleResource, please refer to the document of InboundNatRuleResource
string subscriptionId = "subid";
string resourceGroupName = "testrg";
string loadBalancerName = "lb1";
string inboundNatRuleName = "natRule1.1";
ResourceIdentifier inboundNatRuleResourceId = InboundNatRuleResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, loadBalancerName, inboundNatRuleName);
InboundNatRuleResource inboundNatRule = client.GetInboundNatRuleResource(inboundNatRuleResourceId);

// invoke the operation
await inboundNatRule.DeleteAsync(WaitUntil.Completed);

Console.WriteLine("Succeeded");
