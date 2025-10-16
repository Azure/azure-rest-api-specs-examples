
/**
 * Samples for OracleSubscriptions Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01/oracleSubscriptions_get.json
     */
    /**
     * Sample code: OracleSubscriptions_get.
     * 
     * @param manager Entry point to OracleDatabaseManager.
     */
    public static void oracleSubscriptionsGet(com.azure.resourcemanager.oracledatabase.OracleDatabaseManager manager) {
        manager.oracleSubscriptions().getWithResponse(com.azure.core.util.Context.NONE);
    }
}
