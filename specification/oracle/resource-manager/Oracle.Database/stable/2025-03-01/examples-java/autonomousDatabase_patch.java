
import com.azure.resourcemanager.oracledatabase.models.AutonomousDatabase;

/**
 * Samples for AutonomousDatabases Update.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-03-01/autonomousDatabase_patch.json
     */
    /**
     * Sample code: AutonomousDatabases_update.
     * 
     * @param manager Entry point to OracleDatabaseManager.
     */
    public static void
        autonomousDatabasesUpdate(com.azure.resourcemanager.oracledatabase.OracleDatabaseManager manager) {
        AutonomousDatabase resource = manager.autonomousDatabases()
            .getByResourceGroupWithResponse("rg000", "databasedb1", com.azure.core.util.Context.NONE).getValue();
        resource.update().apply();
    }
}
