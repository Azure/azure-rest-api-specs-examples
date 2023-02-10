/** Samples for SqlPoolSensitivityLabels ListCurrent. */
public final class Main {
    /*
     * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/ListSqlPoolsSensitivityLabelsWithSourceCurrent.json
     */
    /**
     * Sample code: Gets the current sensitivity labels of a given SQL Analytics pool.
     *
     * @param manager Entry point to SynapseManager.
     */
    public static void getsTheCurrentSensitivityLabelsOfAGivenSQLAnalyticsPool(
        com.azure.resourcemanager.synapse.SynapseManager manager) {
        manager
            .sqlPoolSensitivityLabels()
            .listCurrent("myRG", "myServer", "myDatabase", null, com.azure.core.util.Context.NONE);
    }
}
