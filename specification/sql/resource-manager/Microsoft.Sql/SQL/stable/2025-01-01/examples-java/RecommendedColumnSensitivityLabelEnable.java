
import com.azure.resourcemanager.sql.models.RecommendedSensitivityLabelSource;

/**
 * Samples for SensitivityLabels EnableRecommendation.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/RecommendedColumnSensitivityLabelEnable.json
     */
    /**
     * Sample code: Enables sensitivity recommendations on a given column.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void
        enablesSensitivityRecommendationsOnAGivenColumn(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getSensitivityLabels().enableRecommendationWithResponse("myRG", "myServer",
            "myDatabase", "dbo", "myTable", "myColumn", RecommendedSensitivityLabelSource.RECOMMENDED,
            com.azure.core.util.Context.NONE);
    }
}
