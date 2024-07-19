using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.FrontDoor.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.FrontDoor;

// Generated from example definition: specification/frontdoor/resource-manager/Microsoft.Network/stable/2024-02-01/examples/WafPolicyPatch.json
// this example is just showing the usage of "Policies_Update" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this FrontDoorWebApplicationFirewallPolicyResource created on azure
// for more information of creating FrontDoorWebApplicationFirewallPolicyResource, please refer to the document of FrontDoorWebApplicationFirewallPolicyResource
string subscriptionId = "subid";
string resourceGroupName = "rg1";
string policyName = "Policy1";
ResourceIdentifier frontDoorWebApplicationFirewallPolicyResourceId = FrontDoorWebApplicationFirewallPolicyResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, policyName);
FrontDoorWebApplicationFirewallPolicyResource frontDoorWebApplicationFirewallPolicy = client.GetFrontDoorWebApplicationFirewallPolicyResource(frontDoorWebApplicationFirewallPolicyResourceId);

// invoke the operation
FrontDoorWebApplicationFirewallPolicyPatch patch = new FrontDoorWebApplicationFirewallPolicyPatch()
{
    Tags =
    {
    ["key1"] = "value1",
    ["key2"] = "value2",
    },
};
ArmOperation<FrontDoorWebApplicationFirewallPolicyResource> lro = await frontDoorWebApplicationFirewallPolicy.UpdateAsync(WaitUntil.Completed, patch);
FrontDoorWebApplicationFirewallPolicyResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
FrontDoorWebApplicationFirewallPolicyData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
