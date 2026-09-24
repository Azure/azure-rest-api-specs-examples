
/**
 * Samples for InstancePools GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-08-01-preview/GetInstancePool.json
     */
    /**
     * Sample code: Get an instance pool.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void getAnInstancePool(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getInstancePools().getByResourceGroupWithResponse("group1", "testIP",
            com.azure.core.util.Context.NONE);
    }
}
