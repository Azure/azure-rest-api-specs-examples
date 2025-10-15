
/**
 * Samples for OracleSubscriptions ListActivationLinks.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01/oracleSubscriptions_listActivationLinks.json
     */
    /**
     * Sample code: OracleSubscriptions_listActivationLinks.
     * 
     * @param manager Entry point to OracleDatabaseManager.
     */
    public static void
        oracleSubscriptionsListActivationLinks(com.azure.resourcemanager.oracledatabase.OracleDatabaseManager manager) {
        manager.oracleSubscriptions().listActivationLinks(com.azure.core.util.Context.NONE);
    }
}
