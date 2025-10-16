
/**
 * Samples for AutonomousDatabaseCharacterSets Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01/autonomousDatabaseCharacterSet_get.json
     */
    /**
     * Sample code: AutonomousDatabaseCharacterSets_get.
     * 
     * @param manager Entry point to OracleDatabaseManager.
     */
    public static void
        autonomousDatabaseCharacterSetsGet(com.azure.resourcemanager.oracledatabase.OracleDatabaseManager manager) {
        manager.autonomousDatabaseCharacterSets().getWithResponse("eastus", "DATABASE",
            com.azure.core.util.Context.NONE);
    }
}
