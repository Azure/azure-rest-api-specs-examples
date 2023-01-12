/** Samples for BillingSubscriptions Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/stable/2020-05-01/examples/BillingSubscription.json
     */
    /**
     * Sample code: BillingSubscription.
     *
     * @param manager Entry point to BillingManager.
     */
    public static void billingSubscription(com.azure.resourcemanager.billing.BillingManager manager) {
        manager.billingSubscriptions().getWithResponse("{billingAccountName}", com.azure.core.util.Context.NONE);
    }
}
