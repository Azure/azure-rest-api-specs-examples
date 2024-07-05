
import com.azure.resourcemanager.oracledatabase.models.AutonomousDatabaseBackupProperties;

/**
 * Samples for AutonomousDatabaseBackups CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/oracle/resource-manager/Oracle.Database/stable/2023-09-01/examples/autonomousDatabaseBackup_create.
     * json
     */
    /**
     * Sample code: AutonomousDatabaseBackups_CreateOrUpdate.
     * 
     * @param manager Entry point to OracleDatabaseManager.
     */
    public static void autonomousDatabaseBackupsCreateOrUpdate(
        com.azure.resourcemanager.oracledatabase.OracleDatabaseManager manager) {
        manager.autonomousDatabaseBackups().define("1711644130").withExistingAutonomousDatabase("rg000", "databasedb1")
            .withProperties(new AutonomousDatabaseBackupProperties().withDisplayName("Nightly Backup")
                .withRetentionPeriodInDays(365))
            .create();
    }
}
