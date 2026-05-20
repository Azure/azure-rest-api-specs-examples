
/**
 * Samples for ScopingConfiguration Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-06-27/ScopingConfiguration_Get.json
     */
    /**
     * Sample code: ScopingConfiguration.
     * 
     * @param manager Entry point to AppComplianceAutomationManager.
     */
    public static void
        scopingConfiguration(com.azure.resourcemanager.appcomplianceautomation.AppComplianceAutomationManager manager) {
        manager.scopingConfigurations().getWithResponse("testReportName", "default", com.azure.core.util.Context.NONE);
    }
}
