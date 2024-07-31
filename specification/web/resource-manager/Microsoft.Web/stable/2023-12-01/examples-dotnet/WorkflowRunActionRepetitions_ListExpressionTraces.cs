using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.AppService.Models;
using Azure.ResourceManager.AppService;

// Generated from example definition: specification/web/resource-manager/Microsoft.Web/stable/2023-12-01/examples/WorkflowRunActionRepetitions_ListExpressionTraces.json
// this example is just showing the usage of "WorkflowRunActionRepetitions_ListExpressionTraces" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this WorkflowRunActionRepetitionResource created on azure
// for more information of creating WorkflowRunActionRepetitionResource, please refer to the document of WorkflowRunActionRepetitionResource
string subscriptionId = "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
string resourceGroupName = "testResourceGroup";
string name = "test-name";
string workflowName = "testFlow";
string runName = "08586776228332053161046300351";
string actionName = "testAction";
string repetitionName = "000001";
ResourceIdentifier workflowRunActionRepetitionResourceId = WorkflowRunActionRepetitionResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, name, workflowName, runName, actionName, repetitionName);
WorkflowRunActionRepetitionResource workflowRunActionRepetition = client.GetWorkflowRunActionRepetitionResource(workflowRunActionRepetitionResourceId);

// invoke the operation and iterate over the result
await foreach (WorkflowExpressionRoot item in workflowRunActionRepetition.GetExpressionTracesAsync())
{
    Console.WriteLine($"Succeeded: {item}");
}

Console.WriteLine($"Succeeded");
