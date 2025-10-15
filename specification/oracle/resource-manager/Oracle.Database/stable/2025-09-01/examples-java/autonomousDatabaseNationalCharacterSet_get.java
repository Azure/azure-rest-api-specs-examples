
/**
 * Samples for AutonomousDatabaseNationalCharacterSets Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01/autonomousDatabaseNationalCharacterSet_get.json
     */
    /**
     * Sample code: AutonomousDatabaseNationalCharacterSets_get.
     * 
     * @param manager Entry point to OracleDatabaseManager.
     */
    public static void autonomousDatabaseNationalCharacterSetsGet(
        com.azure.resourcemanager.oracledatabase.OracleDatabaseManager manager) {
        manager.autonomousDatabaseNationalCharacterSets().getWithResponse("eastus", "NATIONAL",
            com.azure.core.util.Context.NONE);
    }
}
