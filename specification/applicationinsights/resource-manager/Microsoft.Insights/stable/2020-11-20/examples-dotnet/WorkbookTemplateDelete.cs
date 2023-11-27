using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.ApplicationInsights;
using Azure.ResourceManager.ApplicationInsights.Models;

// Generated from example definition: specification/applicationinsights/resource-manager/Microsoft.Insights/stable/2020-11-20/examples/WorkbookTemplateDelete.json
// this example is just showing the usage of "WorkbookTemplates_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this WorkbookTemplateResource created on azure
// for more information of creating WorkbookTemplateResource, please refer to the document of WorkbookTemplateResource
string subscriptionId = "subid";
string resourceGroupName = "my-resource-group";
string resourceName = "my-template-resource";
ResourceIdentifier workbookTemplateResourceId = WorkbookTemplateResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, resourceName);
WorkbookTemplateResource workbookTemplate = client.GetWorkbookTemplateResource(workbookTemplateResourceId);

// invoke the operation
await workbookTemplate.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
