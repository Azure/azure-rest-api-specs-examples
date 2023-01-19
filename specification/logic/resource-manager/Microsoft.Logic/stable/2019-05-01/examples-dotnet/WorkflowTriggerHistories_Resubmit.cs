using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Logic;

// Generated from example definition: specification/logic/resource-manager/Microsoft.Logic/stable/2019-05-01/examples/WorkflowTriggerHistories_Resubmit.json
// this example is just showing the usage of "WorkflowTriggerHistories_Resubmit" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this LogicWorkflowTriggerHistoryResource created on azure
// for more information of creating LogicWorkflowTriggerHistoryResource, please refer to the document of LogicWorkflowTriggerHistoryResource
string subscriptionId = "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
string resourceGroupName = "testResourceGroup";
string workflowName = "testWorkflowName";
string triggerName = "testTriggerName";
string historyName = "testHistoryName";
ResourceIdentifier logicWorkflowTriggerHistoryResourceId = LogicWorkflowTriggerHistoryResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, workflowName, triggerName, historyName);
LogicWorkflowTriggerHistoryResource logicWorkflowTriggerHistory = client.GetLogicWorkflowTriggerHistoryResource(logicWorkflowTriggerHistoryResourceId);

// invoke the operation
await logicWorkflowTriggerHistory.ResubmitAsync();

Console.WriteLine($"Succeeded");
