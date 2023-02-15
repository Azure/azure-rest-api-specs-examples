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

// Generated from example definition: specification/logic/resource-manager/Microsoft.Logic/stable/2019-05-01/examples/Workflows_CreateOrUpdate.json
// this example is just showing the usage of "Workflows_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

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
            ["Find_pet_by_ID"] = new Dictionary<string, object>()
            {
                ["type"] = "ApiConnection",
                ["inputs"] = new Dictionary<string, object>()
                {
                    ["path"] = "/pet/@{encodeURIComponent('1')}",
                    ["method"] = "get",
                    ["host"] = new Dictionary<string, object>()
                    {
                        ["connection"] = new Dictionary<string, object>()
                        {
                            ["name"] = "@parameters('$connections')['test-custom-connector']['connectionId']"
                        }
                    }
                },
                ["runAfter"] = new Dictionary<string, object>()
                {
                }
            }
        },
        ["contentVersion"] = "1.0.0.0",
        ["outputs"] = new Dictionary<string, object>()
        {
        },
        ["parameters"] = new Dictionary<string, object>()
        {
            ["$connections"] = new Dictionary<string, object>()
            {
                ["type"] = "Object",
                ["defaultValue"] = new Dictionary<string, object>()
                {
                }
            }
        },
        ["triggers"] = new Dictionary<string, object>()
        {
            ["manual"] = new Dictionary<string, object>()
            {
                ["type"] = "Request",
                ["inputs"] = new Dictionary<string, object>()
                {
                    ["schema"] = new Dictionary<string, object>()
                    {
                    }
                },
                ["kind"] = "Http"
            }
        }
    }),
    Parameters =
    {
    ["$connections"] = new LogicWorkflowParameterInfo()
    {
    Value = BinaryData.FromObjectAsJson(new Dictionary<string, object>()
    {
    ["test-custom-connector"] = new Dictionary<string, object>()
    {
    ["connectionId"] = "/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/test-resource-group/providers/Microsoft.Web/connections/test-custom-connector",
    ["connectionName"] = "test-custom-connector",
    ["id"] = "/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/providers/Microsoft.Web/locations/brazilsouth/managedApis/test-custom-connector"}}),
    },
    },
    Tags =
    {
    },
};
ArmOperation<LogicWorkflowResource> lro = await logicWorkflow.UpdateAsync(WaitUntil.Completed, data);
LogicWorkflowResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
LogicWorkflowData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
