
/**
 * Samples for AppServiceEnvironments ListMultiRoleUsages.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01/AppServiceEnvironments_ListMultiRoleUsages.json
     */
    /**
     * Sample code: Get usage metrics for a multi-role pool of an App Service Environment.
     * 
     * @param manager Entry point to AppServiceManager.
     */
    public static void getUsageMetricsForAMultiRolePoolOfAnAppServiceEnvironment(
        com.azure.resourcemanager.appservice.AppServiceManager manager) {
        manager.serviceClient().getAppServiceEnvironments().listMultiRoleUsages("test-rg", "test-ase",
            com.azure.core.util.Context.NONE);
    }
}
