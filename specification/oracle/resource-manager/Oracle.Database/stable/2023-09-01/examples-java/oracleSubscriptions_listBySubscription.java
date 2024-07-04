
/**
 * Samples for OracleSubscriptions List.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/oracle/resource-manager/Oracle.Database/stable/2023-09-01/examples/
     * oracleSubscriptions_listBySubscription.json
     */
    /**
     * Sample code: List Oracle Subscriptions by subscription.
     * 
     * @param manager Entry point to OracleDatabaseManager.
     */
    public static void
        listOracleSubscriptionsBySubscription(com.azure.resourcemanager.oracledatabase.OracleDatabaseManager manager) {
        manager.oracleSubscriptions().list(com.azure.core.util.Context.NONE);
    }
}
