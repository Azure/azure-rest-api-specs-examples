
/**
 * Samples for AppServiceEnvironments ListWebWorkerMetricDefinitions.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2023-01-01/examples/
     * AppServiceEnvironments_ListWebWorkerMetricDefinitions.json
     */
    /**
     * Sample code: Get metric definitions for a worker pool of an App Service Environment.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getMetricDefinitionsForAWorkerPoolOfAnAppServiceEnvironment(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure.webApps().manager().serviceClient().getAppServiceEnvironments().listWebWorkerMetricDefinitions("test-rg",
            "test-ase", "0", com.azure.core.util.Context.NONE);
    }
}
