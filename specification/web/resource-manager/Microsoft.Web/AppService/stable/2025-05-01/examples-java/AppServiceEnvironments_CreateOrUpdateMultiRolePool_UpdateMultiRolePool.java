
import com.azure.resourcemanager.appservice.fluent.models.WorkerPoolResourceInner;

/**
 * Samples for AppServiceEnvironments UpdateMultiRolePool.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01/AppServiceEnvironments_CreateOrUpdateMultiRolePool_UpdateMultiRolePool.json
     */
    /**
     * Sample code: Create or update a multi-role pool.
     * 
     * @param manager Entry point to AppServiceManager.
     */
    public static void createOrUpdateAMultiRolePool(com.azure.resourcemanager.appservice.AppServiceManager manager) {
        manager.serviceClient().getAppServiceEnvironments().updateMultiRolePoolWithResponse("test-rg", "test-ase",
            new WorkerPoolResourceInner().withWorkerSize("Medium").withWorkerCount(3),
            com.azure.core.util.Context.NONE);
    }
}
