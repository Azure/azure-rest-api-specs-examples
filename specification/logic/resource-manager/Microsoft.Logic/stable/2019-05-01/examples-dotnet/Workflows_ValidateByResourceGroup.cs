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

// Generated from example definition: specification/logic/resource-manager/Microsoft.Logic/stable/2019-05-01/examples/Workflows_ValidateByResourceGroup.json
// this example is just showing the usage of "Workflows_ValidateByResourceGroup" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this LogicWorkflowResource created on azure
// for more information of creating LogicWorkflowResource, please refer to the document of LogicWorkflowResource
string subscriptionId = "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
string resourceGroupName = "test-resource-group";
string workflowName = "test-workflow";
ResourceIdentifier logicWorkflowResourceId = LogicWorkflowResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, workflowName);
LogicWorkflowResource logicWorkflow = client.GetLogicWorkflowResource(logicWorkflowResourceId);

// invoke the operation
LogicWorkflowData data = new LogicWorkflowData(new AzureLocation("brazilsouth"))
{
    IntegrationAccount = new LogicResourceReference()
    {
        Id = new ResourceIdentifier("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/test-resource-group/providers/Microsoft.Logic/integrationAccounts/test-integration-account"),
    },
    Definition = BinaryData.FromObjectAsJson(new Dictionary<string, object>()
    {
        ["$schema"] = "https://schema.management.azure.com/providers/Microsoft.Logic/schemas/2016-06-01/workflowdefinition.json#",
        ["actions"] = new Dictionary<string, object>()
        {
        },
        ["contentVersion"] = "1.0.0.0",
        ["outputs"] = new Dictionary<string, object>()
        {
        },
        ["parameters"] = new Dictionary<string, object>()
        {
        },
        ["triggers"] = new Dictionary<string, object>()
        {
        }
    }),
    Tags =
    {
    },
};
await logicWorkflow.ValidateByResourceGroupAsync(data);

Console.WriteLine($"Succeeded");
