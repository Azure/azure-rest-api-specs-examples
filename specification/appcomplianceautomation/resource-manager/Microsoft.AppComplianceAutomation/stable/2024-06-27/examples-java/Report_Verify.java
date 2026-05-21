
/**
 * Samples for Report Verify.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-06-27/Report_Verify.json
     */
    /**
     * Sample code: Report_Verify.
     * 
     * @param manager Entry point to AppComplianceAutomationManager.
     */
    public static void
        reportVerify(com.azure.resourcemanager.appcomplianceautomation.AppComplianceAutomationManager manager) {
        manager.reports().verify("testReport", com.azure.core.util.Context.NONE);
    }
}
