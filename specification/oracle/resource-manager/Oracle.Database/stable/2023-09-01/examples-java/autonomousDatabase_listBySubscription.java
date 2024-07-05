
/**
 * Samples for AutonomousDatabases List.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/oracle/resource-manager/Oracle.Database/stable/2023-09-01/examples/
     * autonomousDatabase_listBySubscription.json
     */
    /**
     * Sample code: List Autonomous Database by subscription.
     * 
     * @param manager Entry point to OracleDatabaseManager.
     */
    public static void
        listAutonomousDatabaseBySubscription(com.azure.resourcemanager.oracledatabase.OracleDatabaseManager manager) {
        manager.autonomousDatabases().list(com.azure.core.util.Context.NONE);
    }
}
