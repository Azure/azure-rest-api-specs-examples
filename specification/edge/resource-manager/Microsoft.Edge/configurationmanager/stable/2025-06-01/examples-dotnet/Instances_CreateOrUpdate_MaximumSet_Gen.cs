using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Resources.Models;
using Azure.ResourceManager.WorkloadOrchestration.Models;
using Azure.ResourceManager.WorkloadOrchestration;

// Generated from example definition: 2025-06-01/Instances_CreateOrUpdate_MaximumSet_Gen.json
// this example is just showing the usage of "Instance_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this EdgeSolutionResource created on azure
// for more information of creating EdgeSolutionResource, please refer to the document of EdgeSolutionResource
string subscriptionId = "9D54FE4C-00AF-4836-8F48-B6A9C4E47192";
string resourceGroupName = "rgconfigurationmanager";
string targetName = "testname";
string solutionName = "testname";
ResourceIdentifier edgeSolutionResourceId = EdgeSolutionResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, targetName, solutionName);
EdgeSolutionResource edgeSolution = client.GetEdgeSolutionResource(edgeSolutionResourceId);

// get the collection of this EdgeDeploymentInstanceResource
EdgeDeploymentInstanceCollection collection = edgeSolution.GetEdgeDeploymentInstances();

// invoke the operation
string instanceName = "testname";
EdgeDeploymentInstanceData data = new EdgeDeploymentInstanceData
{
    Properties = new EdgeDeploymentInstanceProperties("acpddbkfclsgxg", "eguutiftuxrsavvckjrv")
    {
        ActiveState = InstanceActiveState.Active,
        ReconciliationPolicy = new InstanceReconciliationPolicy(InstanceReconciliationState.Inactive, "szucgzdbydcowvhprhx"),
        SolutionScope = "testname",
    },
    ExtendedLocation = new ExtendedLocation
    {
        Name = "szjrwimeqyiue",
    },
};
ArmOperation<EdgeDeploymentInstanceResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, instanceName, data);
EdgeDeploymentInstanceResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
EdgeDeploymentInstanceData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
