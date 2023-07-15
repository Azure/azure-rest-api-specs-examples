using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.PaloAltoNetworks.Ngfw;
using Azure.ResourceManager.PaloAltoNetworks.Ngfw.Models;

// Generated from example definition: specification/paloaltonetworks/resource-manager/PaloAltoNetworks.Cloudngfw/stable/2022-08-29/examples/PostRules_refreshCounters_MaximumSet_Gen.json
// this example is just showing the usage of "PostRules_refreshCounters" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this PostRulestackRuleResource created on azure
// for more information of creating PostRulestackRuleResource, please refer to the document of PostRulestackRuleResource
string globalRulestackName = "lrs1";
string priority = "1";
ResourceIdentifier postRulestackRuleResourceId = PostRulestackRuleResource.CreateResourceIdentifier(globalRulestackName, priority);
PostRulestackRuleResource postRulestackRule = client.GetPostRulestackRuleResource(postRulestackRuleResourceId);

// invoke the operation
string firewallName = "firewall1";
await postRulestackRule.RefreshCountersAsync(firewallName: firewallName);

Console.WriteLine($"Succeeded");
