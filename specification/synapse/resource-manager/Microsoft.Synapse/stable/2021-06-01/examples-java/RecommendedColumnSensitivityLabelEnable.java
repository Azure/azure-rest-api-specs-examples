/** Samples for SqlPoolSensitivityLabels EnableRecommendation. */
public final class Main {
    /*
     * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/RecommendedColumnSensitivityLabelEnable.json
     */
    /**
     * Sample code: Enables sensitivity recommendations on a given column.
     *
     * @param manager Entry point to SynapseManager.
     */
    public static void enablesSensitivityRecommendationsOnAGivenColumn(
        com.azure.resourcemanager.synapse.SynapseManager manager) {
        manager
            .sqlPoolSensitivityLabels()
            .enableRecommendationWithResponse(
                "myRG", "myServer", "myDatabase", "dbo", "myTable", "myColumn", com.azure.core.util.Context.NONE);
    }
}
