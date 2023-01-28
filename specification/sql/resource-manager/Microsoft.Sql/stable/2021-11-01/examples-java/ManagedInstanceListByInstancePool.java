import com.azure.core.util.Context;

/** Samples for ManagedInstances ListByInstancePool. */
public final class Main {
    /*
     * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/ManagedInstanceListByInstancePool.json
     */
    /**
     * Sample code: List managed instances by instance pool.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listManagedInstancesByInstancePool(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .sqlServers()
            .manager()
            .serviceClient()
            .getManagedInstances()
            .listByInstancePool("Test1", "pool1", null, Context.NONE);
    }
}
