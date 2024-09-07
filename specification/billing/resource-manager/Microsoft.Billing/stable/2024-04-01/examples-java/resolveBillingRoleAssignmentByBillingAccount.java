
/**
 * Samples for BillingRoleAssignments ResolveByBillingAccount.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/
     * resolveBillingRoleAssignmentByBillingAccount.json
     */
    /**
     * Sample code: ResolveBillingRoleAssignmentByBillingAccount.
     * 
     * @param manager Entry point to BillingManager.
     */
    public static void
        resolveBillingRoleAssignmentByBillingAccount(com.azure.resourcemanager.billing.BillingManager manager) {
        manager.billingRoleAssignments().resolveByBillingAccount(
            "00000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2018-09-30", null, null,
            com.azure.core.util.Context.NONE);
    }
}
