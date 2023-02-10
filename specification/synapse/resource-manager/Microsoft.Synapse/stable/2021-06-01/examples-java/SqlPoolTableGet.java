/** Samples for SqlPoolTables Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/SqlPoolTableGet.json
     */
    /**
     * Sample code: Get database table.
     *
     * @param manager Entry point to SynapseManager.
     */
    public static void getDatabaseTable(com.azure.resourcemanager.synapse.SynapseManager manager) {
        manager
            .sqlPoolTables()
            .getWithResponse("myRG", "serverName", "myDatabase", "dbo", "table1", com.azure.core.util.Context.NONE);
    }
}
