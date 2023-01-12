/** Samples for BillingRoleAssignments DeleteByBillingProfile. */
public final class Main {
    /*
     * x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/stable/2020-05-01/examples/BillingProfileRoleAssignmentDelete.json
     */
    /**
     * Sample code: BillingProfileRoleAssignmentDelete.
     *
     * @param manager Entry point to BillingManager.
     */
    public static void billingProfileRoleAssignmentDelete(com.azure.resourcemanager.billing.BillingManager manager) {
        manager
            .billingRoleAssignments()
            .deleteByBillingProfileWithResponse(
                "{billingAccountName}",
                "{billingProfileName}",
                "{billingRoleAssignmentName}",
                com.azure.core.util.Context.NONE);
    }
}
