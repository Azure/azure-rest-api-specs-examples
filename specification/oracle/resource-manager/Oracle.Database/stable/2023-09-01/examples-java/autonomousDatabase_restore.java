
import com.azure.resourcemanager.oracledatabase.models.RestoreAutonomousDatabaseDetails;
import java.time.OffsetDateTime;

/**
 * Samples for AutonomousDatabases Restore.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/oracle/resource-manager/Oracle.Database/stable/2023-09-01/examples/autonomousDatabase_restore.json
     */
    /**
     * Sample code: Perform restore action on Autonomous Database.
     * 
     * @param manager Entry point to OracleDatabaseManager.
     */
    public static void performRestoreActionOnAutonomousDatabase(
        com.azure.resourcemanager.oracledatabase.OracleDatabaseManager manager) {
        manager.autonomousDatabases().restore("rg000", "databasedb1",
            new RestoreAutonomousDatabaseDetails().withTimestamp(OffsetDateTime.parse("2024-04-23T00:00:00.000Z")),
            com.azure.core.util.Context.NONE);
    }
}
