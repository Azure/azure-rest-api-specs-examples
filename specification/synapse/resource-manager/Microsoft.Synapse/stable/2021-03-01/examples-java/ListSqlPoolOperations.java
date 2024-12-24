
/**
 * Samples for SqlPoolOperations List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-03-01/examples/ListSqlPoolOperations.json
     */
    /**
     * Sample code: List the Sql Analytics pool management operations.
     * 
     * @param manager Entry point to SynapseManager.
     */
    public static void
        listTheSqlAnalyticsPoolManagementOperations(com.azure.resourcemanager.synapse.SynapseManager manager) {
        manager.sqlPoolOperations().list("sqlcrudtest-7398", "sqlcrudtest-4645", "testdb",
            com.azure.core.util.Context.NONE);
    }
}
