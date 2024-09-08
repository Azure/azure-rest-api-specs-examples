
/**
 * Samples for BillingSubscriptions Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/billingSubscriptionsDelete.
     * json
     */
    /**
     * Sample code: BillingSubscriptionsDelete.
     * 
     * @param manager Entry point to BillingManager.
     */
    public static void billingSubscriptionsDelete(com.azure.resourcemanager.billing.BillingManager manager) {
        manager.billingSubscriptions().delete(
            "00000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31",
            "11111111-1111-1111-1111-111111111111", com.azure.core.util.Context.NONE);
    }
}
