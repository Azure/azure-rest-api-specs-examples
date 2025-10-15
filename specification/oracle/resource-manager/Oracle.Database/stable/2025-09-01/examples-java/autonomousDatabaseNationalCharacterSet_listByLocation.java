
/**
 * Samples for AutonomousDatabaseNationalCharacterSets ListByLocation.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01/autonomousDatabaseNationalCharacterSet_listByLocation.json
     */
    /**
     * Sample code: AutonomousDatabaseNationalCharacterSets_listByLocation.
     * 
     * @param manager Entry point to OracleDatabaseManager.
     */
    public static void autonomousDatabaseNationalCharacterSetsListByLocation(
        com.azure.resourcemanager.oracledatabase.OracleDatabaseManager manager) {
        manager.autonomousDatabaseNationalCharacterSets().listByLocation("eastus", com.azure.core.util.Context.NONE);
    }
}
