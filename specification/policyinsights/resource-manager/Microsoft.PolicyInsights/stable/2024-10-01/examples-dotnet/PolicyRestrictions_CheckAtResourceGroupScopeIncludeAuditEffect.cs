using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.PolicyInsights.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.PolicyInsights;

// Generated from example definition: specification/policyinsights/resource-manager/Microsoft.PolicyInsights/stable/2024-10-01/examples/PolicyRestrictions_CheckAtResourceGroupScopeIncludeAuditEffect.json
// this example is just showing the usage of "PolicyRestrictions_CheckAtResourceGroupScope" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ResourceGroupResource created on azure
// for more information of creating ResourceGroupResource, please refer to the document of ResourceGroupResource
string subscriptionId = "35ee058e-5fa0-414c-8145-3ebb8d09b6e2";
string resourceGroupName = "vmRg";
ResourceIdentifier resourceGroupResourceId = ResourceGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName);
ResourceGroupResource resourceGroupResource = client.GetResourceGroupResource(resourceGroupResourceId);

// invoke the operation
CheckPolicyRestrictionsContent content = new CheckPolicyRestrictionsContent(new CheckRestrictionsResourceDetails(BinaryData.FromObjectAsJson(new
{
    type = "Microsoft.Compute/virtualMachines",
    properties = new
    {
        priority = "Spot",
    },
}))
{
    ApiVersion = "2019-12-01",
})
{
    PendingFields = {new PendingField("name")
    {
    Values = {"myVMName"},
    }, new PendingField("location")
    {
    Values = {"eastus", "westus", "westus2", "westeurope"},
    }, new PendingField("tags")},
    IncludeAuditEffect = true,
};
CheckPolicyRestrictionsResult result = await resourceGroupResource.CheckPolicyRestrictionsAsync(content);

Console.WriteLine($"Succeeded: {result}");
