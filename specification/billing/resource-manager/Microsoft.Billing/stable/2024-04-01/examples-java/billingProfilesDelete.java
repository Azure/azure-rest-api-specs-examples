
/**
 * Samples for BillingProfiles Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/billingProfilesDelete.json
     */
    /**
     * Sample code: BillingProfilesDelete.
     * 
     * @param manager Entry point to BillingManager.
     */
    public static void billingProfilesDelete(com.azure.resourcemanager.billing.BillingManager manager) {
        manager.billingProfiles().delete(
            "00000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31", "xxxx-xxxx-xxx-xxx",
            com.azure.core.util.Context.NONE);
    }
}
