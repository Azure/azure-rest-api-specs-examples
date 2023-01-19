using System;
using System.Threading.Tasks;
using Azure;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager;
using Azure.ResourceManager.AppComplianceAutomation;
using Azure.ResourceManager.AppComplianceAutomation.Models;

// Generated from example definition: specification/appcomplianceautomation/resource-manager/Microsoft.AppComplianceAutomation/preview/2022-11-16-preview/examples/Snapshot_ComplianceDetailedPdfReport_Download.json
// this example is just showing the usage of "Snapshot_Download" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SnapshotResource created on azure
// for more information of creating SnapshotResource, please refer to the document of SnapshotResource
string reportName = "testReportName";
string snapshotName = "testSnapshotName";
ResourceIdentifier snapshotResourceId = SnapshotResource.CreateResourceIdentifier(reportName, snapshotName);
SnapshotResource snapshotResource = client.GetSnapshotResource(snapshotResourceId);

// invoke the operation
SnapshotDownloadContent content = new SnapshotDownloadContent(DownloadType.ComplianceDetailedPdfReport)
{
    ReportCreatorTenantId = "00000000-0000-0000-0000-000000000000",
    OfferGuid = "00000000-0000-0000-0000-000000000000",
};
ArmOperation<DownloadResponse> lro = await snapshotResource.DownloadAsync(WaitUntil.Completed, content);
DownloadResponse result = lro.Value;

Console.WriteLine($"Succeeded: {result}");
