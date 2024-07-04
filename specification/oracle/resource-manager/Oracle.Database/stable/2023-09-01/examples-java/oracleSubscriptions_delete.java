
/**
 * Samples for OracleSubscriptions Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/oracle/resource-manager/Oracle.Database/stable/2023-09-01/examples/oracleSubscriptions_delete.json
     */
    /**
     * Sample code: Delete Oracle Subscription.
     * 
     * @param manager Entry point to OracleDatabaseManager.
     */
    public static void
        deleteOracleSubscription(com.azure.resourcemanager.oracledatabase.OracleDatabaseManager manager) {
        manager.oracleSubscriptions().delete(com.azure.core.util.Context.NONE);
    }
}
