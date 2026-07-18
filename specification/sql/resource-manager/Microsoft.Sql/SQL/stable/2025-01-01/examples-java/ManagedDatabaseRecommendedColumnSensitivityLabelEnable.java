
import com.azure.resourcemanager.sql.models.RecommendedSensitivityLabelSource;

/**
 * Samples for ManagedDatabaseSensitivityLabels EnableRecommendation.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/ManagedDatabaseRecommendedColumnSensitivityLabelEnable.json
     */
    /**
     * Sample code: Enables the sensitivity recommendations on a given column.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void
        enablesTheSensitivityRecommendationsOnAGivenColumn(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getManagedDatabaseSensitivityLabels().enableRecommendationWithResponse("myRG",
            "myManagedInstanceName", "myDatabase", "dbo", "myTable", "myColumn",
            RecommendedSensitivityLabelSource.RECOMMENDED, com.azure.core.util.Context.NONE);
    }
}
