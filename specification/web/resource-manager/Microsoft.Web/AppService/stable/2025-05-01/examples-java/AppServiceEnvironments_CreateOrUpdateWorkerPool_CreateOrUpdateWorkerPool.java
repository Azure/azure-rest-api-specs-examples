
import com.azure.resourcemanager.appservice.fluent.models.WorkerPoolResourceInner;

/**
 * Samples for AppServiceEnvironments CreateOrUpdateWorkerPool.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01/AppServiceEnvironments_CreateOrUpdateWorkerPool_CreateOrUpdateWorkerPool.json
     */
    /**
     * Sample code: Get properties of a worker pool.
     * 
     * @param manager Entry point to AppServiceManager.
     */
    public static void getPropertiesOfAWorkerPool(com.azure.resourcemanager.appservice.AppServiceManager manager) {
        manager.serviceClient().getAppServiceEnvironments().createOrUpdateWorkerPool("test-rg", "test-ase", "0",
            new WorkerPoolResourceInner().withWorkerSize("Small").withWorkerCount(3), com.azure.core.util.Context.NONE);
    }
}
