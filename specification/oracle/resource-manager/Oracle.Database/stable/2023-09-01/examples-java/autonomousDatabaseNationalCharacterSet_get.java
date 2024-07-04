
/**
 * Samples for AutonomousDatabaseNationalCharacterSets Get.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/oracle/resource-manager/Oracle.Database/stable/2023-09-01/examples/
     * autonomousDatabaseNationalCharacterSet_get.json
     */
    /**
     * Sample code: Get autonomous db national character set.
     * 
     * @param manager Entry point to OracleDatabaseManager.
     */
    public static void
        getAutonomousDbNationalCharacterSet(com.azure.resourcemanager.oracledatabase.OracleDatabaseManager manager) {
        manager.autonomousDatabaseNationalCharacterSets().getWithResponse("eastus", "NATIONAL",
            com.azure.core.util.Context.NONE);
    }
}
