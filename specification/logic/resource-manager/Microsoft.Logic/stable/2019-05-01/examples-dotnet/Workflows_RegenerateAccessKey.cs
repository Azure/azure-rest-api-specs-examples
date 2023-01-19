using System;
using System.Collections.Generic;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.Logic;
using Azure.ResourceManager.Logic.Models;
using Azure.ResourceManager.Resources;

// Generated from example definition: specification/logic/resource-manager/Microsoft.Logic/stable/2019-05-01/examples/Workflows_RegenerateAccessKey.json
// this example is just showing the usage of "Workflows_RegenerateAccessKey" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this LogicWorkflowResource created on azure
// for more information of creating LogicWorkflowResource, please refer to the document of LogicWorkflowResource
string subscriptionId = "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
string resourceGroupName = "testResourceGroup";
string workflowName = "testWorkflowName";
ResourceIdentifier logicWorkflowResourceId = LogicWorkflowResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, workflowName);
LogicWorkflowResource logicWorkflow = client.GetLogicWorkflowResource(logicWorkflowResourceId);

// invoke the operation
LogicWorkflowRegenerateActionContent content = new LogicWorkflowRegenerateActionContent()
{
    KeyType = LogicKeyType.Primary,
};
await logicWorkflow.RegenerateAccessKeyAsync(content);

Console.WriteLine($"Succeeded");
