
/**
 * Samples for AppServiceEnvironments ListWorkerPoolInstanceMetricDefinitions.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/web/resource-manager/Microsoft.Web/AppService/stable/2025-03-01/examples/
     * AppServiceEnvironments_ListWorkerPoolInstanceMetricDefinitions.json
     */
    /**
     * Sample code: Get metric definitions for a specific instance of a worker pool of an App Service Environment.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getMetricDefinitionsForASpecificInstanceOfAWorkerPoolOfAnAppServiceEnvironment(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure.webApps().manager().serviceClient().getAppServiceEnvironments().listWorkerPoolInstanceMetricDefinitions(
            "test-rg", "test-ase", "0", "10.8.0.7", com.azure.core.util.Context.NONE);
    }
}
