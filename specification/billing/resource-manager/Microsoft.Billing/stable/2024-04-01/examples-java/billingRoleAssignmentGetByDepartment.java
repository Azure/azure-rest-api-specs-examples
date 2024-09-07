
/**
 * Samples for BillingRoleAssignments GetByDepartment.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/
     * billingRoleAssignmentGetByDepartment.json
     */
    /**
     * Sample code: BillingRoleAssignmentGetByDepartment.
     * 
     * @param manager Entry point to BillingManager.
     */
    public static void billingRoleAssignmentGetByDepartment(com.azure.resourcemanager.billing.BillingManager manager) {
        manager.billingRoleAssignments().getByDepartmentWithResponse("7898901", "225314",
            "9dfd08c2-62a3-4d47-85bd-1cdba1408402", com.azure.core.util.Context.NONE);
    }
}
