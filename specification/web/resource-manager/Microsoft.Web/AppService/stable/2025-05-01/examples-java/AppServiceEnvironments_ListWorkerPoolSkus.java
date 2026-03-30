
/**
 * Samples for AppServiceEnvironments ListWorkerPoolSkus.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01/AppServiceEnvironments_ListWorkerPoolSkus.json
     */
    /**
     * Sample code: Get available SKUs for scaling a worker pool.
     * 
     * @param manager Entry point to AppServiceManager.
     */
    public static void
        getAvailableSKUsForScalingAWorkerPool(com.azure.resourcemanager.appservice.AppServiceManager manager) {
        manager.serviceClient().getAppServiceEnvironments().listWorkerPoolSkus("test-rg", "test-ase", "workerPool1",
            com.azure.core.util.Context.NONE);
    }
}
