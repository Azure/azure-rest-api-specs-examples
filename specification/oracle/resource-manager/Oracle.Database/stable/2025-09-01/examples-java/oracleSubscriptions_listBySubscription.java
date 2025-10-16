
/**
 * Samples for OracleSubscriptions List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01/oracleSubscriptions_listBySubscription.json
     */
    /**
     * Sample code: OracleSubscriptions_listBySubscription.
     * 
     * @param manager Entry point to OracleDatabaseManager.
     */
    public static void
        oracleSubscriptionsListBySubscription(com.azure.resourcemanager.oracledatabase.OracleDatabaseManager manager) {
        manager.oracleSubscriptions().list(com.azure.core.util.Context.NONE);
    }
}
