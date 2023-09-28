import com.azure.resourcemanager.appservice.fluent.models.WorkerPoolResourceInner;

/** Samples for AppServiceEnvironments CreateOrUpdateWorkerPool. */
public final class Main {
    /*
     * x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2022-09-01/examples/AppServiceEnvironments_CreateOrUpdateWorkerPool.json
     */
    /**
     * Sample code: Get properties of a worker pool.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getPropertiesOfAWorkerPool(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .webApps()
            .manager()
            .serviceClient()
            .getAppServiceEnvironments()
            .createOrUpdateWorkerPool(
                "test-rg",
                "test-ase",
                "0",
                new WorkerPoolResourceInner().withWorkerSize("Small").withWorkerCount(3),
                com.azure.core.util.Context.NONE);
    }
}
