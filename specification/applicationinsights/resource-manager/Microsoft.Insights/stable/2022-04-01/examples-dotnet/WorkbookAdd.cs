using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.ApplicationInsights;
using Azure.ResourceManager.ApplicationInsights.Models;
using Azure.ResourceManager.Models;
using Azure.ResourceManager.Resources;

// Generated from example definition: specification/applicationinsights/resource-manager/Microsoft.Insights/stable/2022-04-01/examples/WorkbookAdd.json
// this example is just showing the usage of "Workbooks_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ResourceGroupResource created on azure
// for more information of creating ResourceGroupResource, please refer to the document of ResourceGroupResource
string subscriptionId = "6b643656-33eb-422f-aee8-3ac145d124af";
string resourceGroupName = "my-resource-group";
ResourceIdentifier resourceGroupResourceId = ResourceGroupResource.CreateResourceIdentifier(subscriptionId, resourceGroupName);
ResourceGroupResource resourceGroupResource = client.GetResourceGroupResource(resourceGroupResourceId);

// get the collection of this WorkbookResource
WorkbookCollection collection = resourceGroupResource.GetWorkbooks();

// invoke the operation
string resourceName = "deadb33f-5e0d-4064-8ebb-1a4ed0313eb2";
WorkbookData data = new WorkbookData(new AzureLocation("westus"))
{
    DisplayName = "Sample workbook",
    SerializedData = "{\"version\":\"Notebook/1.0\",\"items\":[{\"type\":1,\"content\":\"{\"json\":\"## New workbook\\r\\n---\\r\\n\\r\\nWelcome to your new workbook.  This area will display text formatted as markdown.\\r\\n\\r\\n\\r\\nWe've included a basic analytics query to get you started. Use the `Edit` button below each section to configure it or add more sections.\"}\",\"halfWidth\":null,\"conditionalVisibility\":null},{\"type\":3,\"content\":\"{\"version\":\"KqlItem/1.0\",\"query\":\"union withsource=TableName *\\n| summarize Count=count() by TableName\\n| render barchart\",\"showQuery\":false,\"size\":1,\"aggregation\":0,\"showAnnotations\":false}\",\"halfWidth\":null,\"conditionalVisibility\":null}],\"isLocked\":false}",
    Category = "workbook",
    Description = "Sample workbook",
    Kind = WorkbookSharedTypeKind.Shared,
    Tags =
    {
    ["TagSample01"] = "sample01",
    ["TagSample02"] = "sample02",
    },
};
string sourceId = "/subscriptions/6b643656-33eb-422f-aee8-3ac145d124af/resourcegroups/my-resource-group";
ArmOperation<WorkbookResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, resourceName, data, sourceId: sourceId);
WorkbookResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
WorkbookData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
