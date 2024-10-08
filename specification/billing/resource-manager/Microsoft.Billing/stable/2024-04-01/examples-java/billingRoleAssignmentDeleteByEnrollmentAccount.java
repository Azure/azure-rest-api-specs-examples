
/**
 * Samples for BillingRoleAssignments DeleteByEnrollmentAccount.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/
     * billingRoleAssignmentDeleteByEnrollmentAccount.json
     */
    /**
     * Sample code: BillingRoleAssignmentDeleteByEnrollmentAccount.
     * 
     * @param manager Entry point to BillingManager.
     */
    public static void
        billingRoleAssignmentDeleteByEnrollmentAccount(com.azure.resourcemanager.billing.BillingManager manager) {
        manager.billingRoleAssignments().deleteByEnrollmentAccountWithResponse("8608480", "123456",
            "9dfd08c2-62a3-4d47-85bd-1cdba1408402", com.azure.core.util.Context.NONE);
    }
}
