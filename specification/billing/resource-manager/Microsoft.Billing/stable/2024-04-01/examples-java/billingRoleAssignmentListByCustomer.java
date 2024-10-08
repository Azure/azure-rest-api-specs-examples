
/**
 * Samples for BillingRoleAssignments ListByCustomer.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/
     * billingRoleAssignmentListByCustomer.json
     */
    /**
     * Sample code: BillingRoleAssignmentListByCustomer.
     * 
     * @param manager Entry point to BillingManager.
     */
    public static void billingRoleAssignmentListByCustomer(com.azure.resourcemanager.billing.BillingManager manager) {
        manager.billingRoleAssignments().listByCustomer(
            "00000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2018-09-30", "BKM6-54VH-BG7-PGB",
            "703ab484-dda2-4402-827b-a74513b61e2d", null, null, null, com.azure.core.util.Context.NONE);
    }
}
