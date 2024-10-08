
/**
 * Samples for BillingRoleAssignments ResolveByInvoiceSection.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/
     * resolveBillingRoleAssignmentByInvoiceSection.json
     */
    /**
     * Sample code: ResolveBillingRoleAssignmentByInvoiceSection.
     * 
     * @param manager Entry point to BillingManager.
     */
    public static void
        resolveBillingRoleAssignmentByInvoiceSection(com.azure.resourcemanager.billing.BillingManager manager) {
        manager.billingRoleAssignments().resolveByInvoiceSection(
            "00000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2018-09-30", "BKM6-54VH-BG7-PGB",
            "xxxx-xxxx-xxx-xxx", null, null, com.azure.core.util.Context.NONE);
    }
}
