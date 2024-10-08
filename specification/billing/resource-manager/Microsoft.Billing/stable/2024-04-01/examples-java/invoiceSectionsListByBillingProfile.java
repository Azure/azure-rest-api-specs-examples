
/**
 * Samples for InvoiceSections ListByBillingProfile.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/
     * invoiceSectionsListByBillingProfile.json
     */
    /**
     * Sample code: InvoiceSectionsListByBillingProfile.
     * 
     * @param manager Entry point to BillingManager.
     */
    public static void invoiceSectionsListByBillingProfile(com.azure.resourcemanager.billing.BillingManager manager) {
        manager.invoiceSections().listByBillingProfile(
            "00000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31", "xxxx-xxxx-xxx-xxx",
            true, null, null, null, null, null, null, com.azure.core.util.Context.NONE);
    }
}
