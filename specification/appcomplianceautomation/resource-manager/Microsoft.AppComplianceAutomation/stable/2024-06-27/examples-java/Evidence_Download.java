
/**
 * Samples for Evidence Download.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-06-27/Evidence_Download.json
     */
    /**
     * Sample code: Evidence_Download.
     * 
     * @param manager Entry point to AppComplianceAutomationManager.
     */
    public static void
        evidenceDownload(com.azure.resourcemanager.appcomplianceautomation.AppComplianceAutomationManager manager) {
        manager.evidences().downloadWithResponse("testReportName", "evidence1", null, com.azure.core.util.Context.NONE);
    }
}
