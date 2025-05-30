using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.FrontDoor.Models;
using Azure.ResourceManager.FrontDoor;

// Generated from example definition: specification/frontdoor/resource-manager/Microsoft.Network/stable/2025-03-01/examples/WafPolicyDelete.json
// this example is just showing the usage of "Policies_Delete" operation, for the dependent resources, they will have to be created separately.

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
await frontDoorWebApplicationFirewallPolicy.DeleteAsync(WaitUntil.Completed);

Console.WriteLine("Succeeded");
