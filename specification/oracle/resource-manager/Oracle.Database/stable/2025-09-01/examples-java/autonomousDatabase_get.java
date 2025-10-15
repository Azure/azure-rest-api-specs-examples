
/**
 * Samples for AutonomousDatabases GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01/autonomousDatabase_get.json
     */
    /**
     * Sample code: AutonomousDatabases_Get.
     * 
     * @param manager Entry point to OracleDatabaseManager.
     */
    public static void autonomousDatabasesGet(com.azure.resourcemanager.oracledatabase.OracleDatabaseManager manager) {
        manager.autonomousDatabases().getByResourceGroupWithResponse("rg000", "databasedb1",
            com.azure.core.util.Context.NONE);
    }
}
