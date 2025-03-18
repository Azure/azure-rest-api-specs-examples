using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.ManagementGroups;
using Azure.ResourceManager.PolicyInsights.Models;
using Azure.ResourceManager.PolicyInsights;

// Generated from example definition: specification/policyinsights/resource-manager/Microsoft.PolicyInsights/stable/2019-10-01/examples/PolicyEvents_QueryManagementGroupScope.json
// this example is just showing the usage of "PolicyEvents_ListQueryResultsForManagementGroup" operation, for the dependent resources, they will have to be created separately.

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
PolicyEventType policyEventType = PolicyEventType.Default;
await foreach (PolicyEvent item in managementGroupResource.GetPolicyEventQueryResultsAsync(policyEventType))
{
    Console.WriteLine($"Succeeded: {item}");
}

Console.WriteLine("Succeeded");
