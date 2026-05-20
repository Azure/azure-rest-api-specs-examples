
/**
 * Samples for Report SyncCertRecord.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-06-27/Report_SyncCertRecord.json
     */
    /**
     * Sample code: Report_SyncCertRecord.
     * 
     * @param manager Entry point to AppComplianceAutomationManager.
     */
    public static void
        reportSyncCertRecord(com.azure.resourcemanager.appcomplianceautomation.AppComplianceAutomationManager manager) {
        manager.reports().syncCertRecord("testReportName", null, com.azure.core.util.Context.NONE);
    }
}
