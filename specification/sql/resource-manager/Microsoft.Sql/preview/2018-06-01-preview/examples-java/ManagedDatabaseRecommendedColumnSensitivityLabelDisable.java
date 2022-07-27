import com.azure.core.util.Context;

/** Samples for ManagedDatabaseSensitivityLabels DisableRecommendation. */
public final class Main {
    /*
     * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2018-06-01-preview/examples/ManagedDatabaseRecommendedColumnSensitivityLabelDisable.json
     */
    /**
     * Sample code: Disables the sensitivity recommendations on a given column.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void disablesTheSensitivityRecommendationsOnAGivenColumn(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .sqlServers()
            .manager()
            .serviceClient()
            .getManagedDatabaseSensitivityLabels()
            .disableRecommendationWithResponse(
                "myRG", "myManagedInstanceName", "myDatabase", "dbo", "myTable", "myColumn", Context.NONE);
    }
}
