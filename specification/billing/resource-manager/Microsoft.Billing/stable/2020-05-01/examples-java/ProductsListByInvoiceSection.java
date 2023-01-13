/** Samples for Products ListByInvoiceSection. */
public final class Main {
    /*
     * x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/stable/2020-05-01/examples/ProductsListByInvoiceSection.json
     */
    /**
     * Sample code: ProductsListByInvoiceSection.
     *
     * @param manager Entry point to BillingManager.
     */
    public static void productsListByInvoiceSection(com.azure.resourcemanager.billing.BillingManager manager) {
        manager
            .products()
            .listByInvoiceSection(
                "{billingAccountName}",
                "{billingProfileName}",
                "{invoiceSectionName}",
                null,
                com.azure.core.util.Context.NONE);
    }
}
