
/**
 * Samples for AutonomousDatabases ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01/autonomousDatabase_listByResourceGroup.json
     */
    /**
     * Sample code: AutonomousDatabases_listByResourceGroup.
     * 
     * @param manager Entry point to OracleDatabaseManager.
     */
    public static void
        autonomousDatabasesListByResourceGroup(com.azure.resourcemanager.oracledatabase.OracleDatabaseManager manager) {
        manager.autonomousDatabases().listByResourceGroup("rg000", com.azure.core.util.Context.NONE);
    }
}
