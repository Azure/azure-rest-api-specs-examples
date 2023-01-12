/** Samples for BillingProfiles Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/stable/2020-05-01/examples/BillingProfile.json
     */
    /**
     * Sample code: BillingProfile.
     *
     * @param manager Entry point to BillingManager.
     */
    public static void billingProfile(com.azure.resourcemanager.billing.BillingManager manager) {
        manager
            .billingProfiles()
            .getWithResponse("{billingAccountName}", "{billingProfileName}", null, com.azure.core.util.Context.NONE);
    }
}
