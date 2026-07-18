
/**
 * Samples for DatabaseAutomaticTuning Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/DatabaseAutomaticTuningGet.json
     */
    /**
     * Sample code: Get a database's automatic tuning settings.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void getADatabaseSAutomaticTuningSettings(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getDatabaseAutomaticTunings().getWithResponse("default-sql-onebox", "testsvr11", "db1",
            com.azure.core.util.Context.NONE);
    }
}
