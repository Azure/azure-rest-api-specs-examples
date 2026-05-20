
/**
 * Samples for Report CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-06-27/Report_CreateOrUpdate.json
     */
    /**
     * Sample code: Report_CreateOrUpdate.
     * 
     * @param manager Entry point to AppComplianceAutomationManager.
     */
    public static void
        reportCreateOrUpdate(com.azure.resourcemanager.appcomplianceautomation.AppComplianceAutomationManager manager) {
        manager.reports().createOrUpdate("testReportName", null, com.azure.core.util.Context.NONE);
    }
}
