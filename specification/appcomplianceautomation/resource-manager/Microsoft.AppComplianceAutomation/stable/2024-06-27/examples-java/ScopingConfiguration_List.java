
/**
 * Samples for ScopingConfiguration List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-06-27/ScopingConfiguration_List.json
     */
    /**
     * Sample code: ScopingConfiguration_List.
     * 
     * @param manager Entry point to AppComplianceAutomationManager.
     */
    public static void scopingConfigurationList(
        com.azure.resourcemanager.appcomplianceautomation.AppComplianceAutomationManager manager) {
        manager.scopingConfigurations().list("testReportName", com.azure.core.util.Context.NONE);
    }
}
