using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.AppComplianceAutomation.Models;
using Azure.ResourceManager.AppComplianceAutomation;

// Generated from example definition: specification/appcomplianceautomation/resource-manager/Microsoft.AppComplianceAutomation/stable/2024-06-27/examples/Evidence_ListByReport.json
// this example is just showing the usage of "Evidence_ListByReport" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this AppComplianceReportResource created on azure
// for more information of creating AppComplianceReportResource, please refer to the document of AppComplianceReportResource
string reportName = "reportName";
ResourceIdentifier appComplianceReportResourceId = AppComplianceReportResource.CreateResourceIdentifier(reportName);
AppComplianceReportResource appComplianceReport = client.GetAppComplianceReportResource(appComplianceReportResourceId);

// get the collection of this AppComplianceReportEvidenceResource
AppComplianceReportEvidenceCollection collection = appComplianceReport.GetAppComplianceReportEvidences();

// invoke the operation and iterate over the result
AppComplianceReportEvidenceCollectionGetAllOptions options = new AppComplianceReportEvidenceCollectionGetAllOptions() { };
await foreach (AppComplianceReportEvidenceResource item in collection.GetAllAsync(options))
{
    // the variable item is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    AppComplianceReportEvidenceData resourceData = item.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}

Console.WriteLine($"Succeeded");
