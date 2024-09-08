
import com.azure.resourcemanager.billing.models.CancelSubscriptionRequest;
import com.azure.resourcemanager.billing.models.CancellationReason;

/**
 * Samples for BillingSubscriptions Cancel.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/billingSubscriptionsCancel.
     * json
     */
    /**
     * Sample code: BillingSubscriptionsCancel.
     * 
     * @param manager Entry point to BillingManager.
     */
    public static void billingSubscriptionsCancel(com.azure.resourcemanager.billing.BillingManager manager) {
        manager.billingSubscriptions().cancel(
            "00000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31",
            "11111111-1111-1111-1111-111111111111",
            new CancelSubscriptionRequest().withCancellationReason(CancellationReason.COMPROMISE)
                .withCustomerId("11111111-1111-1111-1111-111111111111"),
            com.azure.core.util.Context.NONE);
    }
}
