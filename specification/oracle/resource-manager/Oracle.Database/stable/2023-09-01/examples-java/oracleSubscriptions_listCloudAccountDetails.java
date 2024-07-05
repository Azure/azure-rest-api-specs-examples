
/**
 * Samples for OracleSubscriptions ListCloudAccountDetails.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/oracle/resource-manager/Oracle.Database/stable/2023-09-01/examples/
     * oracleSubscriptions_listCloudAccountDetails.json
     */
    /**
     * Sample code: List Cloud Account details for the Oracle Subscription.
     * 
     * @param manager Entry point to OracleDatabaseManager.
     */
    public static void listCloudAccountDetailsForTheOracleSubscription(
        com.azure.resourcemanager.oracledatabase.OracleDatabaseManager manager) {
        manager.oracleSubscriptions().listCloudAccountDetails(com.azure.core.util.Context.NONE);
    }
}
