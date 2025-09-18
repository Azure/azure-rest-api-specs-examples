using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Resources.Models;
using Azure.ResourceManager.WorkloadOrchestration.Models;
using Azure.ResourceManager.WorkloadOrchestration;

// Generated from example definition: 2025-06-01/Executions_CreateOrUpdate_MaximumSet_Gen.json
// this example is just showing the usage of "Execution_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this EdgeWorkflowVersionResource created on azure
// for more information of creating EdgeWorkflowVersionResource, please refer to the document of EdgeWorkflowVersionResource
string subscriptionId = "EE6D9590-0D52-4B1C-935C-FE49DBF838EB";
string resourceGroupName = "rgconfigurationmanager";
string contextName = "abcde";
string workflowName = "abcde";
string versionName = "abcde";
ResourceIdentifier edgeWorkflowVersionResourceId = EdgeWorkflowVersionResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, contextName, workflowName, versionName);
EdgeWorkflowVersionResource edgeWorkflowVersion = client.GetEdgeWorkflowVersionResource(edgeWorkflowVersionResourceId);

// get the collection of this EdgeExecutionResource
EdgeExecutionCollection collection = edgeWorkflowVersion.GetEdgeExecutions();

// invoke the operation
string executionName = "abcde";
EdgeExecutionData data = new EdgeExecutionData
{
    Properties = new EdgeExecutionProperties("souenlqwltljsojdcbpc")
    {
        Specification = { },
    },
    ExtendedLocation = new ExtendedLocation
    {
        Name = "ugf",
    },
};
ArmOperation<EdgeExecutionResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, executionName, data);
EdgeExecutionResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
EdgeExecutionData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
