using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.ApplicationInsights;

// Generated from example definition: specification/applicationinsights/resource-manager/Microsoft.Insights/stable/2023-06-01/examples/WorkbookRevisionGet.json
// this example is just showing the usage of "Workbooks_RevisionGet" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ApplicationInsightsWorkbookResource created on azure
// for more information of creating ApplicationInsightsWorkbookResource, please refer to the document of ApplicationInsightsWorkbookResource
string subscriptionId = "6b643656-33eb-422f-aee8-3ac145d124af";
string resourceGroupName = "my-resource-group";
string resourceName = "deadb33f-5e0d-4064-8ebb-1a4ed0313eb2";
ResourceIdentifier applicationInsightsWorkbookResourceId = ApplicationInsightsWorkbookResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, resourceName);
ApplicationInsightsWorkbookResource applicationInsightsWorkbook = client.GetApplicationInsightsWorkbookResource(applicationInsightsWorkbookResourceId);

// get the collection of this ApplicationInsightsWorkbookRevisionResource
ApplicationInsightsWorkbookRevisionCollection collection = applicationInsightsWorkbook.GetApplicationInsightsWorkbookRevisions();

// invoke the operation
string revisionId = "1e2f8435b98248febee70c64ac22e1ab";
NullableResponse<ApplicationInsightsWorkbookRevisionResource> response = await collection.GetIfExistsAsync(revisionId);
ApplicationInsightsWorkbookRevisionResource result = response.HasValue ? response.Value : null;

if (result == null)
{
    Console.WriteLine($"Succeeded with null as result");
}
else
{
    // the variable result is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    ApplicationInsightsWorkbookData resourceData = result.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}
