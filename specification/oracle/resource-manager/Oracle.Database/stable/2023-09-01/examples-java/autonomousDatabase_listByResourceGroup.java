
/**
 * Samples for AutonomousDatabases ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/oracle/resource-manager/Oracle.Database/stable/2023-09-01/examples/
     * autonomousDatabase_listByResourceGroup.json
     */
    /**
     * Sample code: List Autonomous Database by resource group.
     * 
     * @param manager Entry point to OracleDatabaseManager.
     */
    public static void
        listAutonomousDatabaseByResourceGroup(com.azure.resourcemanager.oracledatabase.OracleDatabaseManager manager) {
        manager.autonomousDatabases().listByResourceGroup("rg000", com.azure.core.util.Context.NONE);
    }
}
