
import com.azure.resourcemanager.sql.models.RecommendedSensitivityLabelSource;

/**
 * Samples for ManagedDatabaseSensitivityLabels DisableRecommendation.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/ManagedDatabaseRecommendedColumnSensitivityLabelDisable.json
     */
    /**
     * Sample code: Disables the sensitivity recommendations on a given column.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void
        disablesTheSensitivityRecommendationsOnAGivenColumn(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getManagedDatabaseSensitivityLabels().disableRecommendationWithResponse("myRG",
            "myManagedInstanceName", "myDatabase", "dbo", "myTable", "myColumn",
            RecommendedSensitivityLabelSource.RECOMMENDED, com.azure.core.util.Context.NONE);
    }
}
