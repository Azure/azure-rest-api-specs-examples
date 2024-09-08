
/**
 * Samples for BillingRoleAssignments GetByEnrollmentAccount.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/
     * billingRoleAssignmentGetByEnrollmentAccount.json
     */
    /**
     * Sample code: BillingRoleAssignmentGetByEnrollmentAccount.
     * 
     * @param manager Entry point to BillingManager.
     */
    public static void
        billingRoleAssignmentGetByEnrollmentAccount(com.azure.resourcemanager.billing.BillingManager manager) {
        manager.billingRoleAssignments().getByEnrollmentAccountWithResponse("7898901", "225314",
            "9dfd08c2-62a3-4d47-85bd-1cdba1408402", com.azure.core.util.Context.NONE);
    }
}
