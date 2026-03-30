
/**
 * Samples for AppServiceEnvironments ListWorkerPoolInstanceMetricDefinitions.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01/AppServiceEnvironments_ListWorkerPoolInstanceMetricDefinitions.json
     */
    /**
     * Sample code: Get metric definitions for a specific instance of a worker pool of an App Service Environment.
     * 
     * @param manager Entry point to AppServiceManager.
     */
    public static void getMetricDefinitionsForASpecificInstanceOfAWorkerPoolOfAnAppServiceEnvironment(
        com.azure.resourcemanager.appservice.AppServiceManager manager) {
        manager.serviceClient().getAppServiceEnvironments().listWorkerPoolInstanceMetricDefinitions("test-rg",
            "test-ase", "0", "10.8.0.7", com.azure.core.util.Context.NONE);
    }
}
