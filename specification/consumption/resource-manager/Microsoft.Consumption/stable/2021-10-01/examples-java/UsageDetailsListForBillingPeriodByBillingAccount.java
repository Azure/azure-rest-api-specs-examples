import com.azure.core.util.Context;

/** Samples for UsageDetails List. */
public final class Main {
    /*
     * x-ms-original-file: specification/consumption/resource-manager/Microsoft.Consumption/stable/2021-10-01/examples/UsageDetailsListForBillingPeriodByBillingAccount.json
     */
    /**
     * Sample code: BillingAccountUsageDetailsListForBillingPeriod-Legacy.
     *
     * @param manager Entry point to ConsumptionManager.
     */
    public static void billingAccountUsageDetailsListForBillingPeriodLegacy(
        com.azure.resourcemanager.consumption.ConsumptionManager manager) {
        manager
            .usageDetails()
            .list("providers/Microsoft.Billing/BillingAccounts/1234", null, null, null, null, null, Context.NONE);
    }
}
