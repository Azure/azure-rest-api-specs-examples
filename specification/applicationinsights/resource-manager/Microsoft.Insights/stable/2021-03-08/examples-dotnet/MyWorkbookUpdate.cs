using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.ApplicationInsights;
using Azure.ResourceManager.ApplicationInsights.Models;
using Azure.ResourceManager.Resources;

// Generated from example definition: specification/applicationinsights/resource-manager/Microsoft.Insights/stable/2021-03-08/examples/MyWorkbookUpdate.json
// this example is just showing the usage of "MyWorkbooks_Update" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this MyWorkbookResource created on azure
// for more information of creating MyWorkbookResource, please refer to the document of MyWorkbookResource
string subscriptionId = "6b643656-33eb-422f-aee8-3ac145d124af";
string resourceGroupName = "my-resource-group";
string resourceName = "deadb33f-5e0d-4064-8ebb-1a4ed0313eb2";
ResourceIdentifier myWorkbookResourceId = MyWorkbookResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, resourceName);
MyWorkbookResource myWorkbook = client.GetMyWorkbookResource(myWorkbookResourceId);

// invoke the operation
MyWorkbookData data = new MyWorkbookData()
{
    Kind = ApplicationInsightsKind.User,
    DisplayName = "Blah Blah Blah",
    SerializedData = "{\"version\":\"Notebook/1.0\",\"items\":[{\"type\":1,\"content\":\"{\"json\":\"## New workbook\\r\\n---\\r\\n\\r\\nWelcome to your new workbook.  This area will display text formatted as markdown.\\r\\n\\r\\n\\r\\nWe've included a basic analytics query to get you started. Use the `Edit` button below each section to configure it or add more sections.\"}\",\"halfWidth\":null,\"conditionalVisibility\":null},{\"type\":3,\"content\":\"{\"version\":\"KqlItem/1.0\",\"query\":\"union withsource=TableName *\\n| summarize Count=count() by TableName\\n| render barchart\",\"showQuery\":false,\"size\":1,\"aggregation\":0,\"showAnnotations\":false}\",\"halfWidth\":null,\"conditionalVisibility\":null}],\"isLocked\":false}",
    Version = "ME",
    Category = "workbook",
    SourceId = "/subscriptions/00000000-0000-0000-0000-00000000/resourceGroups/MyGroup/providers/Microsoft.Web/sites/MyTestApp-CodeLens",
    StorageUri = null,
    Name = "deadb33f-8bee-4d3b-a059-9be8dac93960",
    Location = new AzureLocation("west us"),
    Tags =
    {
    ["0"] = "TagSample01",
    ["1"] = "TagSample02",
    },
};
MyWorkbookResource result = await myWorkbook.UpdateAsync(data);

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
MyWorkbookData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
