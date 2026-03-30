
/**
 * Samples for AppServiceEnvironments GetWorkerPool.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01/AppServiceEnvironments_GetWorkerPool.json
     */
    /**
     * Sample code: Get properties of a worker pool.
     * 
     * @param manager Entry point to AppServiceManager.
     */
    public static void getPropertiesOfAWorkerPool(com.azure.resourcemanager.appservice.AppServiceManager manager) {
        manager.serviceClient().getAppServiceEnvironments().getWorkerPoolWithResponse("test-rg", "test-ase",
            "workerPool1", com.azure.core.util.Context.NONE);
    }
}
