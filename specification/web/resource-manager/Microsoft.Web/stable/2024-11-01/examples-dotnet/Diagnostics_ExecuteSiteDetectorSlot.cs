using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.AppService.Models;
using Azure.ResourceManager.AppService;

// Generated from example definition: specification/web/resource-manager/Microsoft.Web/stable/2024-11-01/examples/Diagnostics_ExecuteSiteDetectorSlot.json
// this example is just showing the usage of "Diagnostics_ExecuteSiteDetector" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SiteDiagnosticDetectorResource created on azure
// for more information of creating SiteDiagnosticDetectorResource, please refer to the document of SiteDiagnosticDetectorResource
string subscriptionId = "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
string resourceGroupName = "Sample-WestUSResourceGroup";
string siteName = "SampleApp";
string diagnosticCategory = "availability";
string detectorName = "sitecrashes";
ResourceIdentifier siteDiagnosticDetectorResourceId = SiteDiagnosticDetectorResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, siteName, diagnosticCategory, detectorName);
SiteDiagnosticDetectorResource siteDiagnosticDetector = client.GetSiteDiagnosticDetectorResource(siteDiagnosticDetectorResourceId);

// invoke the operation
DiagnosticDetectorResponse result = await siteDiagnosticDetector.ExecuteAsync();

Console.WriteLine($"Succeeded: {result}");
