using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.AppService;

// Generated from example definition: specification/web/resource-manager/Microsoft.Web/stable/2024-11-01/examples/WorkflowTriggerHistories_List.json
// this example is just showing the usage of "WorkflowTriggerHistories_List" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this WorkflowTriggerResource created on azure
// for more information of creating WorkflowTriggerResource, please refer to the document of WorkflowTriggerResource
string subscriptionId = "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
string resourceGroupName = "testResourceGroup";
string name = "test-name";
string workflowName = "testWorkflowName";
string triggerName = "testTriggerName";
ResourceIdentifier workflowTriggerResourceId = WorkflowTriggerResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, name, workflowName, triggerName);
WorkflowTriggerResource workflowTrigger = client.GetWorkflowTriggerResource(workflowTriggerResourceId);

// get the collection of this WorkflowTriggerHistoryResource
WorkflowTriggerHistoryCollection collection = workflowTrigger.GetWorkflowTriggerHistories();

// invoke the operation and iterate over the result
await foreach (WorkflowTriggerHistoryResource item in collection.GetAllAsync())
{
    // the variable item is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    WorkflowTriggerHistoryData resourceData = item.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}

Console.WriteLine("Succeeded");
