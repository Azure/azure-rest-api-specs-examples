
/**
 * Samples for BillingRoleAssignments GetByCustomer.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/
     * billingRoleAssignmentGetByCustomer.json
     */
    /**
     * Sample code: BillingRoleAssignmentGetByCustomer.
     * 
     * @param manager Entry point to BillingManager.
     */
    public static void billingRoleAssignmentGetByCustomer(com.azure.resourcemanager.billing.BillingManager manager) {
        manager.billingRoleAssignments().getByCustomerWithResponse(
            "00000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2018-09-30", "xxxx-xxxx-xxx-xxx",
            "703ab484-dda2-4402-827b-a74513b61e2d",
            "30000000-aaaa-bbbb-cccc-100000000000_6fd330f6-7d26-4aff-b9cf-7bd699f965b9",
            com.azure.core.util.Context.NONE);
    }
}
