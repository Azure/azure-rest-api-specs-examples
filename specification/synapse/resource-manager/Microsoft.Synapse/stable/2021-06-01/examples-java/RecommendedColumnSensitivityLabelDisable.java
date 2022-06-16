import com.azure.core.util.Context;

/** Samples for SqlPoolSensitivityLabels DisableRecommendation. */
public final class Main {
    /*
     * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/RecommendedColumnSensitivityLabelDisable.json
     */
    /**
     * Sample code: Disables sensitivity recommendations on a given column.
     *
     * @param manager Entry point to SynapseManager.
     */
    public static void disablesSensitivityRecommendationsOnAGivenColumn(
        com.azure.resourcemanager.synapse.SynapseManager manager) {
        manager
            .sqlPoolSensitivityLabels()
            .disableRecommendationWithResponse(
                "myRG", "myServer", "myDatabase", "dbo", "myTable", "myColumn", Context.NONE);
    }
}
