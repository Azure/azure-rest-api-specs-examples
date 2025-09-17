using Azure;
using Azure.ResourceManager;
using System;
using System.Collections.Generic;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.WorkloadOrchestration.Models;
using Azure.ResourceManager.WorkloadOrchestration;

// Generated from example definition: 2025-06-01/SolutionVersions_Update_MaximumSet_Gen.json
// this example is just showing the usage of "SolutionVersion_Update" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this EdgeSolutionVersionResource created on azure
// for more information of creating EdgeSolutionVersionResource, please refer to the document of EdgeSolutionVersionResource
string subscriptionId = "9D54FE4C-00AF-4836-8F48-B6A9C4E47192";
string resourceGroupName = "rgconfigurationmanager";
string targetName = "testname";
string solutionName = "testname";
string solutionVersionName = "testname";
ResourceIdentifier edgeSolutionVersionResourceId = EdgeSolutionVersionResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, targetName, solutionName, solutionVersionName);
EdgeSolutionVersionResource edgeSolutionVersion = client.GetEdgeSolutionVersionResource(edgeSolutionVersionResourceId);

// invoke the operation
EdgeSolutionVersionData data = new EdgeSolutionVersionData
{
    Properties = new EdgeSolutionVersionProperties(new Dictionary<string, BinaryData>()),
};
ArmOperation<EdgeSolutionVersionResource> lro = await edgeSolutionVersion.UpdateAsync(WaitUntil.Completed, data);
EdgeSolutionVersionResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
EdgeSolutionVersionData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
