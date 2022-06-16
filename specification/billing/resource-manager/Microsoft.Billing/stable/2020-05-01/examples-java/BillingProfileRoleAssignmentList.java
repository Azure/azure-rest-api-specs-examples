import com.azure.core.util.Context;

/** Samples for BillingRoleAssignments ListByBillingProfile. */
public final class Main {
    /*
     * x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/stable/2020-05-01/examples/BillingProfileRoleAssignmentList.json
     */
    /**
     * Sample code: BillingProfileRoleAssignmentList.
     *
     * @param manager Entry point to BillingManager.
     */
    public static void billingProfileRoleAssignmentList(com.azure.resourcemanager.billing.BillingManager manager) {
        manager
            .billingRoleAssignments()
            .listByBillingProfile("{billingAccountName}", "{billingProfileName}", Context.NONE);
    }
}
