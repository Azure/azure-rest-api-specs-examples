
/**
 * Samples for OracleSubscriptions Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01/oracleSubscriptions_delete.json
     */
    /**
     * Sample code: OracleSubscriptions_delete.
     * 
     * @param manager Entry point to OracleDatabaseManager.
     */
    public static void
        oracleSubscriptionsDelete(com.azure.resourcemanager.oracledatabase.OracleDatabaseManager manager) {
        manager.oracleSubscriptions().delete(com.azure.core.util.Context.NONE);
    }
}
