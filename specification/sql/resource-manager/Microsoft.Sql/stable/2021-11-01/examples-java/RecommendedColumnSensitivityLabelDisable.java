
/**
 * Samples for SensitivityLabels DisableRecommendation.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/
     * RecommendedColumnSensitivityLabelDisable.json
     */
    /**
     * Sample code: Disables sensitivity recommendations on a given column.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        disablesSensitivityRecommendationsOnAGivenColumn(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getSensitivityLabels().disableRecommendationWithResponse("myRG",
            "myServer", "myDatabase", "dbo", "myTable", "myColumn", com.azure.core.util.Context.NONE);
    }
}
