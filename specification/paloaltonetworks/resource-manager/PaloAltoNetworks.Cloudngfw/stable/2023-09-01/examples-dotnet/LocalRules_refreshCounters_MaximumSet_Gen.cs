using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.PaloAltoNetworks.Ngfw.Models;
using Azure.ResourceManager.PaloAltoNetworks.Ngfw;

// Generated from example definition: specification/paloaltonetworks/resource-manager/PaloAltoNetworks.Cloudngfw/stable/2023-09-01/examples/LocalRules_refreshCounters_MaximumSet_Gen.json
// this example is just showing the usage of "LocalRules_RefreshCounters" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this LocalRulestackRuleResource created on azure
// for more information of creating LocalRulestackRuleResource, please refer to the document of LocalRulestackRuleResource
string subscriptionId = "2bf4a339-294d-4c25-b0b2-ef649e9f5c27";
string resourceGroupName = "firewall-rg";
string localRulestackName = "lrs1";
string priority = "1";
ResourceIdentifier localRulestackRuleResourceId = LocalRulestackRuleResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, localRulestackName, priority);
LocalRulestackRuleResource localRulestackRule = client.GetLocalRulestackRuleResource(localRulestackRuleResourceId);

// invoke the operation
string firewallName = "firewall1";
await localRulestackRule.RefreshCountersAsync(firewallName: firewallName);

Console.WriteLine("Succeeded");
