using Azure;
using Azure.ResourceManager;
using System;
using System.Collections.Generic;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Resources.Models;
using Azure.ResourceManager.WorkloadOrchestration.Models;
using Azure.ResourceManager.WorkloadOrchestration;

// Generated from example definition: 2025-06-01/WorkflowVersions_ListByWorkflow_MaximumSet_Gen.json
// this example is just showing the usage of "WorkflowVersion_ListByWorkflow" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this EdgeWorkflowResource created on azure
// for more information of creating EdgeWorkflowResource, please refer to the document of EdgeWorkflowResource
string subscriptionId = "9D54FE4C-00AF-4836-8F48-B6A9C4E47192";
string resourceGroupName = "rgconfigurationmanager";
string contextName = "testname";
string workflowName = "testname";
ResourceIdentifier edgeWorkflowResourceId = EdgeWorkflowResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, contextName, workflowName);
EdgeWorkflowResource edgeWorkflow = client.GetEdgeWorkflowResource(edgeWorkflowResourceId);

// get the collection of this EdgeWorkflowVersionResource
EdgeWorkflowVersionCollection collection = edgeWorkflow.GetEdgeWorkflowVersions();

// invoke the operation and iterate over the result
await foreach (EdgeWorkflowVersionResource item in collection.GetAllAsync())
{
    // the variable item is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    EdgeWorkflowVersionData resourceData = item.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}

Console.WriteLine("Succeeded");
