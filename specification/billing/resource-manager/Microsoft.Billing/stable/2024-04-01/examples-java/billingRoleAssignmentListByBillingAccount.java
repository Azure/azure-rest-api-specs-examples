
/**
 * Samples for BillingRoleAssignments ListByBillingAccount.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/
     * billingRoleAssignmentListByBillingAccount.json
     */
    /**
     * Sample code: BillingRoleAssignmentListByBillingAccount.
     * 
     * @param manager Entry point to BillingManager.
     */
    public static void
        billingRoleAssignmentListByBillingAccount(com.azure.resourcemanager.billing.BillingManager manager) {
        manager.billingRoleAssignments().listByBillingAccount(
            "00000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2018-09-30", null, null, null,
            com.azure.core.util.Context.NONE);
    }
}
