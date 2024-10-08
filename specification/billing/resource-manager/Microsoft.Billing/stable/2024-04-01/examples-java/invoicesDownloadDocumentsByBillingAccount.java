
import com.azure.resourcemanager.billing.models.DocumentDownloadRequest;
import java.util.Arrays;

/**
 * Samples for Invoices DownloadDocumentsByBillingAccount.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/
     * invoicesDownloadDocumentsByBillingAccount.json
     */
    /**
     * Sample code: InvoicesDownloadDocumentsByBillingAccount.
     * 
     * @param manager Entry point to BillingManager.
     */
    public static void
        invoicesDownloadDocumentsByBillingAccount(com.azure.resourcemanager.billing.BillingManager manager) {
        manager.invoices().downloadDocumentsByBillingAccount(
            "00000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31",
            Arrays.asList(new DocumentDownloadRequest().withDocumentName("12345678").withInvoiceName("G123456789"),
                new DocumentDownloadRequest().withDocumentName("12345678").withInvoiceName("G987654321")),
            com.azure.core.util.Context.NONE);
    }
}
