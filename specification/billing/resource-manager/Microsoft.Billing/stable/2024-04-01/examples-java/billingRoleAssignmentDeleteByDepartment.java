
/**
 * Samples for BillingRoleAssignments DeleteByDepartment.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/
     * billingRoleAssignmentDeleteByDepartment.json
     */
    /**
     * Sample code: BillingRoleAssignmentDeleteByDepartment.
     * 
     * @param manager Entry point to BillingManager.
     */
    public static void
        billingRoleAssignmentDeleteByDepartment(com.azure.resourcemanager.billing.BillingManager manager) {
        manager.billingRoleAssignments().deleteByDepartmentWithResponse("8608480", "123456",
            "9dfd08c2-62a3-4d47-85bd-1cdba1408402", com.azure.core.util.Context.NONE);
    }
}
