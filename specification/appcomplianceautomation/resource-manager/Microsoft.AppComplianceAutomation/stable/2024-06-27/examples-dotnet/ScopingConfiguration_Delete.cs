using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.AppComplianceAutomation;

// Generated from example definition: specification/appcomplianceautomation/resource-manager/Microsoft.AppComplianceAutomation/stable/2024-06-27/examples/ScopingConfiguration_Delete.json
// this example is just showing the usage of "ScopingConfiguration_Delete" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this AppComplianceReportScopingConfigurationResource created on azure
// for more information of creating AppComplianceReportScopingConfigurationResource, please refer to the document of AppComplianceReportScopingConfigurationResource
string reportName = "testReportName";
string scopingConfigurationName = "default";
ResourceIdentifier appComplianceReportScopingConfigurationResourceId = AppComplianceReportScopingConfigurationResource.CreateResourceIdentifier(reportName, scopingConfigurationName);
AppComplianceReportScopingConfigurationResource appComplianceReportScopingConfiguration = client.GetAppComplianceReportScopingConfigurationResource(appComplianceReportScopingConfigurationResourceId);

// invoke the operation
await appComplianceReportScopingConfiguration.DeleteAsync(WaitUntil.Completed);

Console.WriteLine($"Succeeded");
