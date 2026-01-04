
/**
 * Samples for AppServiceEnvironments ListWorkerPoolSkus.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/web/resource-manager/Microsoft.Web/AppService/stable/2025-03-01/examples/
     * AppServiceEnvironments_ListWorkerPoolSkus.json
     */
    /**
     * Sample code: Get available SKUs for scaling a worker pool.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getAvailableSKUsForScalingAWorkerPool(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.webApps().manager().serviceClient().getAppServiceEnvironments().listWorkerPoolSkus("test-rg", "test-ase",
            "workerPool1", com.azure.core.util.Context.NONE);
    }
}
