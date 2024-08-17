using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.AppComplianceAutomation.Models;
using Azure.ResourceManager.AppComplianceAutomation;

// Generated from example definition: specification/appcomplianceautomation/resource-manager/Microsoft.AppComplianceAutomation/stable/2024-06-27/examples/Evidence_Delete.json
// this example is just showing the usage of "Evidence_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this AppComplianceReportEvidenceResource created on azure
// for more information of creating AppComplianceReportEvidenceResource, please refer to the document of AppComplianceReportEvidenceResource
string reportName = "testReportName";
string evidenceName = "evidence1";
ResourceIdentifier appComplianceReportEvidenceResourceId = AppComplianceReportEvidenceResource.CreateResourceIdentifier(reportName, evidenceName);
AppComplianceReportEvidenceResource appComplianceReportEvidence = client.GetAppComplianceReportEvidenceResource(appComplianceReportEvidenceResourceId);

// invoke the operation
await appComplianceReportEvidence.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
