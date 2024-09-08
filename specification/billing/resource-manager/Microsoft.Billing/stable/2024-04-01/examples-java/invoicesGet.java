
/**
 * Samples for Invoices Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/invoicesGet.json
     */
    /**
     * Sample code: InvoicesGet.
     * 
     * @param manager Entry point to BillingManager.
     */
    public static void invoicesGet(com.azure.resourcemanager.billing.BillingManager manager) {
        manager.invoices().getWithResponse("G123456789", com.azure.core.util.Context.NONE);
    }
}
