
/**
 * Samples for BillingRoleAssignments ResolveByBillingProfile.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/
     * resolveBillingRoleAssignmentByBillingProfile.json
     */
    /**
     * Sample code: ResolveBillingRoleAssignmentByBillingProfile.
     * 
     * @param manager Entry point to BillingManager.
     */
    public static void
        resolveBillingRoleAssignmentByBillingProfile(com.azure.resourcemanager.billing.BillingManager manager) {
        manager.billingRoleAssignments().resolveByBillingProfile(
            "00000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2018-09-30", "xxxx-xxxx-xxx-xxx",
            null, null, com.azure.core.util.Context.NONE);
    }
}
