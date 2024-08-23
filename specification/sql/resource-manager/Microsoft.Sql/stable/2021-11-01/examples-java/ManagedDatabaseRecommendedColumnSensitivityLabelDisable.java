
/**
 * Samples for ManagedDatabaseSensitivityLabels DisableRecommendation.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/
     * ManagedDatabaseRecommendedColumnSensitivityLabelDisable.json
     */
    /**
     * Sample code: Disables the sensitivity recommendations on a given column.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        disablesTheSensitivityRecommendationsOnAGivenColumn(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getManagedDatabaseSensitivityLabels()
            .disableRecommendationWithResponse("myRG", "myManagedInstanceName", "myDatabase", "dbo", "myTable",
                "myColumn", com.azure.core.util.Context.NONE);
    }
}
