
/**
 * Samples for SensitivityLabels ListByDatabase.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/SensitivityLabelsListByDatabase.json
     */
    /**
     * Sample code: Gets the current and recommended sensitivity labels of a given database.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void getsTheCurrentAndRecommendedSensitivityLabelsOfAGivenDatabase(
        com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getSensitivityLabels().listByDatabase("myRG", "myServer", "myDatabase", null,
            com.azure.core.util.Context.NONE);
    }
}
