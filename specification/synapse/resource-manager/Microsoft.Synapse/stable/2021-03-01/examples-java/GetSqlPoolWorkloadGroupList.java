
/**
 * Samples for SqlPoolWorkloadGroup List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-03-01/examples/GetSqlPoolWorkloadGroupList.
     * json
     */
    /**
     * Sample code: Get the list of workload groups of a SQL Analytics pool.
     * 
     * @param manager Entry point to SynapseManager.
     */
    public static void
        getTheListOfWorkloadGroupsOfASQLAnalyticsPool(com.azure.resourcemanager.synapse.SynapseManager manager) {
        manager.sqlPoolWorkloadGroups().list("sqlcrudtest-6852", "sqlcrudtest-2080", "sqlcrudtest-9187",
            com.azure.core.util.Context.NONE);
    }
}
