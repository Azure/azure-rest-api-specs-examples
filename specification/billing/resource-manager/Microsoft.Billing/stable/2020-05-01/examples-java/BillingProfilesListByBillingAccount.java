/** Samples for BillingProfiles ListByBillingAccount. */
public final class Main {
    /*
     * x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/stable/2020-05-01/examples/BillingProfilesListByBillingAccount.json
     */
    /**
     * Sample code: BillingProfilesListByBillingAccount.
     *
     * @param manager Entry point to BillingManager.
     */
    public static void billingProfilesListByBillingAccount(com.azure.resourcemanager.billing.BillingManager manager) {
        manager.billingProfiles().listByBillingAccount("{billingAccountName}", null, com.azure.core.util.Context.NONE);
    }
}
