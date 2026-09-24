
/**
 * Samples for InstancePools Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-08-01-preview/DeleteInstancePool.json
     */
    /**
     * Sample code: Delete an instance pool.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void deleteAnInstancePool(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getInstancePools().delete("group1", "testIP", com.azure.core.util.Context.NONE);
    }
}
