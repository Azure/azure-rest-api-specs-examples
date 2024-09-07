
/**
 * Samples for BillingRoleAssignments ListByBillingProfile.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/stable/2024-04-01/examples/
     * billingRoleAssignmentListByBillingProfile.json
     */
    /**
     * Sample code: BillingRoleAssignmentListByBillingProfile.
     * 
     * @param manager Entry point to BillingManager.
     */
    public static void
        billingRoleAssignmentListByBillingProfile(com.azure.resourcemanager.billing.BillingManager manager) {
        manager.billingRoleAssignments().listByBillingProfile(
            "00000000-0000-0000-0000-000000000000:00000000-0000-0000-0000-000000000000_2018-09-30", "xxxx-xxxx-xxx-xxx",
            null, null, null, com.azure.core.util.Context.NONE);
    }
}
