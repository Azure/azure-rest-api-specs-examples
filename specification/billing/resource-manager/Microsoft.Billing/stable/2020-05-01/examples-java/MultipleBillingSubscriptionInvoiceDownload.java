import java.util.Arrays;

/** Samples for Invoices DownloadMultipleBillingSubscriptionInvoices. */
public final class Main {
    /*
     * x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/stable/2020-05-01/examples/MultipleBillingSubscriptionInvoiceDownload.json
     */
    /**
     * Sample code: BillingSubscriptionInvoiceDownload.
     *
     * @param manager Entry point to BillingManager.
     */
    public static void billingSubscriptionInvoiceDownload(com.azure.resourcemanager.billing.BillingManager manager) {
        manager
            .invoices()
            .downloadMultipleBillingSubscriptionInvoices(
                Arrays
                    .asList(
                        "https://management.azure.com/providers/Microsoft.Billing/billingAccounts/default/billingSubscriptions/{subscriptionId}/invoices/{invoiceName}/download?downloadToken={downloadToken}&useCache=True&api-version=2020-05-01",
                        "https://management.azure.com/providers/Microsoft.Billing/billingAccounts/default/billingSubscriptions/{subscriptionId}/invoices/{invoiceName}/download?downloadToken={downloadToken}&useCache=True&api-version=2020-05-01",
                        "https://management.azure.com/providers/Microsoft.Billing/billingAccounts/default/billingSubscriptions/{subscriptionId}/invoices/{invoiceName}/download?downloadToken={downloadToken}&useCache=True&api-version=2020-05-01"),
                com.azure.core.util.Context.NONE);
    }
}
