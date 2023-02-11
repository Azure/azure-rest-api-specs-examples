/** Samples for SqlPoolRestorePoints Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/SqlPoolRestorePointsGet.json
     */
    /**
     * Sample code: Gets a Sql pool restore point.
     *
     * @param manager Entry point to SynapseManager.
     */
    public static void getsASqlPoolRestorePoint(com.azure.resourcemanager.synapse.SynapseManager manager) {
        manager
            .sqlPoolRestorePoints()
            .getWithResponse(
                "Default-SQL-SouthEastAsia",
                "testws",
                "testpool",
                "131546477590000000",
                com.azure.core.util.Context.NONE);
    }
}
