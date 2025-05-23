using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Logic.Models;
using Azure.ResourceManager.Logic;

// Generated from example definition: specification/logic/resource-manager/Microsoft.Logic/stable/2019-05-01/examples/WorkflowVersionTriggers_ListCallbackUrl.json
// this example is just showing the usage of "WorkflowVersionTriggers_ListCallbackUrl" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this LogicWorkflowVersionResource created on azure
// for more information of creating LogicWorkflowVersionResource, please refer to the document of LogicWorkflowVersionResource
string subscriptionId = "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
string resourceGroupName = "testResourceGroup";
string workflowName = "testWorkflowName";
string versionId = "testWorkflowVersionId";
ResourceIdentifier logicWorkflowVersionResourceId = LogicWorkflowVersionResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, workflowName, versionId);
LogicWorkflowVersionResource logicWorkflowVersion = client.GetLogicWorkflowVersionResource(logicWorkflowVersionResourceId);

// invoke the operation
string triggerName = "testTriggerName";
ListOperationCallbackUrlParameterInfo info = new ListOperationCallbackUrlParameterInfo()
{
    NotAfter = DateTimeOffset.Parse("2017-03-05T08:00:00Z"),
    KeyType = LogicKeyType.Primary,
};
LogicWorkflowTriggerCallbackUri result = await logicWorkflowVersion.GetCallbackUrlWorkflowVersionTriggerAsync(triggerName, info: info);

Console.WriteLine($"Succeeded: {result}");
