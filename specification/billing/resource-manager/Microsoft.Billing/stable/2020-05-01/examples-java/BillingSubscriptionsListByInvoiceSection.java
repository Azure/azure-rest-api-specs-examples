import com.azure.core.util.Context;

/** Samples for BillingSubscriptions ListByInvoiceSection. */
public final class Main {
    /*
     * x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/stable/2020-05-01/examples/BillingSubscriptionsListByInvoiceSection.json
     */
    /**
     * Sample code: BillingSubscriptionsListByInvoiceSection.
     *
     * @param manager Entry point to BillingManager.
     */
    public static void billingSubscriptionsListByInvoiceSection(
        com.azure.resourcemanager.billing.BillingManager manager) {
        manager
            .billingSubscriptions()
            .listByInvoiceSection("{billingAccountName}", "{billingProfileName}", "{invoiceSectionName}", Context.NONE);
    }
}
