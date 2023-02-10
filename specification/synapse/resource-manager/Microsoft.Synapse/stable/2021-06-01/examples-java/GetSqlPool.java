/** Samples for SqlPools Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/GetSqlPool.json
     */
    /**
     * Sample code: Get a SQL Analytics pool.
     *
     * @param manager Entry point to SynapseManager.
     */
    public static void getASQLAnalyticsPool(com.azure.resourcemanager.synapse.SynapseManager manager) {
        manager
            .sqlPools()
            .getWithResponse(
                "sqlcrudtest-6852", "sqlcrudtest-2080", "sqlcrudtest-9187", com.azure.core.util.Context.NONE);
    }
}
