
/**
 * Samples for Report Fix.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-06-27/Report_Fix.json
     */
    /**
     * Sample code: Report_Fix.
     * 
     * @param manager Entry point to AppComplianceAutomationManager.
     */
    public static void
        reportFix(com.azure.resourcemanager.appcomplianceautomation.AppComplianceAutomationManager manager) {
        manager.reports().fix("testReport", com.azure.core.util.Context.NONE);
    }
}
