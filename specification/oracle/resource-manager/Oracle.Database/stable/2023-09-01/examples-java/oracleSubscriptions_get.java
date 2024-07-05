
/**
 * Samples for OracleSubscriptions Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/oracle/resource-manager/Oracle.Database/stable/2023-09-01/examples/oracleSubscriptions_get.json
     */
    /**
     * Sample code: Get Oracle Subscription.
     * 
     * @param manager Entry point to OracleDatabaseManager.
     */
    public static void getOracleSubscription(com.azure.resourcemanager.oracledatabase.OracleDatabaseManager manager) {
        manager.oracleSubscriptions().getWithResponse(com.azure.core.util.Context.NONE);
    }
}
