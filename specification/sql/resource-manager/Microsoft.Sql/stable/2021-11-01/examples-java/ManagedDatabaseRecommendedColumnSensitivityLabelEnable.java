
import com.azure.core.util.Context;

/** Samples for ManagedDatabaseSensitivityLabels EnableRecommendation. */
public final class Main {
    /*
     * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/
     * ManagedDatabaseRecommendedColumnSensitivityLabelEnable.json
     */
    /**
     * Sample code: Enables the sensitivity recommendations on a given column.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        enablesTheSensitivityRecommendationsOnAGivenColumn(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getManagedDatabaseSensitivityLabels()
            .enableRecommendationWithResponse("myRG", "myManagedInstanceName", "myDatabase", "dbo", "myTable",
                "myColumn", Context.NONE);
    }
}
