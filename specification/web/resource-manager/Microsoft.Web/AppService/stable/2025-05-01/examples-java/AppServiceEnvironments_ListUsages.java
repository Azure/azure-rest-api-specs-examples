
/**
 * Samples for AppServiceEnvironments ListUsages.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01/AppServiceEnvironments_ListUsages.json
     */
    /**
     * Sample code: Get global usage metrics of an App Service Environment.
     * 
     * @param manager Entry point to AppServiceManager.
     */
    public static void
        getGlobalUsageMetricsOfAnAppServiceEnvironment(com.azure.resourcemanager.appservice.AppServiceManager manager) {
        manager.serviceClient().getAppServiceEnvironments().listUsages("test-rg", "test-ase", null,
            com.azure.core.util.Context.NONE);
    }
}
