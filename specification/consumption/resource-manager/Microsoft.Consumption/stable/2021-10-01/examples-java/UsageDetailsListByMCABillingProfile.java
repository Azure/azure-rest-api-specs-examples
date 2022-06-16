import com.azure.core.util.Context;

/** Samples for UsageDetails List. */
public final class Main {
    /*
     * x-ms-original-file: specification/consumption/resource-manager/Microsoft.Consumption/stable/2021-10-01/examples/UsageDetailsListByMCABillingProfile.json
     */
    /**
     * Sample code: BillingProfileUsageDetailsList-Modern.
     *
     * @param manager Entry point to ConsumptionManager.
     */
    public static void billingProfileUsageDetailsListModern(
        com.azure.resourcemanager.consumption.ConsumptionManager manager) {
        manager
            .usageDetails()
            .list(
                "providers/Microsoft.Billing/BillingAccounts/1234:56789/billingProfiles/2468",
                null,
                null,
                null,
                null,
                null,
                Context.NONE);
    }
}
