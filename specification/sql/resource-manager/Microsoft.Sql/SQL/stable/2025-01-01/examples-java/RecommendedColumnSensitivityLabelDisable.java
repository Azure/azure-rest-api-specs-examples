
import com.azure.resourcemanager.sql.models.RecommendedSensitivityLabelSource;

/**
 * Samples for SensitivityLabels DisableRecommendation.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/RecommendedColumnSensitivityLabelDisable.json
     */
    /**
     * Sample code: Disables sensitivity recommendations on a given column.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void
        disablesSensitivityRecommendationsOnAGivenColumn(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getSensitivityLabels().disableRecommendationWithResponse("myRG", "myServer",
            "myDatabase", "dbo", "myTable", "myColumn", RecommendedSensitivityLabelSource.RECOMMENDED,
            com.azure.core.util.Context.NONE);
    }
}
