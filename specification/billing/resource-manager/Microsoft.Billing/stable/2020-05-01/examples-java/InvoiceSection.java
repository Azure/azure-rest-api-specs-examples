import com.azure.core.util.Context;

/** Samples for InvoiceSections Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/stable/2020-05-01/examples/InvoiceSection.json
     */
    /**
     * Sample code: InvoiceSection.
     *
     * @param manager Entry point to BillingManager.
     */
    public static void invoiceSection(com.azure.resourcemanager.billing.BillingManager manager) {
        manager
            .invoiceSections()
            .getWithResponse("{billingAccountName}", "{billingProfileName}", "{invoiceSectionName}", Context.NONE);
    }
}
