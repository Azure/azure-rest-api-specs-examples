/** Samples for Invoices DownloadInvoice. */
public final class Main {
    /*
     * x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/stable/2020-05-01/examples/ModernInvoiceDownload.json
     */
    /**
     * Sample code: InvoiceDownload.
     *
     * @param manager Entry point to BillingManager.
     */
    public static void invoiceDownload(com.azure.resourcemanager.billing.BillingManager manager) {
        manager
            .invoices()
            .downloadInvoice("{billingAccountName}", "{invoiceName}", "DRS_12345", com.azure.core.util.Context.NONE);
    }
}
