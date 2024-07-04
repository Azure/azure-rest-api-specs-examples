
/**
 * Samples for AutonomousDatabaseBackups Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/oracle/resource-manager/Oracle.Database/stable/2023-09-01/examples/autonomousDatabaseBackup_get.
     * json
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
