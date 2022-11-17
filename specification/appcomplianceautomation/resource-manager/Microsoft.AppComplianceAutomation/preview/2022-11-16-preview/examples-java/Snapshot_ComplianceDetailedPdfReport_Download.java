import com.azure.core.util.Context;
import com.azure.resourcemanager.appcomplianceautomation.models.DownloadType;
import com.azure.resourcemanager.appcomplianceautomation.models.SnapshotDownloadRequest;

/** Samples for SnapshotOperation Download. */
public final class Main {
    /*
     * x-ms-original-file: specification/appcomplianceautomation/resource-manager/Microsoft.AppComplianceAutomation/preview/2022-11-16-preview/examples/Snapshot_ComplianceDetailedPdfReport_Download.json
     */
    /**
     * Sample code: Snapshot_Download_ComplianceDetailedPdfReport.
     *
     * @param manager Entry point to AppComplianceAutomationManager.
     */
    public static void snapshotDownloadComplianceDetailedPdfReport(
        com.azure.resourcemanager.appcomplianceautomation.AppComplianceAutomationManager manager) {
        manager
            .snapshotOperations()
            .download(
                "testReportName",
                "testSnapshotName",
                new SnapshotDownloadRequest()
                    .withReportCreatorTenantId("00000000-0000-0000-0000-000000000000")
                    .withDownloadType(DownloadType.COMPLIANCE_DETAILED_PDF_REPORT)
                    .withOfferGuid("00000000-0000-0000-0000-000000000000"),
                Context.NONE);
    }
}
