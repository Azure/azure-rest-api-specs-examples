import com.azure.core.util.Context;

/** Samples for BillingSubscriptions ListByBillingProfile. */
public final class Main {
    /*
     * x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/stable/2020-05-01/examples/BillingSubscriptionsListByBillingProfile.json
     */
    /**
     * Sample code: BillingSubscriptionsListByBillingProfile.
     *
     * @param manager Entry point to BillingManager.
     */
    public static void billingSubscriptionsListByBillingProfile(
        com.azure.resourcemanager.billing.BillingManager manager) {
        manager
            .billingSubscriptions()
            .listByBillingProfile("{billingAccountName}", "{billingProfileName}", Context.NONE);
    }
}
