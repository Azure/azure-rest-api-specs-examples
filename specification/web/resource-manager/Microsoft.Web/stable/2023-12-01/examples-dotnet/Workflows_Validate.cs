using Azure;
using Azure.ResourceManager;
using System;
using System.Collections.Generic;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.AppService.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.AppService;

// Generated from example definition: specification/web/resource-manager/Microsoft.Web/stable/2023-12-01/examples/Workflows_Validate.json
// this example is just showing the usage of "Workflows_Validate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this WebSiteResource created on azure
// for more information of creating WebSiteResource, please refer to the document of WebSiteResource
string subscriptionId = "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
string resourceGroupName = "test-resource-group";
string name = "test-name";
ResourceIdentifier webSiteResourceId = WebSiteResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, name);
WebSiteResource webSite = client.GetWebSiteResource(webSiteResourceId);

// invoke the operation
string workflowName = "test-workflow";
WorkflowData data = new WorkflowData(new AzureLocation("placeholder"))
{
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
    Kind = AppServiceKind.Stateful,
};
await webSite.ValidateWorkflowAsync(workflowName, data);

Console.WriteLine($"Succeeded");
