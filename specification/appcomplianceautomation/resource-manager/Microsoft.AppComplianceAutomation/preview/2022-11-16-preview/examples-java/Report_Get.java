import com.azure.core.util.Context;

/** Samples for ReportOperation Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/appcomplianceautomation/resource-manager/Microsoft.AppComplianceAutomation/preview/2022-11-16-preview/examples/Report_Get.json
     */
    /**
     * Sample code: Report_Get.
     *
     * @param manager Entry point to AppComplianceAutomationManager.
     */
    public static void reportGet(
        com.azure.resourcemanager.appcomplianceautomation.AppComplianceAutomationManager manager) {
        manager.reportOperations().getWithResponse("testReport", Context.NONE);
    }
}
