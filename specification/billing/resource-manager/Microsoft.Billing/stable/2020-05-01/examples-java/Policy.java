/** Samples for Policies GetByBillingProfile. */
public final class Main {
    /*
     * x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/stable/2020-05-01/examples/Policy.json
     */
    /**
     * Sample code: PolicyByBillingProfile.
     *
     * @param manager Entry point to BillingManager.
     */
    public static void policyByBillingProfile(com.azure.resourcemanager.billing.BillingManager manager) {
        manager
            .policies()
            .getByBillingProfileWithResponse(
                "{billingAccountName}", "{billingProfileName}", com.azure.core.util.Context.NONE);
    }
}
