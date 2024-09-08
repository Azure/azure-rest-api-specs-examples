
/**
 * Samples for BillingRoleAssignments ListByInvoiceSection.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/
     * billingRoleAssignmentListByInvoiceSection.json
     */
    /**
     * Sample code: BillingRoleAssignmentListByInvoiceSection.
     * 
     * @param manager Entry point to BillingManager.
     */
    public static void
        billingRoleAssignmentListByInvoiceSection(com.azure.resourcemanager.billing.BillingManager manager) {
        manager.billingRoleAssignments().listByInvoiceSection(
            "00000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2018-09-30", "BKM6-54VH-BG7-PGB",
            "xxxx-xxxx-xxx-xxx", null, null, null, com.azure.core.util.Context.NONE);
    }
}
