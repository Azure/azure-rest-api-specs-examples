/** Samples for SqlPools Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/DeleteSqlPool.json
     */
    /**
     * Sample code: Delete a SQL Analytics pool.
     *
     * @param manager Entry point to SynapseManager.
     */
    public static void deleteASQLAnalyticsPool(com.azure.resourcemanager.synapse.SynapseManager manager) {
        manager
            .sqlPools()
            .delete("ExampleResourceGroup", "ExampleWorkspace", "ExampleSqlPool", com.azure.core.util.Context.NONE);
    }
}
