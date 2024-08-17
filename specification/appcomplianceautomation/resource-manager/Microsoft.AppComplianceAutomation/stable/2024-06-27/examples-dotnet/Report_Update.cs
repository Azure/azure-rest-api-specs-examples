using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.AppComplianceAutomation.Models;
using Azure.ResourceManager.AppComplianceAutomation;

// Generated from example definition: specification/appcomplianceautomation/resource-manager/Microsoft.AppComplianceAutomation/stable/2024-06-27/examples/Report_Update.json
// this example is just showing the usage of "Report_Update" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this AppComplianceReportResource created on azure
// for more information of creating AppComplianceReportResource, please refer to the document of AppComplianceReportResource
string reportName = "testReportName";
ResourceIdentifier appComplianceReportResourceId = AppComplianceReportResource.CreateResourceIdentifier(reportName);
AppComplianceReportResource appComplianceReport = client.GetAppComplianceReportResource(appComplianceReportResourceId);

// invoke the operation
AppComplianceReportPatch patch = new AppComplianceReportPatch();
ArmOperation<AppComplianceReportResource> lro = await appComplianceReport.UpdateAsync(WaitUntil.Completed, patch);
AppComplianceReportResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
AppComplianceReportData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
