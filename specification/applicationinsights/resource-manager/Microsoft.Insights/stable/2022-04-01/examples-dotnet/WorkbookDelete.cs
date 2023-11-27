using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.ApplicationInsights;
using Azure.ResourceManager.ApplicationInsights.Models;
using Azure.ResourceManager.Resources;

// Generated from example definition: specification/applicationinsights/resource-manager/Microsoft.Insights/stable/2022-04-01/examples/WorkbookDelete.json
// this example is just showing the usage of "Workbooks_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this WorkbookResource created on azure
// for more information of creating WorkbookResource, please refer to the document of WorkbookResource
string subscriptionId = "6b643656-33eb-422f-aee8-3ac145d124af";
string resourceGroupName = "my-resource-group";
string resourceName = "deadb33f-5e0d-4064-8ebb-1a4ed0313eb2";
ResourceIdentifier workbookResourceId = WorkbookResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, resourceName);
WorkbookResource workbook = client.GetWorkbookResource(workbookResourceId);

// invoke the operation
await workbook.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
