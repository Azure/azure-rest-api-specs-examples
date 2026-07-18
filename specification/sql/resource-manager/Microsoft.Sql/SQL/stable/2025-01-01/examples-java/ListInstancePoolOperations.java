
/**
 * Samples for InstancePoolOperations ListByInstancePool.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/ListInstancePoolOperations.json
     */
    /**
     * Sample code: List the instance pool management operations with some results.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void
        listTheInstancePoolManagementOperationsWithSomeResults(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getInstancePoolOperations().listByInstancePool("resource-group", "test-instance-pool",
            com.azure.core.util.Context.NONE);
    }
}
