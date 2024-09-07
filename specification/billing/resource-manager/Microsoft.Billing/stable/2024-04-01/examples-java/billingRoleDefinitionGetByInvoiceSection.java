
/**
 * Samples for BillingRoleDefinition GetByInvoiceSection.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/
     * billingRoleDefinitionGetByInvoiceSection.json
     */
    /**
     * Sample code: BillingRoleDefinitionGetByInvoiceSection.
     * 
     * @param manager Entry point to BillingManager.
     */
    public static void
        billingRoleDefinitionGetByInvoiceSection(com.azure.resourcemanager.billing.BillingManager manager) {
        manager.billingRoleDefinitions().getByInvoiceSectionWithResponse(
            "10000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31", "xxxx-xxxx-xxx-xxx",
            "yyyy-yyyy-yyy-yyy", "30000000-aaaa-bbbb-cccc-100000000000", com.azure.core.util.Context.NONE);
    }
}
