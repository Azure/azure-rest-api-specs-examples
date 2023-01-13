/** Samples for Invoices GetBySubscriptionAndInvoiceId. */
public final class Main {
    /*
     * x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/stable/2020-05-01/examples/BillingSubscriptionInvoice.json
     */
    /**
     * Sample code: BillingSubscriptionsListByBillingAccount.
     *
     * @param manager Entry point to BillingManager.
     */
    public static void billingSubscriptionsListByBillingAccount(
        com.azure.resourcemanager.billing.BillingManager manager) {
        manager.invoices().getBySubscriptionAndInvoiceIdWithResponse("{invoiceName}", com.azure.core.util.Context.NONE);
    }
}
