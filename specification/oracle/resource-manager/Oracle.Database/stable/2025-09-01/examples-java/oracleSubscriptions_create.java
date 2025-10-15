
import com.azure.resourcemanager.oracledatabase.fluent.models.OracleSubscriptionInner;
import com.azure.resourcemanager.oracledatabase.models.OracleSubscriptionProperties;
import com.azure.resourcemanager.oracledatabase.models.Plan;

/**
 * Samples for OracleSubscriptions CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01/oracleSubscriptions_create.json
     */
    /**
     * Sample code: OracleSubscriptions_createOrUpdate.
     * 
     * @param manager Entry point to OracleDatabaseManager.
     */
    public static void
        oracleSubscriptionsCreateOrUpdate(com.azure.resourcemanager.oracledatabase.OracleDatabaseManager manager) {
        manager.oracleSubscriptions()
            .createOrUpdate(
                new OracleSubscriptionInner().withProperties(new OracleSubscriptionProperties())
                    .withPlan(new Plan().withName("plan1").withPublisher("publisher1").withProduct("product1")
                        .withPromotionCode("fakeTokenPlaceholder").withVersion("alpha")),
                com.azure.core.util.Context.NONE);
    }
}
