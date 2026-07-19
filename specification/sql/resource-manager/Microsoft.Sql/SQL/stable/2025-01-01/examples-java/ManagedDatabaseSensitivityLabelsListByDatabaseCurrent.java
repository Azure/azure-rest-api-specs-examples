
/**
 * Samples for ManagedDatabaseSensitivityLabels ListCurrentByDatabase.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/ManagedDatabaseSensitivityLabelsListByDatabaseCurrent.json
     */
    /**
     * Sample code: Gets the current sensitivity labels of a given database in a managed database.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void getsTheCurrentSensitivityLabelsOfAGivenDatabaseInAManagedDatabase(
        com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getManagedDatabaseSensitivityLabels().listCurrentByDatabase("myRG",
            "myManagedInstanceName", "myDatabase", null, null, null, com.azure.core.util.Context.NONE);
    }
}
