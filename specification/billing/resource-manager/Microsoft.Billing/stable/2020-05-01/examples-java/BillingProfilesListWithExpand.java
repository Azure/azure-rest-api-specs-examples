/** Samples for BillingProfiles ListByBillingAccount. */
public final class Main {
    /*
     * x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/stable/2020-05-01/examples/BillingProfilesListWithExpand.json
     */
    /**
     * Sample code: BillingProfilesListWithExpand.
     *
     * @param manager Entry point to BillingManager.
     */
    public static void billingProfilesListWithExpand(com.azure.resourcemanager.billing.BillingManager manager) {
        manager
            .billingProfiles()
            .listByBillingAccount("{billingAccountName}", "invoiceSections", com.azure.core.util.Context.NONE);
    }
}
