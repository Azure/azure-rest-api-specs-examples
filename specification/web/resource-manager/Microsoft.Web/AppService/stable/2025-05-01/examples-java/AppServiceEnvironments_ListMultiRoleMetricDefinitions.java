
/**
 * Samples for AppServiceEnvironments ListMultiRoleMetricDefinitions.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01/AppServiceEnvironments_ListMultiRoleMetricDefinitions.json
     */
    /**
     * Sample code: Get metric definitions for a multi-role pool of an App Service Environment.
     * 
     * @param manager Entry point to AppServiceManager.
     */
    public static void getMetricDefinitionsForAMultiRolePoolOfAnAppServiceEnvironment(
        com.azure.resourcemanager.appservice.AppServiceManager manager) {
        manager.serviceClient().getAppServiceEnvironments().listMultiRoleMetricDefinitions("test-rg", "test-ase",
            com.azure.core.util.Context.NONE);
    }
}
