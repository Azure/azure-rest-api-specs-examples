/** Samples for SqlPoolTables ListBySchema. */
public final class Main {
    /*
     * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/ListSqlPoolTables.json
     */
    /**
     * Sample code: List the tables of a given schema in a SQL Analytics pool.
     *
     * @param manager Entry point to SynapseManager.
     */
    public static void listTheTablesOfAGivenSchemaInASQLAnalyticsPool(
        com.azure.resourcemanager.synapse.SynapseManager manager) {
        manager
            .sqlPoolTables()
            .listBySchema("myRG", "serverName", "myDatabase", "dbo", null, com.azure.core.util.Context.NONE);
    }
}
