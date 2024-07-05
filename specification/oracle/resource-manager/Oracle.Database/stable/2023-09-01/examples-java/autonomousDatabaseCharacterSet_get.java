
/**
 * Samples for AutonomousDatabaseCharacterSets Get.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/oracle/resource-manager/Oracle.Database/stable/2023-09-01/examples/
     * autonomousDatabaseCharacterSet_get.json
     */
    /**
     * Sample code: Get autonomous db character set.
     * 
     * @param manager Entry point to OracleDatabaseManager.
     */
    public static void
        getAutonomousDbCharacterSet(com.azure.resourcemanager.oracledatabase.OracleDatabaseManager manager) {
        manager.autonomousDatabaseCharacterSets().getWithResponse("eastus", "DATABASE",
            com.azure.core.util.Context.NONE);
    }
}
