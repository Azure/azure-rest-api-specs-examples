/** Samples for BillingSubscriptions ListByBillingAccount. */
public final class Main {
    /*
     * x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/stable/2020-05-01/examples/BillingSubscriptionsListByBillingAccount.json
     */
    /**
     * Sample code: BillingSubscriptionsListByBillingAccount.
     *
     * @param manager Entry point to BillingManager.
     */
    public static void billingSubscriptionsListByBillingAccount(
        com.azure.resourcemanager.billing.BillingManager manager) {
        manager.billingSubscriptions().listByBillingAccount("{billingAccountName}", com.azure.core.util.Context.NONE);
    }
}
