/** Samples for SqlPools Pause. */
public final class Main {
    /*
     * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/PauseSqlPool.json
     */
    /**
     * Sample code: Pause a SQL Analytics pool.
     *
     * @param manager Entry point to SynapseManager.
     */
    public static void pauseASQLAnalyticsPool(com.azure.resourcemanager.synapse.SynapseManager manager) {
        manager.sqlPools().pause("Default-SQL-SouthEastAsia", "testsvr", "testdwdb", com.azure.core.util.Context.NONE);
    }
}
