using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Network.Models;
using Azure.ResourceManager.Network;

// Generated from example definition: specification/network/resource-manager/Microsoft.Network/stable/2024-07-01/examples/ReachabilityAnalysisRunGet.json
// this example is just showing the usage of "ReachabilityAnalysisRuns_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this NetworkVerifierWorkspaceResource created on azure
// for more information of creating NetworkVerifierWorkspaceResource, please refer to the document of NetworkVerifierWorkspaceResource
string subscriptionId = "00000000-0000-0000-0000-000000000000";
string resourceGroupName = "rg1";
string networkManagerName = "testNetworkManager";
string workspaceName = "testWorkspace";
ResourceIdentifier networkVerifierWorkspaceResourceId = NetworkVerifierWorkspaceResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, networkManagerName, workspaceName);
NetworkVerifierWorkspaceResource networkVerifierWorkspace = client.GetNetworkVerifierWorkspaceResource(networkVerifierWorkspaceResourceId);

// get the collection of this ReachabilityAnalysisRunResource
ReachabilityAnalysisRunCollection collection = networkVerifierWorkspace.GetReachabilityAnalysisRuns();

// invoke the operation
string reachabilityAnalysisRunName = "testAnalysisRunName";
NullableResponse<ReachabilityAnalysisRunResource> response = await collection.GetIfExistsAsync(reachabilityAnalysisRunName);
ReachabilityAnalysisRunResource result = response.HasValue ? response.Value : null;

if (result == null)
{
    Console.WriteLine("Succeeded with null as result");
}
else
{
    // the variable result is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    ReachabilityAnalysisRunData resourceData = result.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}
