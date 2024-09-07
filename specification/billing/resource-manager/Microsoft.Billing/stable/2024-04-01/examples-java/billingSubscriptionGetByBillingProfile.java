
/**
 * Samples for BillingSubscriptions GetByBillingProfile.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/
     * billingSubscriptionGetByBillingProfile.json
     */
    /**
     * Sample code: BillingSubscriptionGetByBillingProfile.
     * 
     * @param manager Entry point to BillingManager.
     */
    public static void
        billingSubscriptionGetByBillingProfile(com.azure.resourcemanager.billing.BillingManager manager) {
        manager.billingSubscriptions().getByBillingProfileWithResponse("pcn.94077792", "6478903",
            "6b96d3f2-9008-4a9d-912f-f87744185aa3", null, com.azure.core.util.Context.NONE);
    }
}
