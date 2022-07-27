import com.azure.core.util.Context;

/** Samples for SensitivityLabels DisableRecommendation. */
public final class Main {
    /*
     * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2017-03-01-preview/examples/RecommendedColumnSensitivityLabelDisable.json
     */
    /**
     * Sample code: Disables sensitivity recommendations on a given column.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void disablesSensitivityRecommendationsOnAGivenColumn(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .sqlServers()
            .manager()
            .serviceClient()
            .getSensitivityLabels()
            .disableRecommendationWithResponse(
                "myRG", "myServer", "myDatabase", "dbo", "myTable", "myColumn", Context.NONE);
    }
}
