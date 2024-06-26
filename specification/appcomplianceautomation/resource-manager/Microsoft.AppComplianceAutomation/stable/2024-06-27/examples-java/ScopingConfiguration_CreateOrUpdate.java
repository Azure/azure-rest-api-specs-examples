
/**
 * Samples for ScopingConfiguration CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/appcomplianceautomation/resource-manager/Microsoft.AppComplianceAutomation/stable/2024-06-27/
     * examples/ScopingConfiguration_CreateOrUpdate.json
     */
    /**
     * Sample code: ScopingConfiguration_CreateOrUpdate.
     * 
     * @param manager Entry point to AppComplianceAutomationManager.
     */
    public static void scopingConfigurationCreateOrUpdate(
        com.azure.resourcemanager.appcomplianceautomation.AppComplianceAutomationManager manager) {
        manager.scopingConfigurations().createOrUpdateWithResponse("testReportName", "default", null,
            com.azure.core.util.Context.NONE);
    }
}
