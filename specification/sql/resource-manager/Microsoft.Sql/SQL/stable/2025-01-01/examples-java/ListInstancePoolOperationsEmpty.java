
/**
 * Samples for InstancePoolOperations ListByInstancePool.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/ListInstancePoolOperationsEmpty.json
     */
    /**
     * Sample code: List the instance pool management operations with no results.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void
        listTheInstancePoolManagementOperationsWithNoResults(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getInstancePoolOperations().listByInstancePool("resource-group", "test-instance-pool",
            com.azure.core.util.Context.NONE);
    }
}
