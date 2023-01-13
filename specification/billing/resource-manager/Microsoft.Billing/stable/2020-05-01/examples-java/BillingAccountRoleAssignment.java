/** Samples for BillingRoleAssignments GetByBillingAccount. */
public final class Main {
    /*
     * x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/stable/2020-05-01/examples/BillingAccountRoleAssignment.json
     */
    /**
     * Sample code: BillingAccountRoleAssignment.
     *
     * @param manager Entry point to BillingManager.
     */
    public static void billingAccountRoleAssignment(com.azure.resourcemanager.billing.BillingManager manager) {
        manager
            .billingRoleAssignments()
            .getByBillingAccountWithResponse(
                "{billingAccountName}", "{billingRoleAssignmentId}", com.azure.core.util.Context.NONE);
    }
}
