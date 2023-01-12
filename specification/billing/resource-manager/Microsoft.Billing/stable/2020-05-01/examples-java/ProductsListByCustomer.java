/** Samples for Products ListByCustomer. */
public final class Main {
    /*
     * x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/stable/2020-05-01/examples/ProductsListByCustomer.json
     */
    /**
     * Sample code: ProductsListByInvoiceSection.
     *
     * @param manager Entry point to BillingManager.
     */
    public static void productsListByInvoiceSection(com.azure.resourcemanager.billing.BillingManager manager) {
        manager.products().listByCustomer("{billingAccountName}", "{customerName}", com.azure.core.util.Context.NONE);
    }
}
