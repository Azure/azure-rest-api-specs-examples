using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Cdn.Models;
using Azure.ResourceManager.Cdn;

// Generated from example definition: specification/cdn/resource-manager/Microsoft.Cdn/stable/2024-02-01/examples/WafPolicyDelete.json
// this example is just showing the usage of "Policies_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this CdnWebApplicationFirewallPolicyResource created on azure
// for more information of creating CdnWebApplicationFirewallPolicyResource, please refer to the document of CdnWebApplicationFirewallPolicyResource
string subscriptionId = "subid";
string resourceGroupName = "rg1";
string policyName = "Policy1";
ResourceIdentifier cdnWebApplicationFirewallPolicyResourceId = CdnWebApplicationFirewallPolicyResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, policyName);
CdnWebApplicationFirewallPolicyResource cdnWebApplicationFirewallPolicy = client.GetCdnWebApplicationFirewallPolicyResource(cdnWebApplicationFirewallPolicyResourceId);

// invoke the operation
await cdnWebApplicationFirewallPolicy.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
