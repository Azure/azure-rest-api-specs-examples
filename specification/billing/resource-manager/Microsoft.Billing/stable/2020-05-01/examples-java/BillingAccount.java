/** Samples for BillingAccounts Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/stable/2020-05-01/examples/BillingAccount.json
     */
    /**
     * Sample code: BillingAccounts.
     *
     * @param manager Entry point to BillingManager.
     */
    public static void billingAccounts(com.azure.resourcemanager.billing.BillingManager manager) {
        manager.billingAccounts().getWithResponse("{billingAccountName}", null, com.azure.core.util.Context.NONE);
    }
}
