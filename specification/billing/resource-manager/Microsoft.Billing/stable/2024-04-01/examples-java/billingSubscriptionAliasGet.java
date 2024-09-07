
/**
 * Samples for BillingSubscriptionsAliases Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/billingSubscriptionAliasGet.
     * json
     */
    /**
     * Sample code: BillingSubscriptionAliasGet.
     * 
     * @param manager Entry point to BillingManager.
     */
    public static void billingSubscriptionAliasGet(com.azure.resourcemanager.billing.BillingManager manager) {
        manager.billingSubscriptionsAliases().getWithResponse(
            "00000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31",
            "c356b7c7-7545-4686-b843-c1a49cf853fc", com.azure.core.util.Context.NONE);
    }
}
