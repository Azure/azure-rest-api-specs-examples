import com.azure.core.util.Context;

/** Samples for Invoices Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/stable/2020-05-01/examples/CreditNote.json
     */
    /**
     * Sample code: CreditNote.
     *
     * @param manager Entry point to BillingManager.
     */
    public static void creditNote(com.azure.resourcemanager.billing.BillingManager manager) {
        manager.invoices().getWithResponse("{billingAccountName}", "{invoiceName}", Context.NONE);
    }
}
