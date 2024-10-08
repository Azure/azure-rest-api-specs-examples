
/**
 * Samples for BillingRoleAssignments DeleteByCustomer.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/
     * billingRoleAssignmentDeleteByCustomer.json
     */
    /**
     * Sample code: BillingRoleAssignmentDeleteByCustomer.
     * 
     * @param manager Entry point to BillingManager.
     */
    public static void billingRoleAssignmentDeleteByCustomer(com.azure.resourcemanager.billing.BillingManager manager) {
        manager.billingRoleAssignments().deleteByCustomerWithResponse(
            "00000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2018-09-30", "BKM6-54VH-BG7-PGB",
            "703ab484-dda2-4402-827b-a74513b61e2d",
            "30000000-aaaa-bbbb-cccc-100000000000_00000000-0000-0000-0000-000000000000",
            com.azure.core.util.Context.NONE);
    }
}
