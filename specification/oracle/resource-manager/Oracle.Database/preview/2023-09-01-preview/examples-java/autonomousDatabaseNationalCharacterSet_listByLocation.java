
/**
 * Samples for AutonomousDatabaseNationalCharacterSets ListByLocation.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/oracle/resource-manager/Oracle.Database/preview/2023-09-01-preview/examples/
     * autonomousDatabaseNationalCharacterSet_listByLocation.json
     */
    /**
     * Sample code: List autonomous db national character sets by location.
     * 
     * @param manager Entry point to OracleDatabaseManager.
     */
    public static void listAutonomousDbNationalCharacterSetsByLocation(
        com.azure.resourcemanager.oracledatabase.OracleDatabaseManager manager) {
        manager.autonomousDatabaseNationalCharacterSets().listByLocation("eastus", com.azure.core.util.Context.NONE);
    }
}
