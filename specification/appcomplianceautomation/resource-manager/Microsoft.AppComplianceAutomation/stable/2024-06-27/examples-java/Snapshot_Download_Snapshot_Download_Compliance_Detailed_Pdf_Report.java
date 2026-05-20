
/**
 * Samples for Snapshot Download.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-06-27/Snapshot_Download_Snapshot_Download_Compliance_Detailed_Pdf_Report.json
     */
    /**
     * Sample code: Snapshot_Download_ComplianceDetailedPdfReport.
     * 
     * @param manager Entry point to AppComplianceAutomationManager.
     */
    public static void snapshotDownloadComplianceDetailedPdfReport(
        com.azure.resourcemanager.appcomplianceautomation.AppComplianceAutomationManager manager) {
        manager.snapshots().download("testReportName", "testSnapshotName", null, com.azure.core.util.Context.NONE);
    }
}
