/** Samples for BillingRoleAssignments GetByBillingProfile. */
public final class Main {
    /*
     * x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/stable/2020-05-01/examples/BillingProfileRoleAssignment.json
     */
    /**
     * Sample code: BillingProfileRoleAssignment.
     *
     * @param manager Entry point to BillingManager.
     */
    public static void billingProfileRoleAssignment(com.azure.resourcemanager.billing.BillingManager manager) {
        manager
            .billingRoleAssignments()
            .getByBillingProfileWithResponse(
                "{billingAccountName}",
                "{billingProfileName}",
                "{billingRoleAssignmentName}",
                com.azure.core.util.Context.NONE);
    }
}
