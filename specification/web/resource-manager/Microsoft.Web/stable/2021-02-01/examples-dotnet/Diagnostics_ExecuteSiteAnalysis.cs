using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.AppService;
using Azure.ResourceManager.AppService.Models;

// Generated from example definition: specification/web/resource-manager/Microsoft.Web/stable/2021-02-01/examples/Diagnostics_ExecuteSiteAnalysis.json
// this example is just showing the usage of "Diagnostics_ExecuteSiteAnalysisSlot" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SiteSlotDiagnosticAnalysisResource created on azure
// for more information of creating SiteSlotDiagnosticAnalysisResource, please refer to the document of SiteSlotDiagnosticAnalysisResource
string subscriptionId = "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
string resourceGroupName = "Sample-WestUSResourceGroup";
string siteName = "SampleApp";
string slot = "Production";
string diagnosticCategory = "availability";
string analysisName = "apprestartanalyses";
ResourceIdentifier siteSlotDiagnosticAnalysisResourceId = SiteSlotDiagnosticAnalysisResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, siteName, slot, diagnosticCategory, analysisName);
SiteSlotDiagnosticAnalysisResource siteSlotDiagnosticAnalysis = client.GetSiteSlotDiagnosticAnalysisResource(siteSlotDiagnosticAnalysisResourceId);

// invoke the operation
DiagnosticAnalysis result = await siteSlotDiagnosticAnalysis.ExecuteSiteAnalysisSlotAsync();

Console.WriteLine($"Succeeded: {result}");
