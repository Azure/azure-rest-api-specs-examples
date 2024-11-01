
/**
 * Samples for SqlPoolColumns Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/SqlPoolColumnGet.json
     */
    /**
     * Sample code: Get database column.
     * 
     * @param manager Entry point to SynapseManager.
     */
    public static void getDatabaseColumn(com.azure.resourcemanager.synapse.SynapseManager manager) {
        manager.sqlPoolColumns().getWithResponse("myRG", "serverName", "myDatabase", "dbo", "table1", "column1",
            com.azure.core.util.Context.NONE);
    }
}
