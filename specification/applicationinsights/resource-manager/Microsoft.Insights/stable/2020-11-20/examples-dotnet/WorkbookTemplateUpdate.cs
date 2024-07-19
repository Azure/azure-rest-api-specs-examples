using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.ApplicationInsights.Models;
using Azure.ResourceManager.ApplicationInsights;

// Generated from example definition: specification/applicationinsights/resource-manager/Microsoft.Insights/stable/2020-11-20/examples/WorkbookTemplateUpdate.json
// this example is just showing the usage of "WorkbookTemplates_Update" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ApplicationInsightsWorkbookTemplateResource created on azure
// for more information of creating ApplicationInsightsWorkbookTemplateResource, please refer to the document of ApplicationInsightsWorkbookTemplateResource
string subscriptionId = "6b643656-33eb-422f-aee8-3ac145d124af";
string resourceGroupName = "my-resource-group";
string resourceName = "my-template-resource";
ResourceIdentifier applicationInsightsWorkbookTemplateResourceId = ApplicationInsightsWorkbookTemplateResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, resourceName);
ApplicationInsightsWorkbookTemplateResource applicationInsightsWorkbookTemplate = client.GetApplicationInsightsWorkbookTemplateResource(applicationInsightsWorkbookTemplateResourceId);

// invoke the operation
ApplicationInsightsWorkbookTemplatePatch patch = new ApplicationInsightsWorkbookTemplatePatch();
ApplicationInsightsWorkbookTemplateResource result = await applicationInsightsWorkbookTemplate.UpdateAsync(patch);

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
ApplicationInsightsWorkbookTemplateData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
