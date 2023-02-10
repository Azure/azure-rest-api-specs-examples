/** Samples for SqlPoolSchemas List. */
public final class Main {
    /*
     * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/ListSqlPoolSchema.json
     */
    /**
     * Sample code: List the schema in a SQL Analytics pool.
     *
     * @param manager Entry point to SynapseManager.
     */
    public static void listTheSchemaInASQLAnalyticsPool(com.azure.resourcemanager.synapse.SynapseManager manager) {
        manager.sqlPoolSchemas().list("myRG", "serverName", "myDatabase", null, com.azure.core.util.Context.NONE);
    }
}
