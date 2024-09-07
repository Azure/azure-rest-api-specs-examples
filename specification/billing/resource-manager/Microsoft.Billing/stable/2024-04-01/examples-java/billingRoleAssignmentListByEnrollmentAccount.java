
/**
 * Samples for BillingRoleAssignments ListByEnrollmentAccount.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/
     * billingRoleAssignmentListByEnrollmentAccount.json
     */
    /**
     * Sample code: BillingRoleAssignmentListByEnrollmentAccount.
     * 
     * @param manager Entry point to BillingManager.
     */
    public static void
        billingRoleAssignmentListByEnrollmentAccount(com.azure.resourcemanager.billing.BillingManager manager) {
        manager.billingRoleAssignments().listByEnrollmentAccount("6100092", "123456", com.azure.core.util.Context.NONE);
    }
}
