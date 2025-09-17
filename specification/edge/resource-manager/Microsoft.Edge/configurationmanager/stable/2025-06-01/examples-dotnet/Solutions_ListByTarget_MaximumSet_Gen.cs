using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Resources.Models;
using Azure.ResourceManager.WorkloadOrchestration.Models;
using Azure.ResourceManager.WorkloadOrchestration;

// Generated from example definition: 2025-06-01/Solutions_ListByTarget_MaximumSet_Gen.json
// this example is just showing the usage of "Solution_ListByTarget" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this EdgeTargetResource created on azure
// for more information of creating EdgeTargetResource, please refer to the document of EdgeTargetResource
string subscriptionId = "9D54FE4C-00AF-4836-8F48-B6A9C4E47192";
string resourceGroupName = "rgconfigurationmanager";
string targetName = "testname";
ResourceIdentifier edgeTargetResourceId = EdgeTargetResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, targetName);
EdgeTargetResource edgeTarget = client.GetEdgeTargetResource(edgeTargetResourceId);

// get the collection of this EdgeSolutionResource
EdgeSolutionCollection collection = edgeTarget.GetEdgeSolutions();

// invoke the operation and iterate over the result
await foreach (EdgeSolutionResource item in collection.GetAllAsync())
{
    // the variable item is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    EdgeSolutionData resourceData = item.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}

Console.WriteLine("Succeeded");
