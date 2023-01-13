/** Samples for BillingAccounts List. */
public final class Main {
    /*
     * x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/stable/2020-05-01/examples/BillingAccountsListWithExpand.json
     */
    /**
     * Sample code: BillingAccountsListWithExpand.
     *
     * @param manager Entry point to BillingManager.
     */
    public static void billingAccountsListWithExpand(com.azure.resourcemanager.billing.BillingManager manager) {
        manager
            .billingAccounts()
            .list("soldTo,billingProfiles,billingProfiles/invoiceSections", com.azure.core.util.Context.NONE);
    }
}
