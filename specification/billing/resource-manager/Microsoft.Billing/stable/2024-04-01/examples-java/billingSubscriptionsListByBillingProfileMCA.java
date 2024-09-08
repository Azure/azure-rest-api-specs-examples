
/**
 * Samples for BillingSubscriptions ListByBillingProfile.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/
     * billingSubscriptionsListByBillingProfileMCA.json
     */
    /**
     * Sample code: BillingSubscriptionsListByBillingProfileMCA.
     * 
     * @param manager Entry point to BillingManager.
     */
    public static void
        billingSubscriptionsListByBillingProfileMCA(com.azure.resourcemanager.billing.BillingManager manager) {
        manager.billingSubscriptions().listByBillingProfile(
            "a1a9c77e-4cec-4a6c-a089-867d973a6074:a80d3b1f-c626-4e5e-82ed-1173bd91c838_2019-05-31",
            "ea36e548-1505-41db-bebc-46fff3d37998", null, null, null, null, null, null, null, null,
            com.azure.core.util.Context.NONE);
    }
}
