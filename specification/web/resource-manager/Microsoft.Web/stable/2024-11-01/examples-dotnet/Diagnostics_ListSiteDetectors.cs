using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.AppService;

// Generated from example definition: specification/web/resource-manager/Microsoft.Web/stable/2024-11-01/examples/Diagnostics_ListSiteDetectors.json
// this example is just showing the usage of "Diagnostics_ListSiteDetectors" operation, for the dependent resources, they will have to be created separately.

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

// get the collection of this SiteDiagnosticDetectorResource
SiteDiagnosticDetectorCollection collection = siteDiagnostic.GetSiteDiagnosticDetectors();

// invoke the operation and iterate over the result
await foreach (SiteDiagnosticDetectorResource item in collection.GetAllAsync())
{
    // the variable item is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    DetectorDefinitionResourceData resourceData = item.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}

Console.WriteLine("Succeeded");
