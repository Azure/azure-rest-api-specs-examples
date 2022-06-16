import com.azure.core.util.Context;

/** Samples for UsageDetails List. */
public final class Main {
    /*
     * x-ms-original-file: specification/consumption/resource-manager/Microsoft.Consumption/stable/2021-10-01/examples/UsageDetailsListByMCABillingAccount.json
     */
    /**
     * Sample code: BillingAccountUsageDetailsList-Modern.
     *
     * @param manager Entry point to ConsumptionManager.
     */
    public static void billingAccountUsageDetailsListModern(
        com.azure.resourcemanager.consumption.ConsumptionManager manager) {
        manager
            .usageDetails()
            .list("providers/Microsoft.Billing/BillingAccounts/1234:56789", null, null, null, null, null, Context.NONE);
    }
}
