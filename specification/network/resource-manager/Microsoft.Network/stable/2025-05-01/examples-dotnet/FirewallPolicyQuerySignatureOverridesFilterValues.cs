using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Network.Models;
using Azure.ResourceManager.Network;

// Generated from example definition: specification/network/resource-manager/Microsoft.Network/stable/2025-05-01/examples/FirewallPolicyQuerySignatureOverridesFilterValues.json
// this example is just showing the usage of "FirewallPolicyIdpsSignaturesFilterValues_List" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this FirewallPolicyResource created on azure
// for more information of creating FirewallPolicyResource, please refer to the document of FirewallPolicyResource
string subscriptionId = "e747cc13-97d4-4a79-b463-42d7f4e558f2";
string resourceGroupName = "rg1";
string firewallPolicyName = "firewallPolicy";
ResourceIdentifier firewallPolicyResourceId = FirewallPolicyResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, firewallPolicyName);
FirewallPolicyResource firewallPolicy = client.GetFirewallPolicyResource(firewallPolicyResourceId);

// invoke the operation
SignatureOverridesFilterValuesQueryContent content = new SignatureOverridesFilterValuesQueryContent
{
    FilterName = "severity",
};
SignatureOverridesFilterValuesResult result = await firewallPolicy.GetFirewallPolicyIdpsSignaturesFilterValueAsync(content);

Console.WriteLine($"Succeeded: {result}");
