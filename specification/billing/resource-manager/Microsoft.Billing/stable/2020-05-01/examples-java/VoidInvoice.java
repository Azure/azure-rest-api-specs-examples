/** Samples for Invoices Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/stable/2020-05-01/examples/VoidInvoice.json
     */
    /**
     * Sample code: VoidInvoice.
     *
     * @param manager Entry point to BillingManager.
     */
    public static void voidInvoice(com.azure.resourcemanager.billing.BillingManager manager) {
        manager.invoices().getWithResponse("{billingAccountName}", "{invoiceName}", com.azure.core.util.Context.NONE);
    }
}
