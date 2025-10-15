
/**
 * Samples for AutonomousDatabaseVersions Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01/autonomousDatabaseVersion_get.json
     */
    /**
     * Sample code: AutonomousDatabaseVersions_get.
     * 
     * @param manager Entry point to OracleDatabaseManager.
     */
    public static void
        autonomousDatabaseVersionsGet(com.azure.resourcemanager.oracledatabase.OracleDatabaseManager manager) {
        manager.autonomousDatabaseVersions().getWithResponse("eastus", "18.4.0.0", com.azure.core.util.Context.NONE);
    }
}
