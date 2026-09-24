
/**
 * Samples for SensitivityLabels ListByDatabase.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-08-01-preview/SensitivityLabelsListByDatabase.json
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
