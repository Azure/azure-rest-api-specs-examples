
/**
 * Samples for Report Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-06-27/Report_Get.json
     */
    /**
     * Sample code: Report_Get.
     * 
     * @param manager Entry point to AppComplianceAutomationManager.
     */
    public static void
        reportGet(com.azure.resourcemanager.appcomplianceautomation.AppComplianceAutomationManager manager) {
        manager.reports().getWithResponse("testReport", com.azure.core.util.Context.NONE);
    }
}
