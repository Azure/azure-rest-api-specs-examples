
/**
 * Samples for ScopingConfiguration Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/appcomplianceautomation/resource-manager/Microsoft.AppComplianceAutomation/stable/2024-06-27/
     * examples/ScopingConfiguration_Delete.json
     */
    /**
     * Sample code: ScopingConfiguration_Delete.
     * 
     * @param manager Entry point to AppComplianceAutomationManager.
     */
    public static void scopingConfigurationDelete(
        com.azure.resourcemanager.appcomplianceautomation.AppComplianceAutomationManager manager) {
        manager.scopingConfigurations().deleteByResourceGroupWithResponse("testReportName", "default",
            com.azure.core.util.Context.NONE);
    }
}
