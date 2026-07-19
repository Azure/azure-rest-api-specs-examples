
/**
 * Samples for ManagedDatabaseSensitivityLabels ListRecommendedByDatabase.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/ManagedDatabaseSensitivityLabelsListByDatabaseRecommended.json
     */
    /**
     * Sample code: Gets the recommended sensitivity labels of a given database in a managed database.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void getsTheRecommendedSensitivityLabelsOfAGivenDatabaseInAManagedDatabase(
        com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getManagedDatabaseSensitivityLabels().listRecommendedByDatabase("myRG",
            "myManagedInstanceName", "myDatabase", null, null, null, com.azure.core.util.Context.NONE);
    }
}
