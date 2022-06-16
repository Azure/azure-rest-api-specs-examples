import com.azure.core.util.Context;

/** Samples for SqlPoolSensitivityLabels ListRecommended. */
public final class Main {
    /*
     * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/ListSqlPoolSensitivityLabelsWithSourceRecommended.json
     */
    /**
     * Sample code: Gets the recommended sensitivity labels of a given SQL Analytics pool.
     *
     * @param manager Entry point to SynapseManager.
     */
    public static void getsTheRecommendedSensitivityLabelsOfAGivenSQLAnalyticsPool(
        com.azure.resourcemanager.synapse.SynapseManager manager) {
        manager
            .sqlPoolSensitivityLabels()
            .listRecommended("myRG", "myServer", "myDatabase", null, null, null, Context.NONE);
    }
}
