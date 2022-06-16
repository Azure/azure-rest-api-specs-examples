import com.azure.core.util.Context;

/** Samples for UsageDetails List. */
public final class Main {
    /*
     * x-ms-original-file: specification/consumption/resource-manager/Microsoft.Consumption/stable/2021-10-01/examples/UsageDetailsListForBillingPeriodByEnrollmentAccount.json
     */
    /**
     * Sample code: EnrollmentAccountUsageDetailsListForBillingPeriod-Legacy.
     *
     * @param manager Entry point to ConsumptionManager.
     */
    public static void enrollmentAccountUsageDetailsListForBillingPeriodLegacy(
        com.azure.resourcemanager.consumption.ConsumptionManager manager) {
        manager
            .usageDetails()
            .list("providers/Microsoft.Billing/EnrollmentAccounts/1234", null, null, null, null, null, Context.NONE);
    }
}
