
import com.azure.resourcemanager.billing.fluent.models.BillingSubscriptionAliasInner;

/**
 * Samples for BillingSubscriptionsAliases CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/
     * billingSubscriptionAliasCreateOrUpdate.json
     */
    /**
     * Sample code: BillingSubscriptionAliasCreateOrUpdate.
     * 
     * @param manager Entry point to BillingManager.
     */
    public static void
        billingSubscriptionAliasCreateOrUpdate(com.azure.resourcemanager.billing.BillingManager manager) {
        manager.billingSubscriptionsAliases().createOrUpdate(
            "00000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31",
            "c356b7c7-7545-4686-b843-c1a49cf853fc", new BillingSubscriptionAliasInner().withBillingFrequency("P1M")
                .withDisplayName("Subscription 3").withQuantity(1L).withSkuId("0001").withTermDuration("P1M"),
            com.azure.core.util.Context.NONE);
    }
}
