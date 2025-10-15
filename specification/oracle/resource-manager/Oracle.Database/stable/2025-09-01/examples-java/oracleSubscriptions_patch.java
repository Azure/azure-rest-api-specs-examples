
import com.azure.resourcemanager.oracledatabase.models.OracleSubscriptionUpdate;

/**
 * Samples for OracleSubscriptions Update.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01/oracleSubscriptions_patch.json
     */
    /**
     * Sample code: OracleSubscriptions_update.
     * 
     * @param manager Entry point to OracleDatabaseManager.
     */
    public static void
        oracleSubscriptionsUpdate(com.azure.resourcemanager.oracledatabase.OracleDatabaseManager manager) {
        manager.oracleSubscriptions().update(new OracleSubscriptionUpdate(), com.azure.core.util.Context.NONE);
    }
}
