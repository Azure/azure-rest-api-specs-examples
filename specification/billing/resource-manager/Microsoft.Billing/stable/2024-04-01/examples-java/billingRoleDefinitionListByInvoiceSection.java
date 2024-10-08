
/**
 * Samples for BillingRoleDefinition ListByInvoiceSection.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/
     * billingRoleDefinitionListByInvoiceSection.json
     */
    /**
     * Sample code: BillingRoleDefinitionListByInvoiceSection.
     * 
     * @param manager Entry point to BillingManager.
     */
    public static void
        billingRoleDefinitionListByInvoiceSection(com.azure.resourcemanager.billing.BillingManager manager) {
        manager.billingRoleDefinitions().listByInvoiceSection(
            "10000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2019-05-31", "xxxx-xxxx-xxx-xxx",
            "yyyy-yyyy-yyy-yyy", com.azure.core.util.Context.NONE);
    }
}
