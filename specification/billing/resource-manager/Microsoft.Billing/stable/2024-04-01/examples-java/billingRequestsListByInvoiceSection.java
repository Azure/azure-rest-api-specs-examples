
/**
 * Samples for BillingRequests ListByInvoiceSection.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/
     * billingRequestsListByInvoiceSection.json
     */
    /**
     * Sample code: BillingRequestsListByInvoiceSection.
     * 
     * @param manager Entry point to BillingManager.
     */
    public static void billingRequestsListByInvoiceSection(com.azure.resourcemanager.billing.BillingManager manager) {
        manager.billingRequests().listByInvoiceSection(
            "00000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31", "xxxx-xxxx-xxx-xxx",
            "yyyy-yyyy-yyy-yyy", null, null, null, null, null, null, com.azure.core.util.Context.NONE);
    }
}
