
/**
 * Samples for OracleSubscriptions ListSaasSubscriptionDetails.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/oracle/resource-manager/Oracle.Database/preview/2023-09-01-preview/examples/
     * oracleSubscriptions_listSaasSubscriptionDetails.json
     */
    /**
     * Sample code: List Saas Subscription details for the Oracle Subscription.
     * 
     * @param manager Entry point to OracleDatabaseManager.
     */
    public static void listSaasSubscriptionDetailsForTheOracleSubscription(
        com.azure.resourcemanager.oracledatabase.OracleDatabaseManager manager) {
        manager.oracleSubscriptions().listSaasSubscriptionDetails(com.azure.core.util.Context.NONE);
    }
}
