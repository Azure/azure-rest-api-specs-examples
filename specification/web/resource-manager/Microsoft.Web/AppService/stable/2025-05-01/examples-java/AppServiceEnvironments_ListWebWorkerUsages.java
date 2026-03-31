
/**
 * Samples for AppServiceEnvironments ListWebWorkerUsages.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01/AppServiceEnvironments_ListWebWorkerUsages.json
     */
    /**
     * Sample code: Get usage metrics for a worker pool of an App Service Environment.
     * 
     * @param manager Entry point to AppServiceManager.
     */
    public static void getUsageMetricsForAWorkerPoolOfAnAppServiceEnvironment(
        com.azure.resourcemanager.appservice.AppServiceManager manager) {
        manager.serviceClient().getAppServiceEnvironments().listWebWorkerUsages("test-rg", "test-ase", "0",
            com.azure.core.util.Context.NONE);
    }
}
