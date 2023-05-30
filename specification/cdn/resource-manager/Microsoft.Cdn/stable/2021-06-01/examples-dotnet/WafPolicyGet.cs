using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Cdn;
using Azure.ResourceManager.Cdn.Models;

// Generated from example definition: specification/cdn/resource-manager/Microsoft.Cdn/stable/2021-06-01/examples/WafPolicyGet.json
// this example is just showing the usage of "Policies_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this CdnWebApplicationFirewallPolicyResource created on azure
// for more information of creating CdnWebApplicationFirewallPolicyResource, please refer to the document of CdnWebApplicationFirewallPolicyResource
string subscriptionId = "subid";
string resourceGroupName = "rg1";
string policyName = "MicrosoftCdnWafPolicy";
ResourceIdentifier cdnWebApplicationFirewallPolicyResourceId = CdnWebApplicationFirewallPolicyResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, policyName);
CdnWebApplicationFirewallPolicyResource cdnWebApplicationFirewallPolicy = client.GetCdnWebApplicationFirewallPolicyResource(cdnWebApplicationFirewallPolicyResourceId);

// invoke the operation
CdnWebApplicationFirewallPolicyResource result = await cdnWebApplicationFirewallPolicy.GetAsync();

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
CdnWebApplicationFirewallPolicyData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
