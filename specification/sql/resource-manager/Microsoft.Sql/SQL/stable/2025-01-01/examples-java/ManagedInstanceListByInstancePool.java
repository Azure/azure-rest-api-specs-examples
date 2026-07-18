
/**
 * Samples for ManagedInstances ListByInstancePool.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/ManagedInstanceListByInstancePool.json
     */
    /**
     * Sample code: List managed instances by instance pool.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void listManagedInstancesByInstancePool(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getManagedInstances().listByInstancePool("Test1", "pool1", null,
            com.azure.core.util.Context.NONE);
    }
}
