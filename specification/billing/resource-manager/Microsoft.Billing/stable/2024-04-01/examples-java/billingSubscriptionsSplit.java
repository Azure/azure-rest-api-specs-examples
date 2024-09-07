
import com.azure.resourcemanager.billing.models.BillingSubscriptionSplitRequest;

/**
 * Samples for BillingSubscriptions Split.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/billingSubscriptionsSplit.
     * json
     */
    /**
     * Sample code: BillingSubscriptionsSplit.
     * 
     * @param manager Entry point to BillingManager.
     */
    public static void billingSubscriptionsSplit(com.azure.resourcemanager.billing.BillingManager manager) {
        manager.billingSubscriptions()
            .split("00000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31",
                "11111111-1111-1111-1111-111111111111",
                new BillingSubscriptionSplitRequest().withTargetProductTypeId("XYZ56789").withTargetSkuId("0001")
                    .withQuantity(1).withTermDuration("P1M").withBillingFrequency("P1M"),
                com.azure.core.util.Context.NONE);
    }
}
