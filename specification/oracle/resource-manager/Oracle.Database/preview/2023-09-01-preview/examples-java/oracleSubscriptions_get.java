
/**
 * Samples for OracleSubscriptions Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/oracle/resource-manager/Oracle.Database/preview/2023-09-01-preview/examples/oracleSubscriptions_get
     * .json
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
