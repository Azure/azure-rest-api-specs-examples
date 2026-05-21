
/**
 * Samples for Report Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-06-27/Report_Delete.json
     */
    /**
     * Sample code: Report_Delete.
     * 
     * @param manager Entry point to AppComplianceAutomationManager.
     */
    public static void
        reportDelete(com.azure.resourcemanager.appcomplianceautomation.AppComplianceAutomationManager manager) {
        manager.reports().delete("testReportName", com.azure.core.util.Context.NONE);
    }
}
