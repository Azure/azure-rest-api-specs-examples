
/**
 * Samples for AppServiceEnvironments GetWorkerPool.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/web/resource-manager/Microsoft.Web/AppService/stable/2025-03-01/examples/
     * AppServiceEnvironments_GetWorkerPool.json
     */
    /**
     * Sample code: Get properties of a worker pool.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getPropertiesOfAWorkerPool(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.webApps().manager().serviceClient().getAppServiceEnvironments().getWorkerPoolWithResponse("test-rg",
            "test-ase", "workerPool1", com.azure.core.util.Context.NONE);
    }
}
