import com.azure.core.util.Context;
import com.azure.resourcemanager.appcomplianceautomation.models.DownloadType;
import com.azure.resourcemanager.appcomplianceautomation.models.SnapshotDownloadRequest;

/** Samples for SnapshotOperation Download. */
public final class Main {
    /*
     * x-ms-original-file: specification/appcomplianceautomation/resource-manager/Microsoft.AppComplianceAutomation/preview/2022-11-16-preview/examples/Snapshot_ComplianceReport_Download.json
     */
    /**
     * Sample code: Snapshot_Download_ComplianceReport.
     *
     * @param manager Entry point to AppComplianceAutomationManager.
     */
    public static void snapshotDownloadComplianceReport(
        com.azure.resourcemanager.appcomplianceautomation.AppComplianceAutomationManager manager) {
        manager
            .snapshotOperations()
            .download(
                "testReportName",
                "testSnapshotName",
                new SnapshotDownloadRequest()
                    .withReportCreatorTenantId("00000000-0000-0000-0000-000000000000")
                    .withDownloadType(DownloadType.COMPLIANCE_REPORT)
                    .withOfferGuid("00000000-0000-0000-0000-000000000000"),
                Context.NONE);
    }
}
