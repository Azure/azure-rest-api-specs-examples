
/**
 * Samples for BillingPermissions ListByInvoiceSection.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/
     * billingPermissionsListByInvoiceSection.json
     */
    /**
     * Sample code: BillingPermissionsListByInvoiceSection.
     * 
     * @param manager Entry point to BillingManager.
     */
    public static void
        billingPermissionsListByInvoiceSection(com.azure.resourcemanager.billing.BillingManager manager) {
        manager.billingPermissions().listByInvoiceSection(
            "10000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31", "xxxx-xxxx-xxx-xxx",
            "XXXX-XXXX-XXX-XXX", com.azure.core.util.Context.NONE);
    }
}
