
/**
 * Samples for OracleSubscriptions ListCloudAccountDetails.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01/oracleSubscriptions_listCloudAccountDetails.json
     */
    /**
     * Sample code: OracleSubscriptions_listCloudAccountDetails.
     * 
     * @param manager Entry point to OracleDatabaseManager.
     */
    public static void oracleSubscriptionsListCloudAccountDetails(
        com.azure.resourcemanager.oracledatabase.OracleDatabaseManager manager) {
        manager.oracleSubscriptions().listCloudAccountDetails(com.azure.core.util.Context.NONE);
    }
}
