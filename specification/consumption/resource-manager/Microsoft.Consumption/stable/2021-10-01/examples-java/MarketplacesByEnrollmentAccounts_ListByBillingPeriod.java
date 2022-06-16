import com.azure.core.util.Context;

/** Samples for Marketplaces List. */
public final class Main {
    /*
     * x-ms-original-file: specification/consumption/resource-manager/Microsoft.Consumption/stable/2021-10-01/examples/MarketplacesByEnrollmentAccounts_ListByBillingPeriod.json
     */
    /**
     * Sample code: EnrollmentAccountMarketplacesListForBillingPeriod.
     *
     * @param manager Entry point to ConsumptionManager.
     */
    public static void enrollmentAccountMarketplacesListForBillingPeriod(
        com.azure.resourcemanager.consumption.ConsumptionManager manager) {
        manager
            .marketplaces()
            .list("providers/Microsoft.Billing/enrollmentAccounts/123456", null, null, null, Context.NONE);
    }
}
