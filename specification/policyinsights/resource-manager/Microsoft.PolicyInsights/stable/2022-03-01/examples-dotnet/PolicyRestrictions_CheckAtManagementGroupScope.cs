using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.ManagementGroups;
using Azure.ResourceManager.PolicyInsights;
using Azure.ResourceManager.PolicyInsights.Models;

// Generated from example definition: specification/policyinsights/resource-manager/Microsoft.PolicyInsights/stable/2022-03-01/examples/PolicyRestrictions_CheckAtManagementGroupScope.json
// this example is just showing the usage of "PolicyRestrictions_CheckAtManagementGroupScope" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ManagementGroupResource created on azure
// for more information of creating ManagementGroupResource, please refer to the document of ManagementGroupResource
string managementGroupId = "financeMg";
ResourceIdentifier managementGroupResourceId = ManagementGroupResource.CreateResourceIdentifier(managementGroupId);
ManagementGroupResource managementGroupResource = client.GetManagementGroupResource(managementGroupResourceId);

// invoke the operation
CheckManagementGroupPolicyRestrictionsContent content = new CheckManagementGroupPolicyRestrictionsContent()
{
    PendingFields =
    {
    new PendingField("type")
    },
};
CheckPolicyRestrictionsResult result = await managementGroupResource.CheckPolicyRestrictionsAsync(content);

Console.WriteLine($"Succeeded: {result}");
