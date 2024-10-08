
import com.azure.resourcemanager.oracledatabase.models.AutonomousDatabase;

/**
 * Samples for AutonomousDatabases Update.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/oracle/resource-manager/Oracle.Database/stable/2023-09-01/examples/autonomousDatabase_patch.json
     */
    /**
     * Sample code: Patch Autonomous Database.
     * 
     * @param manager Entry point to OracleDatabaseManager.
     */
    public static void patchAutonomousDatabase(com.azure.resourcemanager.oracledatabase.OracleDatabaseManager manager) {
        AutonomousDatabase resource = manager.autonomousDatabases()
            .getByResourceGroupWithResponse("rg000", "databasedb1", com.azure.core.util.Context.NONE).getValue();
        resource.update().apply();
    }
}
