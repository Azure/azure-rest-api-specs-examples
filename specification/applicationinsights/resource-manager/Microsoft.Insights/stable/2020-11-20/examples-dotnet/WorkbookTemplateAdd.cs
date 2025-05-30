using Azure;
using Azure.ResourceManager;
using System;
using System.Collections.Generic;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.ApplicationInsights.Models;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.ApplicationInsights;

// Generated from example definition: specification/applicationinsights/resource-manager/Microsoft.Insights/stable/2020-11-20/examples/WorkbookTemplateAdd.json
// this example is just showing the usage of "WorkbookTemplates_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ResourceGroupResource created on azure
// for more information of creating ResourceGroupResource, please refer to the document of ResourceGroupResource
string subscriptionId = "subid";
string resourceGroupName = "my-resource-group";
ResourceIdentifier resourceGroupResourceId = ResourceGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName);
ResourceGroupResource resourceGroupResource = client.GetResourceGroupResource(resourceGroupResourceId);

// get the collection of this ApplicationInsightsWorkbookTemplateResource
ApplicationInsightsWorkbookTemplateCollection collection = resourceGroupResource.GetApplicationInsightsWorkbookTemplates();

// invoke the operation
string resourceName = "testtemplate2";
ApplicationInsightsWorkbookTemplateData data = new ApplicationInsightsWorkbookTemplateData(new AzureLocation("west us"))
{
    Priority = 1,
    Author = "Contoso",
    TemplateData = BinaryData.FromObjectAsJson(new Dictionary<string, object>
    {
        ["$schema"] = "https://github.com/Microsoft/Application-Insights-Workbooks/blob/master/schema/workbook.json",
        ["items"] = new object[]
{
new
{
name = "text - 2",
type = "1",
content = new
{
json = "## New workbook\n---\n\nWelcome to your new workbook.  This area will display text formatted as markdown.\n\n\nWe've included a basic analytics query to get you started. Use the `Edit` button below each section to configure it or add more sections.",
},
},
new
{
name = "query - 2",
type = "3",
content = new
{
exportToExcelOptions = "visible",
query = "union withsource=TableName *\n| summarize Count=count() by TableName\n| render barchart",
queryType = "0",
resourceType = "microsoft.operationalinsights/workspaces",
size = "1",
version = "KqlItem/1.0",
},
}
},
        ["styleSettings"] = new object(),
        ["version"] = "Notebook/1.0"
    }),
    Galleries = {new WorkbookTemplateGallery
    {
    Name = "Simple Template",
    Category = "Failures",
    WorkbookType = "tsg",
    Order = 100,
    ResourceType = "microsoft.insights/components",
    }},
    Tags = { },
};
ArmOperation<ApplicationInsightsWorkbookTemplateResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, resourceName, data);
ApplicationInsightsWorkbookTemplateResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
ApplicationInsightsWorkbookTemplateData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
