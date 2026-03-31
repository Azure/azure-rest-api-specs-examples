
/**
 * Samples for AppServiceEnvironments ListMultiRolePoolInstanceMetricDefinitions.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01/AppServiceEnvironments_ListMultiRolePoolInstanceMetricDefinitions.json
     */
    /**
     * Sample code: Get metric definitions for a specific instance of a multi-role pool of an App Service Environment.
     * 
     * @param manager Entry point to AppServiceManager.
     */
    public static void getMetricDefinitionsForASpecificInstanceOfAMultiRolePoolOfAnAppServiceEnvironment(
        com.azure.resourcemanager.appservice.AppServiceManager manager) {
        manager.serviceClient().getAppServiceEnvironments().listMultiRolePoolInstanceMetricDefinitions("test-rg",
            "test-ase", "10.7.1.8", com.azure.core.util.Context.NONE);
    }
}
