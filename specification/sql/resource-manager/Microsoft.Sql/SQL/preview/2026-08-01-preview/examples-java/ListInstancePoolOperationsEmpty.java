
/**
 * Samples for InstancePoolOperations ListByInstancePool.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-08-01-preview/ListInstancePoolOperationsEmpty.json
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
