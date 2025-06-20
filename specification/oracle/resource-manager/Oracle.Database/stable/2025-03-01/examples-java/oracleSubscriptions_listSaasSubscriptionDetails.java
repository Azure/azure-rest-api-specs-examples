
/**
 * Samples for OracleSubscriptions ListSaasSubscriptionDetails.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-03-01/oracleSubscriptions_listSaasSubscriptionDetails.json
     */
    /**
     * Sample code: OracleSubscriptions_listSaasSubscriptionDetails.
     * 
     * @param manager Entry point to OracleDatabaseManager.
     */
    public static void oracleSubscriptionsListSaasSubscriptionDetails(
        com.azure.resourcemanager.oracledatabase.OracleDatabaseManager manager) {
        manager.oracleSubscriptions().listSaasSubscriptionDetails(com.azure.core.util.Context.NONE);
    }
}
