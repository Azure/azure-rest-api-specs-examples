
import com.azure.resourcemanager.oracledatabase.fluent.models.OracleSubscriptionInner;
import com.azure.resourcemanager.oracledatabase.models.OracleSubscriptionProperties;
import com.azure.resourcemanager.oracledatabase.models.Plan;

/**
 * Samples for OracleSubscriptions CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/oracle/resource-manager/Oracle.Database/stable/2023-09-01/examples/oracleSubscriptions_create.json
     */
    /**
     * Sample code: Create or Update Oracle Subscription.
     * 
     * @param manager Entry point to OracleDatabaseManager.
     */
    public static void
        createOrUpdateOracleSubscription(com.azure.resourcemanager.oracledatabase.OracleDatabaseManager manager) {
        manager.oracleSubscriptions()
            .createOrUpdate(
                new OracleSubscriptionInner().withProperties(new OracleSubscriptionProperties())
                    .withPlan(new Plan().withName("plan1").withPublisher("publisher1").withProduct("product1")
                        .withPromotionCode("fakeTokenPlaceholder").withVersion("alpha")),
                com.azure.core.util.Context.NONE);
    }
}
