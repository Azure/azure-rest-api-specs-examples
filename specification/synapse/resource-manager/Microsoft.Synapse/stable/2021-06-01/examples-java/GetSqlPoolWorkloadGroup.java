/** Samples for SqlPoolWorkloadGroup Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/GetSqlPoolWorkloadGroup.json
     */
    /**
     * Sample code: Get a a workload group of a SQL Analytics pool.
     *
     * @param manager Entry point to SynapseManager.
     */
    public static void getAAWorkloadGroupOfASQLAnalyticsPool(com.azure.resourcemanager.synapse.SynapseManager manager) {
        manager
            .sqlPoolWorkloadGroups()
            .getWithResponse(
                "sqlcrudtest-6852",
                "sqlcrudtest-2080",
                "sqlcrudtest-9187",
                "smallrc",
                com.azure.core.util.Context.NONE);
    }
}
