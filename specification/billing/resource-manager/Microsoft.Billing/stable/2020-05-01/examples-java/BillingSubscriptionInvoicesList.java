import com.azure.core.util.Context;

/** Samples for Invoices ListByBillingSubscription. */
public final class Main {
    /*
     * x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/stable/2020-05-01/examples/BillingSubscriptionInvoicesList.json
     */
    /**
     * Sample code: BillingSubscriptionsListByBillingAccount.
     *
     * @param manager Entry point to BillingManager.
     */
    public static void billingSubscriptionsListByBillingAccount(
        com.azure.resourcemanager.billing.BillingManager manager) {
        manager.invoices().listByBillingSubscription("2018-01-01", "2018-06-30", Context.NONE);
    }
}
