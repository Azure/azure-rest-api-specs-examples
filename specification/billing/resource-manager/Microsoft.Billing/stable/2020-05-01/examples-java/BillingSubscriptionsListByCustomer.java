/** Samples for BillingSubscriptions ListByCustomer. */
public final class Main {
    /*
     * x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/stable/2020-05-01/examples/BillingSubscriptionsListByCustomer.json
     */
    /**
     * Sample code: BillingSubscriptionsListByCustomer.
     *
     * @param manager Entry point to BillingManager.
     */
    public static void billingSubscriptionsListByCustomer(com.azure.resourcemanager.billing.BillingManager manager) {
        manager
            .billingSubscriptions()
            .listByCustomer("{billingAccountName}", "{customerName}", com.azure.core.util.Context.NONE);
    }
}
