
/**
 * Samples for AppServiceEnvironments ListWebWorkerMetricDefinitions.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01/AppServiceEnvironments_ListWebWorkerMetricDefinitions.json
     */
    /**
     * Sample code: Get metric definitions for a worker pool of an App Service Environment.
     * 
     * @param manager Entry point to AppServiceManager.
     */
    public static void getMetricDefinitionsForAWorkerPoolOfAnAppServiceEnvironment(
        com.azure.resourcemanager.appservice.AppServiceManager manager) {
        manager.serviceClient().getAppServiceEnvironments().listWebWorkerMetricDefinitions("test-rg", "test-ase", "0",
            com.azure.core.util.Context.NONE);
    }
}
