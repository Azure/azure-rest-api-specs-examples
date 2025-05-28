using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Network;

// Generated from example definition: specification/network/resource-manager/Microsoft.Network/stable/2024-07-01/examples/FirewallPolicySignatureOverridesGet.json
// this example is just showing the usage of "FirewallPolicyIdpsSignaturesOverrides_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this PolicySignaturesOverridesForIdpsResource created on azure
// for more information of creating PolicySignaturesOverridesForIdpsResource, please refer to the document of PolicySignaturesOverridesForIdpsResource
string subscriptionId = "e747cc13-97d4-4a79-b463-42d7f4e558f2";
string resourceGroupName = "rg1";
string firewallPolicyName = "firewallPolicy";
ResourceIdentifier policySignaturesOverridesForIdpsResourceId = PolicySignaturesOverridesForIdpsResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, firewallPolicyName);
PolicySignaturesOverridesForIdpsResource policySignaturesOverridesForIdps = client.GetPolicySignaturesOverridesForIdpsResource(policySignaturesOverridesForIdpsResourceId);

// invoke the operation
PolicySignaturesOverridesForIdpsResource result = await policySignaturesOverridesForIdps.GetAsync();

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
PolicySignaturesOverridesForIdpsData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
