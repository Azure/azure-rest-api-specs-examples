using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.AppComplianceAutomation;

// Generated from example definition: specification/appcomplianceautomation/resource-manager/Microsoft.AppComplianceAutomation/preview/2022-11-16-preview/examples/Snapshot_Get.json
// this example is just showing the usage of "Snapshot_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ReportResource created on azure
// for more information of creating ReportResource, please refer to the document of ReportResource
string reportName = "testReportName";
ResourceIdentifier reportResourceId = ReportResource.CreateResourceIdentifier(reportName);
ReportResource reportResource = client.GetReportResource(reportResourceId);

// get the collection of this SnapshotResource
SnapshotResourceCollection collection = reportResource.GetSnapshotResources();

// invoke the operation
string snapshotName = "testSnapshot";
bool result = await collection.ExistsAsync(snapshotName);

Console.WriteLine($"Succeeded: {result}");
