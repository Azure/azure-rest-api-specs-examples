
/**
 * Samples for AutonomousDatabaseBackups Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01/autonomousDatabaseBackup_get.json
     */
    /**
     * Sample code: AutonomousDatabaseBackups_Get.
     * 
     * @param manager Entry point to OracleDatabaseManager.
     */
    public static void
        autonomousDatabaseBackupsGet(com.azure.resourcemanager.oracledatabase.OracleDatabaseManager manager) {
        manager.autonomousDatabaseBackups().getWithResponse("rg000", "databasedb1", "1711644130",
            com.azure.core.util.Context.NONE);
    }
}
