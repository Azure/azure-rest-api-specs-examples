using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.PolicyInsights.Models;
using Azure.ResourceManager.PolicyInsights;

// Generated from example definition: specification/policyinsights/resource-manager/Microsoft.PolicyInsights/stable/2024-10-01/examples/ComponentPolicyStates_QueryResourceScopeGroupByComponentTypeWithAggregate.json
// this example is just showing the usage of "ComponentPolicyStates_ListQueryResultsForResource" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// invoke the operation and iterate over the result
ResourceIdentifier scope = null;
ComponentPolicyStatesResource componentPolicyStatesResource = ComponentPolicyStatesResource.Latest;
ArmResourceGetQueryResultsForResourceComponentPolicyStatesOptions options = new ArmResourceGetQueryResultsForResourceComponentPolicyStatesOptions(componentPolicyStatesResource) { Filter = "policyAssignmentId eq '/subscriptions/e78961ba-36fe-4739-9212-e3031b4c8db7/providers/microsoft.authorization/policyassignments/560050f83dbb4a24974323f8'", Apply = "groupby((componentType,complianceState),aggregate($count as count))" };
await foreach (ComponentPolicyState item in client.GetQueryResultsForResourceComponentPolicyStatesAsync(scope, options))
{
    Console.WriteLine($"Succeeded: {item}");
}

Console.WriteLine("Succeeded");
