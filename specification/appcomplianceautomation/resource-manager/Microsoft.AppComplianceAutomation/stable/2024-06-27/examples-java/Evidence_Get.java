
/**
 * Samples for Evidence Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-06-27/Evidence_Get.json
     */
    /**
     * Sample code: Evidence_Get.
     * 
     * @param manager Entry point to AppComplianceAutomationManager.
     */
    public static void
        evidenceGet(com.azure.resourcemanager.appcomplianceautomation.AppComplianceAutomationManager manager) {
        manager.evidences().getWithResponse("testReportName", "evidence1", com.azure.core.util.Context.NONE);
    }
}
