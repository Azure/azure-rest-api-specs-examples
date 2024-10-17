using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Logic;

// Generated from example definition: specification/logic/resource-manager/Microsoft.Logic/stable/2019-05-01/examples/WorkflowRuns_Cancel.json
// this example is just showing the usage of "WorkflowRuns_Cancel" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this LogicWorkflowRunResource created on azure
// for more information of creating LogicWorkflowRunResource, please refer to the document of LogicWorkflowRunResource
string subscriptionId = "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
string resourceGroupName = "test-resource-group";
string workflowName = "test-workflow";
string runName = "08586676746934337772206998657CU22";
ResourceIdentifier logicWorkflowRunResourceId = LogicWorkflowRunResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, workflowName, runName);
LogicWorkflowRunResource logicWorkflowRun = client.GetLogicWorkflowRunResource(logicWorkflowRunResourceId);

// invoke the operation
await logicWorkflowRun.CancelAsync();

Console.WriteLine($"Succeeded");
