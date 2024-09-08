
import com.azure.resourcemanager.billing.models.DocumentDownloadRequest;
import java.util.Arrays;

/**
 * Samples for Invoices DownloadDocumentsByBillingSubscription.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/
     * invoicesDownloadDocumentsByBillingSubscription.json
     */
    /**
     * Sample code: InvoicesDownloadDocumentsByBillingSubscription.
     * 
     * @param manager Entry point to BillingManager.
     */
    public static void
        invoicesDownloadDocumentsByBillingSubscription(com.azure.resourcemanager.billing.BillingManager manager) {
        manager.invoices().downloadDocumentsByBillingSubscription(
            Arrays.asList(new DocumentDownloadRequest().withDocumentName("12345678").withInvoiceName("E123456789"),
                new DocumentDownloadRequest().withDocumentName("12345678").withInvoiceName("E987654321")),
            com.azure.core.util.Context.NONE);
    }
}
