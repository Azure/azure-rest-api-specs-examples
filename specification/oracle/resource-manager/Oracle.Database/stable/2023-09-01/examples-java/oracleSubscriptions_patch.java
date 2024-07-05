
import com.azure.resourcemanager.oracledatabase.models.OracleSubscriptionUpdate;

/**
 * Samples for OracleSubscriptions Update.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/oracle/resource-manager/Oracle.Database/stable/2023-09-01/examples/oracleSubscriptions_patch.json
     */
    /**
     * Sample code: Patch Oracle Subscription.
     * 
     * @param manager Entry point to OracleDatabaseManager.
     */
    public static void patchOracleSubscription(com.azure.resourcemanager.oracledatabase.OracleDatabaseManager manager) {
        manager.oracleSubscriptions().update(new OracleSubscriptionUpdate(), com.azure.core.util.Context.NONE);
    }
}
