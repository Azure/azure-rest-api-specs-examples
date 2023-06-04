using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.ManagementGroups;
using Azure.ResourceManager.PolicyInsights;
using Azure.ResourceManager.PolicyInsights.Models;

// Generated from example definition: specification/policyinsights/resource-manager/Microsoft.PolicyInsights/stable/2019-10-01/examples/PolicyStates_SummarizeManagementGroupScope.json
// this example is just showing the usage of "PolicyStates_SummarizeForManagementGroup" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ManagementGroupResource created on azure
// for more information of creating ManagementGroupResource, please refer to the document of ManagementGroupResource
string managementGroupName = "myManagementGroup";
ResourceIdentifier managementGroupResourceId = ManagementGroupResource.CreateResourceIdentifier(managementGroupName);
ManagementGroupResource managementGroupResource = client.GetManagementGroupResource(managementGroupResourceId);

// invoke the operation and iterate over the result
PolicyStateSummaryType policyStateSummaryType = PolicyStateSummaryType.Latest;
await foreach (PolicySummary item in managementGroupResource.SummarizePolicyStatesAsync(policyStateSummaryType))
{
    Console.WriteLine($"Succeeded: {item}");
}

Console.WriteLine($"Succeeded");
