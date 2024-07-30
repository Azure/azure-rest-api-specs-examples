using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.AppService;

// Generated from example definition: specification/web/resource-manager/Microsoft.Web/stable/2023-12-01/examples/Diagnostics_GetSiteAnalysisSlot.json
// this example is just showing the usage of "Diagnostics_GetSiteAnalysis" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SiteDiagnosticResource created on azure
// for more information of creating SiteDiagnosticResource, please refer to the document of SiteDiagnosticResource
string subscriptionId = "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
string resourceGroupName = "Sample-WestUSResourceGroup";
string siteName = "SampleApp";
string diagnosticCategory = "availability";
ResourceIdentifier siteDiagnosticResourceId = SiteDiagnosticResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, siteName, diagnosticCategory);
SiteDiagnosticResource siteDiagnostic = client.GetSiteDiagnosticResource(siteDiagnosticResourceId);

// get the collection of this SiteDiagnosticAnalysisResource
SiteDiagnosticAnalysisCollection collection = siteDiagnostic.GetSiteDiagnosticAnalyses();

// invoke the operation
string analysisName = "appanalysis";
NullableResponse<SiteDiagnosticAnalysisResource> response = await collection.GetIfExistsAsync(analysisName);
SiteDiagnosticAnalysisResource result = response.HasValue ? response.Value : null;

if (result == null)
{
    Console.WriteLine($"Succeeded with null as result");
}
else
{
    // the variable result is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    WebSiteAnalysisDefinitionData resourceData = result.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}
