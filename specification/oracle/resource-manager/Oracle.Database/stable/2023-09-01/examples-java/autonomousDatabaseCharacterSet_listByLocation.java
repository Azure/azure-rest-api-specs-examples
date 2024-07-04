
/**
 * Samples for AutonomousDatabaseCharacterSets ListByLocation.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/oracle/resource-manager/Oracle.Database/stable/2023-09-01/examples/
     * autonomousDatabaseCharacterSet_listByLocation.json
     */
    /**
     * Sample code: List autonomous db character sets by location.
     * 
     * @param manager Entry point to OracleDatabaseManager.
     */
    public static void listAutonomousDbCharacterSetsByLocation(
        com.azure.resourcemanager.oracledatabase.OracleDatabaseManager manager) {
        manager.autonomousDatabaseCharacterSets().listByLocation("eastus", com.azure.core.util.Context.NONE);
    }
}
