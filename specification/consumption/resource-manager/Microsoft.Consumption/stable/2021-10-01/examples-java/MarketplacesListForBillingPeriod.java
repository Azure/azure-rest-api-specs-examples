import com.azure.core.util.Context;

/** Samples for Marketplaces List. */
public final class Main {
    /*
     * x-ms-original-file: specification/consumption/resource-manager/Microsoft.Consumption/stable/2021-10-01/examples/MarketplacesListForBillingPeriod.json
     */
    /**
     * Sample code: SubscriptionMarketplacesListForBillingPeriod.
     *
     * @param manager Entry point to ConsumptionManager.
     */
    public static void subscriptionMarketplacesListForBillingPeriod(
        com.azure.resourcemanager.consumption.ConsumptionManager manager) {
        manager
            .marketplaces()
            .list("subscriptions/00000000-0000-0000-0000-000000000000", null, null, null, Context.NONE);
    }
}
