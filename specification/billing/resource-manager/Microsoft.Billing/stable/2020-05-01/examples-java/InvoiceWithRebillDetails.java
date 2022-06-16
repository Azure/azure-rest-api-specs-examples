import com.azure.core.util.Context;

/** Samples for Invoices Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/stable/2020-05-01/examples/InvoiceWithRebillDetails.json
     */
    /**
     * Sample code: InvoiceWithRebillDetails.
     *
     * @param manager Entry point to BillingManager.
     */
    public static void invoiceWithRebillDetails(com.azure.resourcemanager.billing.BillingManager manager) {
        manager.invoices().getWithResponse("{billingAccountName}", "{invoiceName}", Context.NONE);
    }
}
