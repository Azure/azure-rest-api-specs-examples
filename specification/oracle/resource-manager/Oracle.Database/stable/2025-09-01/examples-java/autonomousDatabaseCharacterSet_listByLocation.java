
/**
 * Samples for AutonomousDatabaseCharacterSets ListByLocation.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01/autonomousDatabaseCharacterSet_listByLocation.json
     */
    /**
     * Sample code: AutonomousDatabaseCharacterSets_listByLocation.
     * 
     * @param manager Entry point to OracleDatabaseManager.
     */
    public static void autonomousDatabaseCharacterSetsListByLocation(
        com.azure.resourcemanager.oracledatabase.OracleDatabaseManager manager) {
        manager.autonomousDatabaseCharacterSets().listByLocation("eastus", com.azure.core.util.Context.NONE);
    }
}
