
import com.azure.resourcemanager.billing.models.BillingSubscriptionMergeRequest;

/**
 * Samples for BillingSubscriptions Merge.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/billingSubscriptionsMerge.
     * json
     */
    /**
     * Sample code: BillingSubscriptionsMerge.
     * 
     * @param manager Entry point to BillingManager.
     */
    public static void billingSubscriptionsMerge(com.azure.resourcemanager.billing.BillingManager manager) {
        manager.billingSubscriptions().merge(
            "00000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31",
            "11111111-1111-1111-1111-111111111111",
            new BillingSubscriptionMergeRequest()
                .withTargetBillingSubscriptionName("22222222-2222-2222-2222-222222222222").withQuantity(1),
            com.azure.core.util.Context.NONE);
    }
}
