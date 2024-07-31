using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.AppService;

// Generated from example definition: specification/web/resource-manager/Microsoft.Web/stable/2023-12-01/examples/WorkflowTriggerHistories_Resubmit.json
// this example is just showing the usage of "WorkflowTriggerHistories_Resubmit" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this WorkflowTriggerHistoryResource created on azure
// for more information of creating WorkflowTriggerHistoryResource, please refer to the document of WorkflowTriggerHistoryResource
string subscriptionId = "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
string resourceGroupName = "testResourceGroup";
string name = "test-name";
string workflowName = "testWorkflowName";
string triggerName = "testTriggerName";
string historyName = "testHistoryName";
ResourceIdentifier workflowTriggerHistoryResourceId = WorkflowTriggerHistoryResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, name, workflowName, triggerName, historyName);
WorkflowTriggerHistoryResource workflowTriggerHistory = client.GetWorkflowTriggerHistoryResource(workflowTriggerHistoryResourceId);

// invoke the operation
await workflowTriggerHistory.ResubmitAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
