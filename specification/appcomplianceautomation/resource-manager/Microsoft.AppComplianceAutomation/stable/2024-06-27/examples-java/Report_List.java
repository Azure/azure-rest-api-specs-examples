
/**
 * Samples for Report List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-06-27/Report_List.json
     */
    /**
     * Sample code: Report_List.
     * 
     * @param manager Entry point to AppComplianceAutomationManager.
     */
    public static void
        reportList(com.azure.resourcemanager.appcomplianceautomation.AppComplianceAutomationManager manager) {
        manager.reports().list("1", 100, null, null, null, "00000000-0000-0000-0000-000000000000",
            "00000000-0000-0000-0000-000000000000", com.azure.core.util.Context.NONE);
    }
}
