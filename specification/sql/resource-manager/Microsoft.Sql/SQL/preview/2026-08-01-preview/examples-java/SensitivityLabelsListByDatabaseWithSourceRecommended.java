
/**
 * Samples for SensitivityLabels ListRecommendedByDatabase.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-08-01-preview/SensitivityLabelsListByDatabaseWithSourceRecommended.json
     */
    /**
     * Sample code: Gets the recommended sensitivity labels of a given database.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void
        getsTheRecommendedSensitivityLabelsOfAGivenDatabase(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getSensitivityLabels().listRecommendedByDatabase("myRG", "myServer", "myDatabase", null,
            null, null, com.azure.core.util.Context.NONE);
    }
}
